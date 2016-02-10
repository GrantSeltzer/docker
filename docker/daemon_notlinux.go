// +build daemon,!linux

package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/docker/docker/daemon"
	"github.com/docker/docker/pkg/jsonlog"
)

// configureLogging sets up the logging configuration that handles messages
// generated by the daemon itself
func configureLogging(config *daemon.Config) {
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: jsonlog.RFC3339NanoFixed,
		DisableColors:   config.RawLogs,
	})
}
