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

package lifecycle

import (
	"fmt"
	"os"
	"syscall"

	"github.com/golang/glog"

	"github.com/mfdlabs/ns1-github-comparator/cache"
	lifecycleconstants "github.com/mfdlabs/ns1-github-comparator/constants/lifecycle_constants"
	"github.com/mfdlabs/ns1-github-comparator/sg"
)

func OnExit(sig chan os.Signal) {
	signal := <-sig

	sg.SendMail(fmt.Sprintf(lifecycleconstants.OnExitCalled, signal))
	glog.Infof(lifecycleconstants.OnExitCalled, signal)

	switch signal {
	case syscall.SIGTERM:
		glog.Info(lifecycleconstants.SIGTERMRecieved)
		cache.PurgeCache()
		os.Exit(0)
	case syscall.SIGINT:
		glog.Info(lifecycleconstants.SIGINTRecievied)
		os.Exit(0)
	}
}
