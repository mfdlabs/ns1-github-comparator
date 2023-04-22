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

package sendgridconstants

import flagconstants "github.com/mfdlabs/ns1-github-comparator/constants/flag_constants"

// Logged if you don't include the FROM.
const NoFromVar string = "If you are using SendGrid, you must specify " + flagconstants.SendGridFromFlagEnvVarName + " or -" + flagconstants.SendGridFromFlagName + ".\n\n"

// Logged if you don't include the FROM_EMAIL.
const NoFromEmailVar string = "If you are using SendGrid, you must specify " + flagconstants.SendGridFromEmailFlagEnvVarName + " or -" + flagconstants.SendGridFromEmailFlagName + ".\n\n"

// Logged if you don't include the MAILING_LIST.
const NoMailingListVar string = "If you are using SendGrid, you must specify " + flagconstants.SendGridMailingListFlagEnvVarName + " or -" + flagconstants.SendGridMailingListFlagName + ".\n\n"

// Logged if you provide an invalid from email.
const InvalidFromEmailVar string = "If you are using SendGrid, the from email you supply must be valid: %s\n\n"

// Logged if you provide an invalid mailing list.
const InvalidMailingListVar string = "No mailing list specified\n"

// Logged when setup complete
const SendGridSetupComplete string = "SendGrid setup complete.\n"
