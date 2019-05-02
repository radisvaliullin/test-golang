package main

import (
	logrus_stack "github.com/Gurpartap/logrus-stack"
	"github.com/sirupsen/logrus"
)

func main() {

	lg := logrus.New()
	lg.Info("new logger")
	lg.AddHook(logrus_stack.StandardHook())
	lg.Info("logrus stack hook")

	lg.Error("test err")
}
