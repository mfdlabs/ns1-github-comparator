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

package flagconstants

// The name of the NS1Url flag.
const NS1UrlFlagName string = "ns1-url"

// The description of the NS1Url flag.
const NS1UrlFlagDescription string = "The URL for the NS1 API, warning: It has to be prepended with a slash or ns1 client will just strip out the last part of the URL. (environment variable: NS1_URL)"

// The enviornment variable name for the NS1Url flag.
const NS1UrlFlagEnvVarName string = "NS1_URL"

// The name of the NS1ApiKey flag.
const NS1ApiKeyFlagName string = "ns1-api-key"

// The description of the NS1ApiKey flag.
const NS1ApiKeyFlagDescription string = "The API key for the NS1 API. (environment variable: NS1_API_KEY)"

// The enviornment variable name for the NS1ApiKey flag.
const NS1ApiKeyFlagEnvVarName string = "NS1_API_KEY"

// The name of the NS1ApiTimeout flag.
const NS1ApiTimeoutFlagName string = "ns1-api-timeout"

// The description of the NS1ApiTimeout flag.
const NS1ApiTimeoutFlagDescription string = "The max amount of time to wait for the API to respond. (environment variable: NS1_API_TIMEOUT)"

// The enviornment variable name for the NS1ApiTimeout flag.
const NS1ApiTimeoutFlagEnvVarName string = "NS1_API_TIMEOUT"
