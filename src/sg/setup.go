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

package sg

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/golang/glog"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	flagconstants "github.com/mfdlabs/ns1-github-comparator/constants/flag_constants"
	sendgridconstants "github.com/mfdlabs/ns1-github-comparator/constants/sendgrid_constants"
	"github.com/mfdlabs/ns1-github-comparator/env"
	"github.com/mfdlabs/ns1-github-comparator/flags"
)

var (
	setupOnce   sync.Once
	mailingList []*mail.Email    = []*mail.Email{}
	mailClient  *sendgrid.Client = nil
	from        *mail.Email      = nil
	machineName string           = ""
)

// Called in main.main(). Sets up the SG API.
func PreSetupSendGrid() {
	setupOnce.Do(func() {
		machineName, _ = os.Hostname()

		env.GetEnvironmentVariableOrFlag(flagconstants.SendGridApiKeyFlagEnvVarName, flags.SendGridApiKey)
		env.GetEnvironmentVariableOrFlag(flagconstants.SendGridFromEmailFlagEnvVarName, flags.SendGridFrom)
		env.GetEnvironmentVariableOrFlag(flagconstants.SendGridFromEmailFlagEnvVarName, flags.SendGridFromEmail)
		env.GetEnvironmentVariableOrFlag(flagconstants.SendGridMailingListFlagEnvVarName, flags.SendGridMailingList)

		if *flags.SendGridApiKey != "" {
			if *flags.SendGridFrom == "" {
				os.Stderr.WriteString(sendgridconstants.NoFromVar)
				flag.Usage()
				os.Exit(1)
			}

			if *flags.SendGridFromEmail == "" {
				os.Stderr.WriteString(sendgridconstants.NoFromEmailVar)
				flag.Usage()
				os.Exit(1)
			}

			if *flags.SendGridMailingList == "" {
				os.Stderr.WriteString(sendgridconstants.NoMailingListVar)
				flag.Usage()
				os.Exit(1)
			}

			_, err := mail.ParseEmail(*flags.SendGridFromEmail)

			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf(sendgridconstants.InvalidFromEmailVar, err))
				flag.Usage()
				os.Exit(1)
			}

			setupSendGrid()
		}
	})
}

func setupSendGrid() {
	// Setup the sendgrid client
	mailClient = sendgrid.NewSendClient(*flags.SendGridApiKey)

	// Setup the from email
	from = mail.NewEmail(*flags.SendGridFrom, *flags.SendGridFromEmail)

	// Setup the mailing list
	for _, email := range strings.Split(*flags.SendGridMailingList, ",") {
		email, err := mail.ParseEmail(email)
		if err != nil {
			glog.Fatal(err)
		}

		mailingList = append(mailingList, email)
	}

	if len(mailingList) == 0 {
		glog.Fatal(sendgridconstants.InvalidMailingListVar)
	}

	glog.Info(sendgridconstants.SendGridSetupComplete)
}
