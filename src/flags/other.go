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
	// Purge all known information.
	Purge = flag.Bool(flagconstants.PurgeFlagName, false, flagconstants.PurgeFlagDescription)

	// Get the app version.
	Version = flag.Bool(flagconstants.VersionFlagName, false, flagconstants.VersionFlagDescription)

	// The polling interval for each fetch operation in the background, this cannot be less than 5 seconds.
	PollingInterval = flag.Duration(flagconstants.PollingIntervalFlagName, time.Minute, flagconstants.PollingIntervalFlagDescription)

	// The name of the branch to use for ingress.
	BranchName = flag.String(flagconstants.BranchNameFlagName, "", flagconstants.BranchNameFlagDescription)

	// The name of the remote to use for ingress.
	RemoteName = flag.String(flagconstants.RemoteNameFlagName, "origin", flagconstants.RemoteNameFlagDescription)

	// Determines if we should just write the zone files locally and not push. Defaults to false.
	NoPushToGit = flag.Bool(flagconstants.NoPushToGitFlagName, false, flagconstants.NoPushToGitFlagDescription)

	// Should we skip verifying the TLS certificate of the remote server? Defaults to false.
	SkipCertificateVerify = flag.Bool(flagconstants.SkipCertificateVerifyFlagName, false, flagconstants.SkipCertificateVerifyFlagDescription)

	// Should the Zones be downloaded in background threads? False by default for now because it may introduce some issues.
	ConcurrentZoneDownloadEnabled = flag.Bool(flagconstants.ConcurrentZoneDownloadEnabledFlagName, false, flagconstants.ConcurrentZoneDownloadEnabledFlagDescription)

	// Should log the amount of time it takes for work to execute.
	LogWorkTiming = flag.Bool(flagconstants.LogWorkTimingFlagName, false, flagconstants.LogWorkTimingFlagName)

	// Should pulse work be done? Basically it will run the work task once and then exit.
	PulseWork = flag.Bool(flagconstants.PulseWorkFlagName, false, flagconstants.PulseWorkFlagDescription)
)
