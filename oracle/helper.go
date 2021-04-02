package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"oracle/controller/chainlisten"
	"time"
)

var oracleListener *chainlisten.VORCoordinatorListener

func Stop() (err error) {
	if stop1 {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "start",
			"action":   "stop service",
			"result":   fmt.Sprintf("stopped %d %s", PID, time.Now().String()),
		}).Warning()
	}
	stop1 = true
	log.WithFields(logrus.Fields{
		"package":  "main",
		"function": "start",
		"action":   "stop service",
		"result":   fmt.Sprintf("stopped %d %s", PID, time.Now().String()),
	}).Warning()
	err = e.Shutdown(context.Background())
	if err != nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "start",
			"action":   "stop echo server",
			"result":   err.Error(),
		}).Error()
		err = e.Close()
	}
	oracleListener.Shutdown()
	return err
}
