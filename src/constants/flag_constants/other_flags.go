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

// The name of the Purge flag.
const PurgeFlagName string = "purge"

// The description of the PurgeFlag flag.
const PurgeFlagDescription string = "Purge all information relating to this."

// The name of the Version flag.
const VersionFlagName string = "version"

// The description of the Version flag.
const VersionFlagDescription string = "Get the app version."

// The name of the PollingInterval flag.
const PollingIntervalFlagName string = "polling-interval"

// The description of the PollingInterval flag.
const PollingIntervalFlagDescription string = "The polling interval for each fetch operation in the background, this cannot be less than 5 seconds. (environment variable: POLLING_INTERVAL)"

// The enviornment variable name for the PollingInterval flag.
const PollingIntervalFlagEnvVarName string = "POLLING_INTERVAL"

// The name of the BranchName flag.
const BranchNameFlagName string = "branch-name"

// The description of the BranchName flag.
const BranchNameFlagDescription string = "The name of the branch to use for ingress. (environment variable: BRANCH_NAME)"

// The enviornment variable name for the BranchName flag.
const BranchNameFlagEnvVarName string = "BRANCH_NAME"

// The name of the RemoteName flag.
const RemoteNameFlagName string = "remote-name"

// The description of the RemoteName flag.
const RemoteNameFlagDescription string = "The name of the remote to use for ingress. (environment variable: REMOTE_NAME)"

// The enviornment variable name for the RemoteName flag.
const RemoteNameFlagEnvVarName string = "REMOTE_NAME"

// The name of the NoPushToGit flag.
const NoPushToGitFlagName string = "no-push-to-git"

// The description of the NoPushToGit flag.
const NoPushToGitFlagDescription string = "Determines if we should just write the zone files locally and not push. Defaults to false. (environment variable: NO_PUSH_TO_GIT)"

// The enviornment variable name for the NoPushToGit flag.
const NoPushToGitFlagEnvVarName string = "NO_PUSH_TO_GIT"

// The name of the SkipCertificateVerify flag.
const SkipCertificateVerifyFlagName string = "skip-certificate-verify"

// The description of the SkipCertificateVerify flag.
const SkipCertificateVerifyFlagDescription string = "Should we skip verifying the TLS certificate of the remote server? Defaults to false. (environment variable: SKIP_CERTIFICATE_VERIFY)"

// The enviornment variable name for the SkipCertificateVerify flag.
const SkipCertificateVerifyFlagEnvVarName string = "SKIP_CERTIFICATE_VERIFY"

// The name of the ConcurrentZoneDownloadEnabled flag.
const ConcurrentZoneDownloadEnabledFlagName string = "concurrent-zone-download-enabled"

// The description of the ConcurrentZoneDownloadEnabled flag.
const ConcurrentZoneDownloadEnabledFlagDescription string = "Should the Zones be downloaded in background threads? (environment variable: CONCURRENT_ZONE_DOWNLOAD_ENABLED)"

// The enviornment variable name for the ConcurrentZoneDownloadEnabled flag.
const ConcurrentZoneDownloadEnabledFlagEnvVarName string = "CONCURRENT_ZONE_DOWNLOAD_ENABLED"

// The name of the LogWorkTiming flag.
const LogWorkTimingFlagName string = "log-work-timing"

// The description of the LogWorkTiming flag.
const LogWorkTimingFlagDescription string = "Should log the amount of time it takes for work to execute. (environment variable: LOG_WORK_TIMING)"

// The enviornment variable name for the LogWorkTiming flag.
const LogWorkTimingFlagEnvVarName string = "LOG_WORK_TIMING"

// The name of the PulseWork flag.
const PulseWorkFlagName string = "pulse-work"

// The description of the PulseWork flag.
const PulseWorkFlagDescription string = "Should pulse work be done? Basically it will run the work task once and then exit. (environment variable: PULSE_WORK)"

// The enviornment variable name for the PulseWork flag.
const PulseWorkFlagEnvVarName string = "PULSE_WORK"
