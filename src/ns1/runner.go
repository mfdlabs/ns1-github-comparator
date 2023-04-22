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

package ns1

import (
	"time"

	"github.com/golang/glog"

	ns1constants "github.com/mfdlabs/ns1-github-comparator/constants/ns1_constants"
	"github.com/mfdlabs/ns1-github-comparator/flags"
)

// Starts the NS1 Job worker.
func Start() {
	client := PrepareNS1Client()

	go func() {
		for {
			// Time how long it takes to do the work
			t := time.Now()
			doWork(client)

			if *flags.LogWorkTiming {
				glog.V(10).Infof(ns1constants.TimingMessage, time.Since(t).String())
			}

			time.Sleep(*flags.PollingInterval)
		}
	}()
}

// Run a single NS1 worker job.
func Pulse() {
	// Time how long it takes to do the work
	t := time.Now()
	doWork(PrepareNS1Client())

	if *flags.LogWorkTiming {
		glog.V(10).Infof(ns1constants.TimingMessage, time.Since(t).String())
	}
}
