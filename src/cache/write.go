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
	"os"

	cacheconstants "github.com/mfdlabs/ns1-github-comparator/constants/cache_constants"
)

func WriteFileToCache(data []byte, cachefileName string) error {
	// create the cache directory if it doesn't exist
	err := os.MkdirAll(cacheconstants.CacheDirectoryName, 0755)
	if err != nil {
		return err
	}

	// write the json to the cache directory
	err = os.WriteFile(cachefileName, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
