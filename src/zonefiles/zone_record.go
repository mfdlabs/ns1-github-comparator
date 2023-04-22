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
	"regexp"

	"github.com/mfdlabs/ns1-github-comparator/ns1/client/model/dns"

	zonefilesconstants "github.com/mfdlabs/ns1-github-comparator/constants/zonefiles_constants"
)

func parseRecord(record *dns.ZoneRecord, zoneName string, defaultTTL int) string {
	recordType := record.Type

	// if the record is NS or SOA, skip it. (We write it via the API response)
	if recordType == "NS" || recordType == "SOA" {
		return ""
	}

	// Write the base of the record, if the record matches the zone name, use @ as the record name, else remove the zone name from the record name.
	recordName := record.Domain

	if recordName == zoneName {
		recordName = "@"
	} else {
		recordName = regexp.MustCompile(fmt.Sprintf(".%s", zoneName)).ReplaceAllString(recordName, "")
	}

	outRecord := ""

	for _, shortAnswer := range record.ShortAns {

		// The record looks like this `@ OR <record_name>				<ttl (if not default ttl)>	IN	<record_type>	<record_data>\n`

		recordString := ""

		if record.TTL == defaultTTL {
			recordString = fmt.Sprintf(zonefilesconstants.ZoneRecordPrefixNoTTL, recordName, recordType)
		} else {
			recordString = fmt.Sprintf(zonefilesconstants.ZoneRecordPrefix, recordName, record.TTL, record.Type)
		} // We basically only need to append `<data_parsed>\n` to the record string.

		// if the type is TXT, we need to add the quotes around the value and escape the quotes.

		// The record can have multiple "short answers" (A, AAAA, CNAME, MX, NS, PTR, SRV, TXT)
		// We need to iterate through the short answers and append them to the record string.
		// Second case is to search regex for `[a-zA-Z]` in the short answer, and also check if it does not end with a period, OR it is a CNAME and does not end with a period.
		if recordType == "TXT" {
			recordString += fmt.Sprintf("\"%s\"\n", shortAnswer)
		} else if (regexp.MustCompile(`[a-zA-Z]`).MatchString(shortAnswer) && shortAnswer[len(shortAnswer)-1:] != ".") || (recordType == "CNAME" && shortAnswer[len(shortAnswer)-1:] != ".") {
			recordString += fmt.Sprintf("%s.\n", shortAnswer)
		} else {
			recordString += fmt.Sprintf("%s\n", shortAnswer)
		}

		outRecord += recordString
	}

	return outRecord
}
