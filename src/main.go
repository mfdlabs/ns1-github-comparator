/*
   Copyright 2022 MFDLABS

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang/glog"

	"github.com/mfdlabs/ns1-github-comparator/cache"
	"github.com/mfdlabs/ns1-github-comparator/constants"
	"github.com/mfdlabs/ns1-github-comparator/flags"
	"github.com/mfdlabs/ns1-github-comparator/lifecycle"
	"github.com/mfdlabs/ns1-github-comparator/ns1"
	"github.com/mfdlabs/ns1-github-comparator/sg"
)

var commitSha string
var buildMode string
var applicationName string

func init() {
	flag.Usage = func() {
		os.Stderr.WriteString(constants.UsageHeader)
		flag.PrintDefaults()
	}

	flag.Set("logtostderr", "true")
	flag.Set("v", "100")

	flag.Parse()
}

func main() {
	if len(os.Args) <= 1 {
		flag.Usage()
		os.Exit(0)
	}

	if *flags.Version {
		os.Stderr.WriteString(fmt.Sprintf(constants.VersionFormat, applicationName, buildMode, commitSha))
		return
	}

	if *flags.Purge {
		glog.Warningf(constants.PurgeInProgress)

		cache.PurgeCache()
		ns1.PurgeZoneFileDir()

		return
	}

	flags.SetupFlags()
	sg.PreSetupSendGrid()

	if *flags.BranchName == "" {
		os.Stderr.WriteString(constants.NoBranchName)
		flag.Usage()
		os.Exit(1)
	}

	// check if the fetch timeout is less than 5 seconds
	if *flags.PollingInterval < 5*time.Second && !*flags.PulseWork {
		os.Stderr.WriteString(constants.PollingIntervalTooSmall)
		flag.Usage()
		os.Exit(1)
	}

	glog.Infof(constants.StartupLogMessage, *flags.BranchName)
	sg.SendMail(fmt.Sprintf(constants.StartupLogMessage, *flags.BranchName))

	if *flags.PulseWork {
		ns1.Pulse()

		os.Exit(0)
		return
	}

	ns1.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	defer lifecycle.OnExit(sig)
}
