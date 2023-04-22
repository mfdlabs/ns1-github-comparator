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

package lifecycleconstants

// Logged when the OnExit is called.
const OnExitCalled string = "Received signal '%s'."

// logged when a SIGTERM is received.
const SIGTERMRecieved string = "Received SIGTERM, purging cache...\n"

// Logged when a SIGINT is received.
const SIGINTRecievied string = "Received SIGINIT. Exiting...\n"
