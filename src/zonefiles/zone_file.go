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
	"fmt"
	"os"
	"regexp"

	zonefilesconstants "github.com/mfdlabs/ns1-github-comparator/constants/zonefiles_constants"
	"github.com/mfdlabs/ns1-github-comparator/ns1/client/model/dns"
)

func getBaseZoneFileHeader(zone *dns.Zone) string {
	zoneName := zone.Zone
	zoneVersion := zone.UpdatedAt
	soa := zone.PrimaryMaster
	soaAuthority := zone.Hostmaster
	soaAuthority = regexp.MustCompile(`@`).ReplaceAllString(soaAuthority, ".")
	header := fmt.Sprintf(zonefilesconstants.ZoneHeader, zoneName, zoneName, zoneVersion)
	soaRecord := fmt.Sprintf(zonefilesconstants.ZoneSOARecord, soa, soaAuthority, zoneVersion, zone.Refresh, zone.Retry, zone.Expiry, zone.TTL)

	header += soaRecord
	header += zonefilesconstants.ZoneNSHeader

	for _, ns := range zone.DNSServers {
		header += fmt.Sprintf(zonefilesconstants.ZoneNSRecord, ns)
	}

	header += zonefilesconstants.ZoneRecordHeader

	for _, record := range zone.Records {
		header += parseRecord(record, zoneName, zone.TTL)
	}

	return header
}

func jsonZoneFileToZoneFile(zone *dns.Zone, outputDir string) error {
	zoneFileHeader := getBaseZoneFileHeader(zone)
	filename := fmt.Sprintf(zonefilesconstants.FullZoneFileFormat, outputDir, zone.Zone)

	return os.WriteFile(filename, []byte(zoneFileHeader), 0644)
}
