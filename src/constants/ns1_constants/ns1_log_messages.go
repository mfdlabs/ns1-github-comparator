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

package ns1constants

// Logged when cache file fails to open
const FailedToReadZoneCacheFile string = "Error opening zone %s cache file: %s\n"

// Logged when cache file fails to be decoded
const FailedToDecodeZoneCacheFile string = "Error decoding zone %s cache file: %s\n"

// Logged when surplus cache read fails
const FailedToReadSurplusCache string = "Error reading surplus cache: %s\n"

// Logged when surplus cache fails to be decoded
const FailedToDecodeSurplusCache string = "Error decoding surplus cache: %s\n"

// Logged when zone file dir fails to be read
const FailedToReadZoneFileDir string = "Error reading zone file directory: %s\n"

// Logged when remote zone is deleted
const RemoteZoneDeleted string = "Zone %s does not exist in the NS1 API, deleting zone file and cache file...\n"

// Logged when remote zone cannot be fetched
const FailedToFetchRemoteZone string = "Error fetching remote zone %s: %s\n"

// Logged when remote zone fails to be encoded
const FailedToEncodeRemoteZone string = "Error serializing remote zone %s: %s\n"

// Logged when zone cache file fails to be written
const FailedToWriteZoneCacheFile string = "Error writing zone %s cache file: %s\n"

// Logged when API is unauthorized
const UnauthorizedApiKey string = "The API key '%s' is unauthorized to access API '%s'\n"

// Logged when remote zones fail to be fetched
const FailedToFetchRemoteZones string = "Error fetching remote zones: %s\n"

// Logged when the updated timestamp for a zone is different
const ZoneTimestampDifferent string = "Updated timestamp for zone %s is different (old: %d, new: %d), writing to cache...\n"

// Logged when NS1toZoneFile fails
const FailedToRunNS1toZoneFile string = "Error running NS1toZoneFile: %s\n"

// Logged when flags.NoPushToGit is enabled.
const DoNotPushGitMessage string = "!! NOT PUSHING TO GIT REMOTE !!"

// Logged when timing is took.
const TimingMessage string = "Took %s to execute work\n"

// Occurs when zone file dir fails to be purged.
const ZonePurgeFailed string = "Error purging zone file directory: %s\n"
