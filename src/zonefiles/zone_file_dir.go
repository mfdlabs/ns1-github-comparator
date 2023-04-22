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
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/golang/glog"

	"github.com/mfdlabs/ns1-github-comparator/ns1/client/model/dns"

	zonefilesconstants "github.com/mfdlabs/ns1-github-comparator/constants/zonefiles_constants"
)

func readJsonZoneFilesDirectory(dir string) ([]dns.Zone, error) {
	// First check if the directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		glog.Errorf(zonefilesconstants.JsonZoneFileDirNotExist, dir)
		return nil, err
	}

	// Read all the files in the directory
	files, err := os.ReadDir(dir)

	if err != nil {
		glog.Errorf(zonefilesconstants.FailedToReadJsonZoneFileDir, dir)
		return nil, err
	}

	// Return an error if there are no files in the directory
	if len(files) == 0 {
		glog.Errorf(zonefilesconstants.JsonZoneFileDirEmpty, dir)
		//go:nowarn
		return nil, fmt.Errorf(strings.ToLower(zonefilesconstants.JsonZoneFileDirEmpty), dir)
	}

	// Return an error if there are no json files in the directory
	var jsonFiles []os.FileInfo

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if file.Name()[len(file.Name())-5:] == zonefilesconstants.CachedZoneFileFormat {
			info, _ := file.Info()

			jsonFiles = append(jsonFiles, info)
		}
	}

	if len(jsonFiles) == 0 {
		glog.Errorf(zonefilesconstants.JsonZoneFileDirNoJsonZoneFiles, dir)
		return nil, fmt.Errorf(strings.ToLower(zonefilesconstants.JsonZoneFileDirNoJsonZoneFiles), dir)
	}

	// Create a slice to hold all the zones
	zones := make([]dns.Zone, 0)

	// Iterate over the files
	for _, file := range files {
		// Only go pass if the file is a JSON file
		if file.Name()[len(file.Name())-5:] == zonefilesconstants.CachedZoneFileFormat {
			// Read the file
			fileBytes, err := os.ReadFile(dir + "/" + file.Name())

			if err != nil {
				glog.Errorf(zonefilesconstants.FailedToReadJsonZoneFile, file.Name())
				return nil, err
			}

			// Create a new zone
			zone := dns.Zone{}

			// Unmarshal the file into the zone
			err = json.Unmarshal(fileBytes, &zone)

			if err != nil {
				glog.Errorf(zonefilesconstants.FailedToUnmarshalJsonZoneFile, file.Name())
				return nil, err
			}

			// Add the zone to the slice
			zones = append(zones, zone)
		}
	}

	return zones, nil
}
