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

package env

import (
	"os"
	"strconv"
	"time"
)

// Reassigns a flag if an enviornment variable is present.
func GetEnvironmentVariableOrFlag(name string, flag interface{}) error {
	if env, ok := os.LookupEnv(name); ok {
		var err error = nil
		switch v := flag.(type) {
		case *uint:
			var r uint64
			r, err = strconv.ParseUint(env, 10, 0)

			*v = uint(r)
		case *uint64:
			*v, err = strconv.ParseUint(env, 10, 64)
		case *int:
			var r int64
			r, err = strconv.ParseInt(env, 10, 0)

			*v = int(r)
		case *int64:
			*v, err = strconv.ParseInt(env, 10, 64)
		case *float64:
			*v, err = strconv.ParseFloat(env, 64)
		case *bool:
			*v, err = strconv.ParseBool(env)
		case *time.Duration:
			*v, err = time.ParseDuration(env)
		case *string:
			*v = env
		default:
			flag = env
		}

		return err
	}

	return nil
}
