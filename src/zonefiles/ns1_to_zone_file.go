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

package zonefiles

import (
	"os"

	"github.com/golang/glog"

	zonefilesconstants "github.com/mfdlabs/ns1-github-comparator/constants/zonefiles_constants"
)

// Converts NS1 Json Zone file to an actual Zone file.
func NS1toZoneFile(zoneFilesPath string, dnsZoneFilesOutputPath string) error {
	// if the dns zone files output path doesn't exist, create it forcefully
	if _, err := os.Stat(dnsZoneFilesOutputPath); os.IsNotExist(err) {
		os.MkdirAll(dnsZoneFilesOutputPath, 0755)
	}

	// Get the list of zone files.
	zoneFiles, err := readJsonZoneFilesDirectory(zoneFilesPath)

	if err != nil {
		return err
	}

	glog.Infof(zonefilesconstants.ZoneCountLog, len(zoneFiles))

	// Iterate through the zone files.
	for _, zoneFile := range zoneFiles {
		err = jsonZoneFileToZoneFile(&zoneFile, dnsZoneFilesOutputPath)

		if err != nil {
			return err
		}
	}

	return nil
}
