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

package git

import (
	"strings"

	"github.com/mfdlabs/ns1-github-comparator/cmd"
	gitconstants "github.com/mfdlabs/ns1-github-comparator/constants/git_constants"
	zonefilesconstants "github.com/mfdlabs/ns1-github-comparator/constants/zonefiles_constants"
	"github.com/mfdlabs/ns1-github-comparator/flags"
	"github.com/mfdlabs/ns1-github-comparator/sg"
)

func PushToRepo(commitMessage string) {
	cmd.Exec(gitconstants.GitExecutableName, gitconstants.GitStageCommand, zonefilesconstants.ZoneFilesDirectoryName)
	cmd.Exec(gitconstants.GitExecutableName, gitconstants.GitCommitCommand, gitconstants.GitCommitCommandMessageArg, commitMessage)

	_, stdout, _, _ := cmd.ExecuteWithOutput(
		gitconstants.GitExecutableName,
		gitconstants.GitSymbolicRefCommand,
		gitconstants.GitSymbolicRefCommandShortNameArg,
		gitconstants.GitHeadCommit,
	)

	if strings.TrimSpace(stdout) != *flags.BranchName {
		command, _ := cmd.Exec(gitconstants.GitExecutableName, gitconstants.GitShowBranchCommand, *flags.BranchName)

		if command.ProcessState.Success() {
			cmd.Exec(gitconstants.GitExecutableName, gitconstants.GitCheckoutCommand, *flags.BranchName)
		} else {
			cmd.Exec(gitconstants.GitExecutableName, gitconstants.GitCheckoutCommand, gitconstants.GitCheckoutCommandNonExitantBranch, *flags.BranchName)
		}
	}

	cmd.Exec(gitconstants.GitExecutableName, gitconstants.GitPushCommand, *flags.RemoteName, *flags.BranchName)

	if commitMessage != gitconstants.DefaultGitCommitMessage {
		sg.SendMail(commitMessage)
	}
}
