package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/openshift/image-build-daemon/pkg/cmd"
	"github.com/openshift/image-build-daemon/pkg/logs"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	rand.Seed(time.Now().UTC().UnixNano())

	command := cmd.New("")
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
