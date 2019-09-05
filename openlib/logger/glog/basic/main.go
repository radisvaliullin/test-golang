package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	// set log to stderr
	flag.Set("logtostderr", "true")
	// set V() define log level
	flag.Set("v", "1")
	flag.Parse()
	defer glog.Flush()

	// simple log
	glog.Info("test glog, info lever log")
	glog.Error("test glog, error lever log")

	// user defined log levels
	if glog.V(0) {
		glog.Info("test glog, V(0) info log")
	}
	glog.V(1).Info("test glog, V(1) info log")
	glog.V(2).Info("test glog, V(2) info log")
}
