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

package cache

import (
	"fmt"
	"os"

	"github.com/golang/glog"

	cacheconstants "github.com/mfdlabs/ns1-github-comparator/constants/cache_constants"
	"github.com/mfdlabs/ns1-github-comparator/sg"
)

// Purge the cache.
func PurgeCache() {
	err := os.RemoveAll(cacheconstants.CacheDirectoryName)

	if err != nil {
		glog.Errorf(cacheconstants.CachePurgeFailed, err)

		sg.SendMail(fmt.Sprintf(cacheconstants.CachePurgeFailed, err))
	}
}
