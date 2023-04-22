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

	"github.com/golang/glog"

	ns1constants "github.com/mfdlabs/ns1-github-comparator/constants/ns1_constants"
	zonefilesconstants "github.com/mfdlabs/ns1-github-comparator/constants/zonefiles_constants"
	"github.com/mfdlabs/ns1-github-comparator/sg"
)

// Purge the zone file dir.
func PurgeZoneFileDir() {
	err := os.RemoveAll(zonefilesconstants.ZoneFilesDirectoryName)

	if err != nil {
		glog.Errorf(ns1constants.ZonePurgeFailed, err)

		sg.SendMail(fmt.Sprintf(ns1constants.ZonePurgeFailed, err))
	}
}
