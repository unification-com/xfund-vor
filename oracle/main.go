package main

import (
	"flag"
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
	"net"
	"oracle/config"
	"oracle/version"
	"os"
	"os/signal"
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

var options struct {
	Config       string `short:"c" long:"config" description:"Config path" required:"false" default:"./config.json"`
	PasswordFile string `short:"k" long:"key" description:"Path to file containing decryption key, or the key itself" required:"false"`
	Version      bool   `short:"v" long:"version" description:"Show version information and exit" required:"false"`
}

var parser = flags.NewParser(&options, flags.Default)

func main() {
	var err error

	if _, err := parser.Parse(); err != nil {
		panic(err)
	}

	// show version info and exit
	if options.Version {
		vers := version.NewInfo()
		fmt.Println(vers.String())
		return
	}

	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "main",
		}).Info("exiting oracle daemon...")
		err = Stop()
	}()

	fmt.Println(options.Config)
	config.Conf, err = config.NewConfig(options.Config)
	if err != nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "main",
		}).Error("can't read config file")

		panic(err)
	}
	os.Setenv("ORACLE_PORT", string(config.Conf.Serve.Port))
	os.Setenv("ORACLE_HOST", config.Conf.Serve.Host)

	switch os.Args[1] {
	case "start":
		err = start()
	default:
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "main",
		}).Error("no command specified")
		return
	}

	if err != nil {
		panic(err)
	}
}
