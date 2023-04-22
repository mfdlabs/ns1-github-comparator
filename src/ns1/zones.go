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
	"encoding/json"
	"fmt"
	"sync"

	"github.com/golang/glog"
	"github.com/mfdlabs/ns1-github-comparator/ns1/client"

	"github.com/mfdlabs/ns1-github-comparator/ns1/client/model/dns"

	"github.com/mfdlabs/ns1-github-comparator/cache"
	ns1constants "github.com/mfdlabs/ns1-github-comparator/constants/ns1_constants"
	"github.com/mfdlabs/ns1-github-comparator/sg"
)

func zoneExists(zones []*dns.Zone, zoneName string) bool {
	for _, z := range zones {
		if z.Zone == zoneName {
			return true
		}
	}

	return false
}

func downloadZone(client *client.Client, zone string, commitMessage *string) error {
	// zone has to be rewritten here because it has minimal information
	// the reason we advise you keep caches is so this takes less time
	newZone, _, err := client.Zones.Get(zone)
	if err != nil {
		glog.Errorf(ns1constants.FailedToFetchRemoteZone, zone, err)
		sg.SendMail(fmt.Sprintf(ns1constants.FailedToFetchRemoteZone, zone, err))

		return err
	}

	// serialize the zone to json using encoding/json
	json, err := json.Marshal(newZone)
	if err != nil {
		glog.Errorf(ns1constants.FailedToEncodeRemoteZone, zone, err)
		sg.SendMail(fmt.Sprintf(ns1constants.FailedToEncodeRemoteZone, zone, err))

		return err
	}

	// 3. write the json to the cache directory
	err = cache.WriteFileToCache(json, zone)
	*commitMessage += ns1constants.ZoneUpdatedFormat
	if err != nil {
		glog.Errorf(ns1constants.FailedToWriteZoneCacheFile, zone, err)
		sg.SendMail(fmt.Sprintf(ns1constants.FailedToWriteZoneCacheFile, zone, err))

		return err
	}

	return nil
}

func downloadZoneConcurrently(wg *sync.WaitGroup, client *client.Client, zone string, commitMessage *string) {
	go func() {
		downloadZone(client, zone, commitMessage)

		wg.Done()
	}()
}
