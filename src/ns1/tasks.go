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
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/glog"

	"github.com/mfdlabs/ns1-github-comparator/ns1/client/model/dns"

	ns1constants "github.com/mfdlabs/ns1-github-comparator/constants/ns1_constants"
	zonefilesconstants "github.com/mfdlabs/ns1-github-comparator/constants/zonefiles_constants"
	"github.com/mfdlabs/ns1-github-comparator/sg"
)

func zoneInvalidationTask(commitMessage *string, zones []*dns.Zone) {
	// This function will query all zones within the NS1 API and compare the current state to the state we have stored locally.
	// If the a zone is in the local `ns1` directory, but not in the NS1 API, we will delete the zone file. And if there is a cache file, we will delete it.

	// Get all the .zone files in the `ns1` directory
	files, err := ioutil.ReadDir(zonefilesconstants.ZoneFilesDirectoryName)
	if err != nil {
		glog.Errorf(ns1constants.FailedToReadZoneFileDir, err)
		sg.SendMail(fmt.Sprintf(ns1constants.FailedToReadZoneFileDir, err))

		return
	}

	// Filter out all the files that are not .zone files
	zoneFiles := []os.FileInfo{}
	for _, f := range files {
		if strings.HasSuffix(f.Name(), zonefilesconstants.ZoneFileExtension) {
			zoneFiles = append(zoneFiles, f)
		}
	}

	// Loop through all the zone files and see if they exist in the NS1 API
	for _, f := range zoneFiles {
		// Get the zone name from the file name
		zoneName := strings.TrimSuffix(f.Name(), zonefilesconstants.ZoneFileExtension)

		if !zoneExists(zones, zoneName) {
			// If we can't find a match, delete the zone file and the cache file
			glog.Warningf(ns1constants.RemoteZoneDeleted, zoneName)
			sg.SendMail(fmt.Sprintf(ns1constants.RemoteZoneDeleted, zoneName))

			os.Remove(f.Name())
			os.Remove(fmt.Sprintf(zonefilesconstants.CachedZoneFileFormat, zoneName))
			*commitMessage += fmt.Sprintf(ns1constants.ZoneRemovedFormat, zoneName)
		}
	}
}
