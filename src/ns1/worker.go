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
	"fmt"
	"os"
	"sync"

	"github.com/golang/glog"

	cacheconstants "github.com/mfdlabs/ns1-github-comparator/constants/cache_constants"
	gitconstants "github.com/mfdlabs/ns1-github-comparator/constants/git_constants"
	ns1constants "github.com/mfdlabs/ns1-github-comparator/constants/ns1_constants"
	zonefilesconstants "github.com/mfdlabs/ns1-github-comparator/constants/zonefiles_constants"
	"github.com/mfdlabs/ns1-github-comparator/flags"
	"github.com/mfdlabs/ns1-github-comparator/git"
	"github.com/mfdlabs/ns1-github-comparator/ns1/client"
	"github.com/mfdlabs/ns1-github-comparator/sg"
	"github.com/mfdlabs/ns1-github-comparator/zonefiles"
)

func doWork(client *client.Client) {
	zones, response, err := client.Zones.List()
	if err != nil {
		if response != nil {
			if response.StatusCode == 401 && client.APIKey != "" {
				glog.Errorf(ns1constants.UnauthorizedApiKey, client.APIKey, client.Endpoint)
				sg.SendMail(fmt.Sprintf(ns1constants.UnauthorizedApiKey, client.APIKey, client.Endpoint))

				os.Exit(1)
			}
		}

		glog.Errorf(ns1constants.FailedToFetchRemoteZones, err)
		sg.SendMail(fmt.Sprintf(ns1constants.FailedToFetchRemoteZones, err))

		return
	}

	var commitMessage string = gitconstants.DefaultGitCommitMessage
	// var updatedZones int = 0
	var wg sync.WaitGroup

	allChanged, updatedZones, err := collectActivityLogDiffs(client, &commitMessage)

	if !allChanged && err != nil {
		for _, zone := range updatedZones {
			// 	var oldTimestamp int = 0
			// 	if checkAlreadyExistingTimestamp(zone.Zone, zone.UpdatedAt, &oldTimestamp) {
			// 		continue
			// 	}

			// 	updatedZones++

			// 	if oldTimestamp != 0 {
			// 		glog.Infof(ns1constants.ZoneTimestampDifferent, zone.Zone, oldTimestamp, zone.UpdatedAt)
			// 		sg.SendMail(fmt.Sprintf(ns1constants.ZoneTimestampDifferent, zone.Zone, oldTimestamp, zone.UpdatedAt))
			// 	}

			if *flags.ConcurrentZoneDownloadEnabled {
				wg.Add(1)
				downloadZoneConcurrently(&wg, client, zone, &commitMessage)
			} else {
				if err := downloadZone(client, zone, &commitMessage); err != nil {
					continue
				}
			}
		}
	} else {
		for _, zone := range zones {
			if *flags.ConcurrentZoneDownloadEnabled {
				wg.Add(1)
				downloadZoneConcurrently(&wg, client, zone.Zone, &commitMessage)
			} else {
				if err := downloadZone(client, zone.Zone, &commitMessage); err != nil {
					continue
				}
			}
		}
	}

	// if updatedZones == 0 {
	// 	return
	// }

	if len(updatedZones) == 0 && !allChanged {
		return
	}

	if *flags.ConcurrentZoneDownloadEnabled {
		wg.Wait()
	}

	err = zonefiles.NS1toZoneFile(cacheconstants.CacheDirectoryName, zonefilesconstants.ZoneFilesDirectoryName)
	if err != nil {
		glog.Errorf(ns1constants.FailedToRunNS1toZoneFile, err)
		sg.SendMail(fmt.Sprintf(ns1constants.FailedToRunNS1toZoneFile, err))

		return
	}

	go zoneInvalidationTask(&commitMessage, zones)

	if *flags.NoPushToGit {
		glog.Warningf(ns1constants.DoNotPushGitMessage)
	} else {
		go git.PushToRepo(commitMessage)
	}
}
