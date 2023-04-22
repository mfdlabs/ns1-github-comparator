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
	"os"

	"github.com/golang/glog"

	"github.com/mfdlabs/ns1-github-comparator/ns1/client/model/dns"

	ns1constants "github.com/mfdlabs/ns1-github-comparator/constants/ns1_constants"
	zonefilesconstants "github.com/mfdlabs/ns1-github-comparator/constants/zonefiles_constants"
	"github.com/mfdlabs/ns1-github-comparator/sg"
)

func checkAlreadyExistingTimestamp(zoneName string, newTimestamp int, oldTimestamp *int) bool {
	cacheFilename := fmt.Sprintf(zonefilesconstants.CachedZoneFileFormat, zoneName)

	if _, err := os.Stat(cacheFilename); err == nil {
		file, err := os.Open(cacheFilename)
		if err != nil {
			glog.Errorf(ns1constants.FailedToReadZoneCacheFile, zoneName, err)
			sg.SendMail(fmt.Sprintf(ns1constants.FailedToReadZoneCacheFile, zoneName, err))

			return false
		}

		defer file.Close()

		var zone dns.Zone
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&zone)

		if err != nil {
			glog.Errorf(ns1constants.FailedToDecodeZoneCacheFile, zoneName, err)
			sg.SendMail(fmt.Sprintf(ns1constants.FailedToDecodeZoneCacheFile, zoneName, err))

			return true
		}

		*oldTimestamp = zone.UpdatedAt

		return zone.UpdatedAt == newTimestamp
	}

	return false
}
