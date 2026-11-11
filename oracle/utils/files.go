package utils

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// FileExists returns true if a file at the passed string exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}
	return true
}

func TooPermissive(fileMode, maxAllowedPerms os.FileMode) bool {
	return fileMode&^maxAllowedPerms != 0
}

//func IsFileOwnedByOracle(fileInfo os.FileInfo) (bool, error) {
//	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
//	if !ok {
//		return false, errors.Errorf("Unable to determine file owner of %s", fileInfo.Name())
//	}
//	return int(stat.Uid) == os.Getuid(), nil
//}

// Ensures that the given path exists, that it's a directory, and that it has
// permissions that are no more permissive than the given ones.
//
// - If the path does not exist, it is created
// - If the path exists, but is not a directory, an error is returned
// - If the path exists, and is a directory, but has the wrong perms, it is chmod'ed
func EnsureDirAndMaxPerms(path string, perms os.FileMode) error {
	stat, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		// Regular error
		return err
	} else if os.IsNotExist(err) {
		// Dir doesn't exist, create it with desired perms
		return os.MkdirAll(path, perms)
	} else if !stat.IsDir() {
		// Path exists, but it's a file, so don't clobber
		return errors.Errorf("%v already exists and is not a directory", path)
	} else if stat.Mode() != perms {
		// Dir exists, but wrong perms, so chmod
		return os.Chmod(path, (stat.Mode() & perms))
	}
	return nil
}

// Writes `data` to `path` with permissions no more permissive than the given ones. The write is atomic:
// the data goes to a temporary file in the same directory, which is synced and then renamed over `path`,
// so a crash leaves the old file intact and a shorter rewrite can never leave a stale tail behind.
func WriteFileWithMaxPerms(path string, data []byte, perms os.FileMode) error {
	tmp, err := os.CreateTemp(filepath.Dir(path), "."+filepath.Base(path)+".tmp-*")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name()) // no-op once renamed

	if err = tmp.Chmod(perms); err == nil {
		if _, err = tmp.Write(data); err == nil {
			err = tmp.Sync()
		}
	}
	if closeErr := tmp.Close(); err == nil {
		err = closeErr
	}
	if err != nil {
		return err
	}
	return RenameDurably(tmp.Name(), path)
}

// RenameDurably renames from to to and syncs the directory, so the rename survives a crash or power cut.
func RenameDurably(from, to string) error {
	if err := os.Rename(from, to); err != nil {
		return err
	}
	dir, err := os.Open(filepath.Dir(to))
	if err != nil {
		return err
	}
	defer dir.Close()
	return dir.Sync()
}

// Copies the file at `srcPath` to `dstPath` and ensures that it has
// permissions that are no more permissive than the given ones.
func CopyFileWithMaxPerms(srcPath, dstPath string, perms os.FileMode) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return errors.Wrap(err, "could not open source file")
	}

	dst, err := os.OpenFile(dstPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perms)
	if err != nil {
		return errors.Wrap(err, "could not open destination file")
	}

	err = EnsureFileMaxPerms(dst, perms)
	if err != nil {
		return errors.Wrap(err, "could not set file permissions")
	}

	_, err = io.Copy(dst, src)
	return errors.Wrap(err, "could not copy file contents")
}

// Ensures that the given file has permissions that are no more
// permissive than the given ones.
func EnsureFileMaxPerms(file *os.File, perms os.FileMode) error {
	stat, err := file.Stat()
	if err != nil {
		return err
	}
	if stat.Mode() == perms {
		return nil
	}
	return file.Chmod(stat.Mode() & perms)
}

// Ensures that the file at the given filepath has permissions that are
// no more permissive than the given ones.
func EnsureFilepathMaxPerms(filepath string, perms os.FileMode) error {
	dst, err := os.OpenFile(filepath, os.O_RDWR, perms)
	if err != nil {
		return err
	}

	return EnsureFileMaxPerms(dst, perms)
}

// FilesInDir returns an array of filenames in the directory.
func FilesInDir(dir string) ([]string, error) {
	f, err := os.Open(dir)
	if err != nil {
		return []string{}, err
	}

	r, err := f.Readdirnames(-1)
	if err != nil {
		return []string{}, err
	}

	return r, nil
}

// FileContents returns the contents of a file as a string.
func FileContents(path string) (string, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}
