package main

import (
	"flag"
	"github.com/sevlyar/go-daemon"
	"github.com/sirupsen/logrus"
	"oracle/config"
	"os"
)

// Create a new instance of the logger. You can have any number of instances.
var log = logrus.New()

var configuration *config.Config

// To terminate the daemon use:
//  kill `cat oracled.pid`
func main() {
	var err error
	var configFile string

	//// flags declaration using flag package
	flag.StringVar(&configFile, "c", "./config.json", "Specify config json file.Default is ./config.json")

	configuration, err = config.NewConfig("")
	if err != nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "main",
			"action":   "open config file",
			"result":   "can't read config file",
		}).Error()
		panic(err)
	}

	daemonContext := &daemon.Context{
		PidFileName: "oracled.pid",
		PidFilePerm: 0644,
		LogFileName: "oraclog.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[oracled daemon]"},
	}

	d, err := daemonContext.Reborn()
	if err != nil || d != nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "main",
			"action":   "start oracle daemon",
			"result":   "can't start daemon",
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

	switch os.Args[0] {
	case "start":
		err = start()
	default:
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "main",
			"action":   "open log file",
			"result":   "log file is not specified",
		}).Error()
		return
	}
	if err != nil {
		panic(err)
	}
}
