package main

import (
	"flag"
	"github.com/sevlyar/go-daemon"
	"github.com/sirupsen/logrus"
	"net"
	"oracle/config"
	"os"
	"syscall"
)

// Create a new instance of the logger. You can have any number of instances.
var log = logrus.New()
var FD *int = flag.Int("fd", 0, "Server socket FD")
var PID int = syscall.Getpid()
var listener1 net.Listener
var file1 *os.File = nil
var exit1 chan int = make(chan int)
var stop1 = false

// To terminate the daemon use:
//  kill `cat oracled.pid`
func main() {
	var err error
	var configFile string

	//// flags declaration using flag package
	flag.StringVar(&configFile, "c", "./config.json", "Specify config json file.Default is ./config.json")

	config.Conf, err = config.NewConfig(configFile)
	if err != nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "main",
			"action":   "open config file",
			"result":   "can't read config file",
		}).Error()
		panic(err)
	}
	os.Setenv("ORACLE_PORT", string(config.Conf.Serve.Port))
	os.Setenv("ORACLE_HOST", config.Conf.Serve.Host)

	daemonContext := &daemon.Context{
		PidFileName: "oracled.pid",
		PidFilePerm: 0644,
		LogFileName: "oraclog.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		//Args:        []string{"[oracled]"},
	}

	_, err = daemonContext.Reborn()
	if err != nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "main",
			"action":   "start oracle daemon",
			"result":   err,
		}).Error()
		return
	}

	defer daemonContext.Release()

	log.WithFields(logrus.Fields{
		"package":  "main",
		"function": "main",
		"action":   "start oracle daemon",
		"result":   "daemon started",
	}).Info()

	switch os.Args[1] {
	case "start":
		err = start()
	default:
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "main",
			"action":   "open log file",
			"result":   "no command specified",
		}).Error()
		return
	}

	if err != nil {
		panic(err)
	}
}
