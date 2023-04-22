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

package gitconstants

// The git executable name.
const GitExecutableName string = "git"

// The git stage command
const GitStageCommand string = "add"

// The git commit command
const GitCommitCommand string = "commit"

// The git commit with message arg
const GitCommitCommandMessageArg string = "-m"

// The git symbolic-ref command
const GitSymbolicRefCommand string = "symbolic-ref"

// The git symbolic-ref short-name arg
const GitSymbolicRefCommandShortNameArg string = "--short"

// The git head commit name.
const GitHeadCommit string = "HEAD"

// The git show-branch command
const GitShowBranchCommand string = "show-branch"

// The git checkout command
const GitCheckoutCommand string = "checkout"

// The git checkout to non existent branch arg
const GitCheckoutCommandNonExitantBranch string = "-b"

// The git push command
const GitPushCommand string = "push"

// The default git commit message
const DefaultGitCommitMessage string = "Update to NS1 zones."
