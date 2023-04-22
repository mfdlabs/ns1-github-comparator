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

package ns1

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"sync"

	"github.com/mfdlabs/ns1-github-comparator/ns1/client"

	"github.com/mfdlabs/ns1-github-comparator/flags"
)

var (
	clientOnce sync.Once
	apiClient  *client.Client
)

// Creates the NS1 Client.
func PrepareNS1Client() *client.Client {
	clientOnce.Do(func() {
		doer := &http.Client{
			Timeout: *flags.NS1ApiTimeout,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: *flags.SkipCertificateVerify},
			},
		}

		apiClient = client.NewClient(doer, func(c *client.Client) {
			if *flags.NS1Url != "" {
				c.Endpoint, _ = url.Parse(*flags.NS1Url)
			}

			if *flags.NS1ApiKey != "" {
				c.APIKey = *flags.NS1ApiKey
			}
		})
	})

	return apiClient
}
