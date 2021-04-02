package keystorage_test

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"oracle/store/keystorage"
	"os"
	"testing"
)

var Log = logrus.New()

func TestKeystorage_NewKeyStorage(t *testing.T) {
	keystoragePath := os.Args[len(os.Args)-1]

	keystore, err := keystorage.NewKeyStorage(Log, keystoragePath)
	if err != nil {
		t.Error(err)
	}
	t.Log(keystore.KeyStore.GetKey())
}

func TestKeystorage_Exists(t *testing.T) {
	keystoragePath := os.Args[len(os.Args)-1]

	keystore, err := keystorage.NewKeyStorage(Log, keystoragePath)
	if err != nil {
		t.Error(err)
	}
	t.Log(keystore.Exists())
}

func TestKeystorage_AddGenerated(t *testing.T) {
	keystoragePath := os.Args[len(os.Args)-1]

	keystore, err := keystorage.NewKeyStorage(Log, keystoragePath)
	if err != nil {
		t.Error(err)
	}
	private, err := keystore.GeneratePrivate("testaccount")
	if err != nil {
		t.Error(err)
	}
	t.Log(private)
}

// rod0gbc3mhyxdiah2vwialx1q3osk5cw
func TestKeystorage_Add(t *testing.T) {
	keystoragePath := os.Args[len(os.Args)-1]

	keystore, err := keystorage.NewKeyStorage(Log, keystoragePath)
	if err != nil {
		t.Error(err)
	}
	err = keystore.CheckToken("rod0gbc3mhyxdiah2vwialx1q3osk5cw")
	if err != nil {
		t.Error(err)
	}
	err = keystore.AddExisting("rootuser", "ba37bd76fa2efb78d29cc55c026786c368e34cd97e64aebe4184f4e822079c74")
	if err != nil {
		t.Error(err)
	}
}

func TestKeystorage_GenerateAndCheckToken(t *testing.T) {
	keystoragePath := os.Args[len(os.Args)-1]

	assert := assert.New(t)
	keystore, err := keystorage.NewKeyStorage(Log, keystoragePath)
	token, err := keystore.GenerateToken()
	if err != nil {
		t.Error(err)
	}
	keystore.KeyStore.Token = ""
	t.Log("token generated: ", token)
	err = keystore.CheckToken(token)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(token, keystore.KeyStore.Token)
}

func TestKeystorage_EncryptDecrypt(t *testing.T) {
	assert := assert.New(t)

	stringToEncrypt := "0x90fc9ac3c1d46d2bff1c57cc24e1068a03ea933489a17a23708cd3b5c23168d6"
	key := "jgz9t811see4sie1f5f1ku1vu04"

	encryptedString, err := keystorage.Encrypt(stringToEncrypt, key)
	if err != nil {
		t.Error(err)
	}
	t.Log("encrypted string: ", encryptedString)
	decryptedString, err := keystorage.Decrypt(encryptedString, key)
	if err != nil {
		t.Error(err)
	}
	t.Log("decrypted string: ", decryptedString)

	assert.Equal(stringToEncrypt, decryptedString)
}

func TestKeystorage_Decrypt(t *testing.T) {
	assert := assert.New(t)

	stringToEncrypt := "0x90fc9ac3c1d46d2bff1c57cc24e1068a03ea933489a17a23708cd3b5c23168d6"
	stringToDecrypt := "WRava3YlduujSq8OCmGW6u3MAPQjsTfBUVlaz_Jv90OcG3SbPNmD_S-xmWPW3natOIBEXUicKgI0bGY9X1OVI7sEZZW7T6YKYWs2clLWO8A="
	key := "jfclqiinxy5ccgrn4am5rbv10mlo3mnf"

	t.Log("encrypted string: ", stringToDecrypt)
	decryptedString, err := keystorage.Decrypt(stringToDecrypt, key)
	if err != nil {
		t.Error(err)
	}
	t.Log("decrypted string: ", decryptedString)

	assert.Equal(stringToEncrypt, decryptedString)
}

func TestKeystorage_GetByAccount(t *testing.T) {
	keystoragePath := os.Args[len(os.Args)-1]

	assert := assert.New(t)
	privateKeyExpected := "0x90fc9ac3c1d46d2bff1c57cc24e1068a03ea933489a17a23708cd3b5c23168d6"
	keystore, err := keystorage.NewKeyStorage(Log, keystoragePath)
	if err != nil {
		t.Error(err)
	}
	err = keystore.CheckToken("rod0gbc3mhyxdiah2vwialx1q3osk5cw")
	if err != nil {
		t.Error(err)
	}
	keyModel := keystore.GetByUsername("username")
	t.Log(keyModel.Private)

	assert.Equal(keyModel.Private, privateKeyExpected)
}

func TestKeystorage_SetRegistered(t *testing.T) {
	keystoragePath := os.Args[len(os.Args)-1]

	assert := assert.New(t)
	//privateKeyExpected := "0x90fc9ac3c1d46d2bff1c57cc24e1068a03ea933489a17a23708cd3b5c23168d6"
	keystore, err := keystorage.NewKeyStorage(Log, keystoragePath)
	if err != nil {
		t.Error(err)
	}
	err = keystore.CheckToken("rod0gbc3mhyxdiah2vwialx1q3osk5cw")
	if err != nil {
		t.Error(err)
	}
	keyModel := keystore.GetByUsername("username")
	t.Log(keyModel.Private)

	err = keystore.SetRegistered(keyModel.Private)
	if err != nil {
		t.Error(err)
	}

	keystore2, err := keystorage.NewKeyStorage(Log, keystoragePath)
	if err != nil {
		t.Error(err)
	}
	err = keystore2.CheckToken("rod0gbc3mhyxdiah2vwialx1q3osk5cw")
	if err != nil {
		t.Error(err)
	}
	keyModel2 := keystore2.GetByUsername("gkozyrev")
	t.Log(keyModel.Private)

	assert.Equal(keyModel2.Registered, true)
}
