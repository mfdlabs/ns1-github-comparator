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

// The name of the SendGridApiKey flag.
const SendGridApiKeyFlagName string = "sendgrid-api-key"

// The description of the SendGridApiKey flag.
const SendGridApiKeyFlagDescription string = "The SendGrid API key. This is optional. (environment variable: SENDGRID_API_KEY)"

// The enviornment variable name for the SendGridApiKey flag.
const SendGridApiKeyFlagEnvVarName string = "SENDGRID_API_KEY"

// The name of the SendGridFrom flag.
const SendGridFromFlagName string = "sendgrid-from"

// The description of the SendGridFrom flag.
const SendGridFromFlagDescription string = "The name to use as the sender. This is required if the API Key is specified. (environment variable: SENDGRID_FROM)"

// The enviornment variable name for the SendGridFrom flag.
const SendGridFromFlagEnvVarName string = "SENDGRID_FROM"

// The name of the SendGridFromEmail flag.
const SendGridFromEmailFlagName string = "sendgrid-from-email"

// The description of the SendGridFromEmail flag.
const SendGridFromEmailFlagDescription string = "The email address to use as the sender. This is required if the API Key is specified. (environment variable: SENDGRID_FROM_EMAIL)"

// The enviornment variable name for the SendGridFromEmail flag.
const SendGridFromEmailFlagEnvVarName string = "SENDGRID_FROM_EMAIL"

// The name of the SendGridMailingList flag.
const SendGridMailingListFlagName string = "sendgrid-mailing-list"

// The description of the SendGridMailingList flag.
const SendGridMailingListFlagDescription string = "The mailing list to send the commit message to. This is required if the API Key is specified. (environment variable: SENDGRID_MAILING_LIST)"

// The enviornment variable name for the SendGridMailingList flag.
const SendGridMailingListFlagEnvVarName string = "SENDGRID_MAILING_LIST"
