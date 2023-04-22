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
	"sync"

	flagconstants "github.com/mfdlabs/ns1-github-comparator/constants/flag_constants"
	"github.com/mfdlabs/ns1-github-comparator/env"
)

var setupOnce sync.Once

// Sets up environment based flags
func SetupFlags() {
	setupOnce.Do(func() {
		// NS1
		env.GetEnvironmentVariableOrFlag(flagconstants.NS1UrlFlagEnvVarName, NS1Url)
		env.GetEnvironmentVariableOrFlag(flagconstants.NS1ApiKeyFlagEnvVarName, NS1ApiKey)
		env.GetEnvironmentVariableOrFlag(flagconstants.NS1ApiTimeoutFlagEnvVarName, NS1ApiTimeout)

		// Other
		env.GetEnvironmentVariableOrFlag(flagconstants.PollingIntervalFlagEnvVarName, PollingInterval)
		env.GetEnvironmentVariableOrFlag(flagconstants.BranchNameFlagEnvVarName, BranchName)
		env.GetEnvironmentVariableOrFlag(flagconstants.RemoteNameFlagEnvVarName, RemoteName)
		env.GetEnvironmentVariableOrFlag(flagconstants.NoPushToGitFlagEnvVarName, NoPushToGit)
		env.GetEnvironmentVariableOrFlag(flagconstants.SkipCertificateVerifyFlagEnvVarName, SkipCertificateVerify)
		env.GetEnvironmentVariableOrFlag(flagconstants.ConcurrentZoneDownloadEnabledFlagEnvVarName, ConcurrentZoneDownloadEnabled)
		env.GetEnvironmentVariableOrFlag(flagconstants.LogWorkTimingFlagEnvVarName, LogWorkTiming)
	})
}
