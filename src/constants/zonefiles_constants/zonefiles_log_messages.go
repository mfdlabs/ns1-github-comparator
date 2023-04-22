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

package zonefilesconstants

// Logged when json zone file dir does not exist.
const JsonZoneFileDirNotExist string = "Json zone file directory does not exist: %s\n"

// Logged when we fail to read the json zone file dir
const FailedToReadJsonZoneFileDir string = "Error reading json zone file directory: %s\n"

// Logged when the json zone file dir is empty.
const JsonZoneFileDirEmpty string = "No files in directory: %s\n"

// Logged when the json zone file dir contains no json zone files.
const JsonZoneFileDirNoJsonZoneFiles string = "No json zone files in directory: %s\n"

// Logged when there was an error reading a json zone file.
const FailedToReadJsonZoneFile string = "Error reading json zone file: %s\n"

// Logged when there was an error unmarshalling a json zone file.
const FailedToUnmarshalJsonZoneFile string = "Error unmarshalling json zone file: %s\n"

// Logged to inform user of zone count.
const ZoneCountLog string = "Found %d zone files.\n"
