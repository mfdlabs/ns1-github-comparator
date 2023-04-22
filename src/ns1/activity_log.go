package ns1

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/golang/glog"
	"github.com/mfdlabs/ns1-github-comparator/cache"
	cacheconstants "github.com/mfdlabs/ns1-github-comparator/constants/cache_constants"
	ns1constants "github.com/mfdlabs/ns1-github-comparator/constants/ns1_constants"
	"github.com/mfdlabs/ns1-github-comparator/ns1/client"
	"github.com/mfdlabs/ns1-github-comparator/ns1/client/model/account"
	"github.com/mfdlabs/ns1-github-comparator/ns1/client/model/dns"
	"github.com/mfdlabs/ns1-github-comparator/sg"
	diff "github.com/r3labs/diff/v3"
)

// This package compares NS1 Activity Logs to cached activity logs and tries to collect the differences.

func collectActivityLogDiffs(client *client.Client, commitMessage *string) (allChanged bool, changedZones []string, err error) {
	// Collects the cached Activity Logs.
	const cachefileName string = cacheconstants.CacheDirectoryName + "/activity_logs.json"

	// Collects the NS1 Activity Logs.
	activityLogs, _, err := client.Activity.List()
	if err != nil {
		return false, nil, err
	}

	if _, err := os.Stat(cachefileName); err == nil {
		// The cached activity logs exist, so we need to compare them to the NS1 activity logs.
		file, err := os.Open(cachefileName)
		if err != nil {
			glog.Errorf(ns1constants.FailedToReadSurplusCache, err)
			sg.SendMail(fmt.Sprintf(ns1constants.FailedToReadSurplusCache, err))

			return false, nil, err
		}

		defer file.Close()

		var cachedActivity []*account.ActivityLog
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&cachedActivity)

		if err != nil {
			glog.Errorf(ns1constants.FailedToDecodeSurplusCache, err)
			sg.SendMail(fmt.Sprintf(ns1constants.FailedToDecodeSurplusCache, err))

			return false, nil, err
		}

		changelog, err := diff.Diff(cachedActivity, activityLogs)
		if err != nil {
			return false, nil, err
		}

		if len(changelog) == 0 {
			// No changes.
			return false, nil, nil
		}

		completedIndexes := make([]int, 0)

		// There are changes, so we need to collect the zones that were changed. Filter by resource type "record" and "zone".
		for _, change := range changelog {
			// To determine the new changes we need to look at the change's paths, the first part of the path is the index of the change.
			// The second part of the path is the field that was changed.

			if len(change.Path) < 2 {
				continue
			}

			indexAsStr := change.Path[0]
			index, err := strconv.Atoi(indexAsStr)
			if err != nil {
				continue
			}

			// Check if the index is already in the list.
			indexAlreadyInList := false
			for _, i := range completedIndexes {
				if i == index {
					indexAlreadyInList = true
					break
				}
			}

			if indexAlreadyInList {
				continue
			}

			activity := activityLogs[index]
			if activity == nil {
				continue
			}

			if activity.ResourceType == "record" {
				// Cast the resource to a record.
				record, ok := activity.Resource.(*dns.Record)
				if !ok {
					continue
				}

				// Check if the zone is already in the list.
				zoneAlreadyInList := false
				for _, zone := range changedZones {
					if zone == record.Zone {
						zoneAlreadyInList = true
						break
					}
				}

				if !zoneAlreadyInList {
					changedZones = append(changedZones, record.Zone)

					// Add the zone to the commit message. Use "record updated" for "udpate", "record created" for "create" and "record deleted" for "delete".
					if activity.Action == "update" {
						*commitMessage += fmt.Sprintf(ns1constants.RecordUpdatedFormat, record.Domain)
					} else if activity.Action == "create" {
						*commitMessage += fmt.Sprintf(ns1constants.RecordCreatedFormat, record.Domain)
					} else if activity.Action == "delete" {
						*commitMessage += fmt.Sprintf(ns1constants.RecordRemovedFormat, record.Domain)
					}
				}
			}

			if activity.ResourceType == "zone" {
				// Cast the resource to a zone.
				zone, ok := activity.Resource.(*dns.Zone)
				if !ok {
					continue
				}

				// Check if the zone is already in the list.
				zoneAlreadyInList := false
				for _, z := range changedZones {
					if z == zone.Zone {
						zoneAlreadyInList = true
						break
					}
				}

				if !zoneAlreadyInList {
					changedZones = append(changedZones, zone.Zone)
				}
			}

			completedIndexes = append(completedIndexes, index)
		}

		return false, changedZones, nil
	}

	// Write the new activity logs to the cache.

	// Marshal the activity logs.
	activityLogsJSON, err := json.Marshal(activityLogs)
	if err != nil {
		return false, nil, err
	}

	// Write the activity logs to the cache.
	cache.WriteFileToCache(activityLogsJSON, cachefileName)

	// The cached activity logs don't exist, so we assume all zones have changed.
	return true, nil, nil
}
