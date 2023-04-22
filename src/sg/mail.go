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
	"fmt"
	"time"

	"github.com/golang/glog"

	"github.com/sendgrid/sendgrid-go/helpers/mail"

	sendgridconstants "github.com/mfdlabs/ns1-github-comparator/constants/sendgrid_constants"
)

func getSgSubject() string {
	// Return {date} {getSgSubject()}
	return time.Now().String() + " " + sendgridconstants.SendGridDefaultSubject
}

// Sends mail to SendGrid with a subject.
func SendMailWithSubject(subject string, body string) {
	if mailClient == nil {
		return
	}

	if machineName != "" {
		body = fmt.Sprintf(sendgridconstants.BodyMutationFormat, machineName, body)
	}

	// Check if the mailing list is just 1 email
	if len(mailingList) == 1 {
		// If it is, we can use the single email as the from email
		message := mail.NewSingleEmail(from, subject, mailingList[0], body, body)

		response, err := mailClient.Send(message)
		if err != nil {
			if response != nil {
				// If it's an access denied error, then just skip it
				if response.StatusCode == 401 {
					return
				}
			}
			glog.Fatal(err)
		}
	} else {
		// If it isn't, we need to send the email to each address
		for _, email := range mailingList {
			message := mail.NewSingleEmail(from, subject, email, body, body)

			response, err := mailClient.Send(message)
			if err != nil {
				if response != nil {
					// If it's an access denied error, then just skip it
					if response.StatusCode == 401 {
						return
					}
				}

				glog.Fatal(err)
			}
		}
	}
}

// Sends mail to SendGrid.
func SendMail(body string) {
	SendMailWithSubject(getSgSubject(), body)
}
