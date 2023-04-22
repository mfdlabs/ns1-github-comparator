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

package flags

import (
	"flag"
	"time"

	flagconstants "github.com/mfdlabs/ns1-github-comparator/constants/flag_constants"
)

var (
	// The URL for the NS1 API, warning: It has to be prepended with a slash or ns1 client will just strip out the last part of the URL.
	NS1Url = flag.String(flagconstants.NS1UrlFlagName, "", flagconstants.NS1UrlFlagDescription)

	// The API key for the NS1 API.
	NS1ApiKey = flag.String(flagconstants.NS1ApiKeyFlagName, "", flagconstants.NS1ApiKeyFlagDescription)

	// The max amount of time to wait for the API to respond.
	NS1ApiTimeout = flag.Duration(flagconstants.NS1ApiTimeoutFlagName, 5*time.Minute, flagconstants.NS1ApiTimeoutFlagDescription)
)
