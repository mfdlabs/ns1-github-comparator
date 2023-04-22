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

	flagconstants "github.com/mfdlabs/ns1-github-comparator/constants/flag_constants"
)

var (
	// The SendGrid API key. This is optional.
	SendGridApiKey = flag.String(flagconstants.SendGridApiKeyFlagName, "", flagconstants.SendGridApiKeyFlagDescription)

	// The name to use as the sender. This is required if the API Key is specified.
	SendGridFrom = flag.String(flagconstants.SendGridFromFlagName, "", flagconstants.SendGridFromFlagDescription)

	// The email address to use as the sender. This is required if the API Key is specified.
	SendGridFromEmail = flag.String(flagconstants.SendGridFromEmailFlagName, "", flagconstants.SendGridFromEmailFlagDescription)

	// The mailing list to send the commit message to. This is required if the API Key is specified.
	SendGridMailingList = flag.String(flagconstants.SendGridMailingListFlagName, "", flagconstants.SendGridMailingListFlagDescription)
)
