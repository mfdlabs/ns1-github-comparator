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

import cacheconstants "github.com/mfdlabs/ns1-github-comparator/constants/cache_constants"

// The name of the zone files directory
const ZoneFilesDirectoryName string = "ns1"

// The extension for zone files.
const ZoneFileExtension string = ".zone"

// The format for a zone file
const ZoneFileFormat string = ZoneFilesDirectoryName + "/%s" + ZoneFileExtension

// The full format for a zone file
const FullZoneFileFormat string = "%s/%s" + ZoneFileExtension

// The NS records header
const ZoneNSHeader string = "\n;\n;  Zone NS records\n;\n\n"

// The rest of the records header
const ZoneRecordHeader string = "\n;\n;  Zone records\n;\n\n"

// Record prefix with no TTL
const ZoneRecordPrefixNoTTL string = "%s					  IN	%s	"

// Record prefix with TTL
const ZoneRecordPrefix string = "%s				%d	IN	%s	"

// Top of Zone file
const ZoneHeader string = ";\n;  Database file %s.zone for Default Zone scope in zone %s.\n;      Zone version:  %d\n;\n\n"

// Format for SOA record.
const ZoneSOARecord string = "@ 					  IN	SOA %s. %s. (\n                        		%d         ; serial number\n                        		%d         ; refresh\n                        		%d         ; retry\n                        		%d         ; expire\n                        		%d         ) ; default ttl\n"

// Format for NS record.
const ZoneNSRecord string = "@ 					  IN	NS	%s.\n"

// The extension for the cached zone files.
const CachedZoneFileExtension string = ".json"

// The format for a zone file.
const CachedZoneFileFormat string = cacheconstants.CacheDirectoryName + "/zones/%s" + CachedZoneFileExtension
