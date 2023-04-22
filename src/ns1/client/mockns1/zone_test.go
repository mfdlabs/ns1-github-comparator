package mockns1_test

import (
	"encoding/json"
	"testing"

	api "github.com/mfdlabs/ns1-github-comparator/ns1/client"
	"github.com/mfdlabs/ns1-github-comparator/ns1/client/mockns1"
	"github.com/mfdlabs/ns1-github-comparator/ns1/client/model/dns"
	"github.com/stretchr/testify/require"
)

func TestZone(t *testing.T) {
	mock, doer, err := mockns1.New(t)
	require.Nil(t, err)
	defer mock.Shutdown()

	client := api.NewClient(doer, api.SetEndpoint("https://"+mock.Address+"/v1/"))

	t.Run("AddZoneListTestCase", func(t *testing.T) {
		zones := []*dns.Zone{
			{Zone: "a.list.zone"},
			{Zone: "b.list.zone"},
			{Zone: "c.list.zone"},
			{Zone: "d.list.zone"},
		}

		require.Nil(t, mock.AddZoneListTestCase(nil, nil, zones))

		resp, _, err := client.Zones.List()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, len(zones), len(resp))

		for i := range zones {
			require.Equal(t, zones[i].Zone, resp[i].Zone, i)
		}
	})

	t.Run("AddZoneGetTestCase", func(t *testing.T) {
		zone := &dns.Zone{
			Zone: "get.zone",
			Records: []*dns.ZoneRecord{
				{Domain: "a.get.zone"},
				{Domain: "b.get.zone"},
			},
		}

		require.Nil(t, mock.AddZoneGetTestCase(zone.Zone, nil, nil, zone))

		resp, _, err := client.Zones.Get(zone.Zone)
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, zone.Zone, resp.Zone)
		require.Equal(t, len(zone.Records), len(resp.Records))

		for i := range zone.Records {
			require.Equal(t, zone.Records[i].Domain, resp.Records[i].Domain, i)
		}
	})

	t.Run("AddZoneCreateTestCase", func(t *testing.T) {
		zone := &dns.Zone{
			Zone: "create.zone",
		}
		var resp *dns.Zone
		deepcopy(t, zone, &resp)

		resp.TTL = 42

		require.Nil(t, mock.AddZoneCreateTestCase(nil, nil, zone, resp))
		require.Zero(t, zone.TTL)

		_, err := client.Zones.Create(zone)
		require.Nil(t, err)
		require.Equal(t, zone.TTL, resp.TTL)
	})

	t.Run("AddZoneUpdateTestCase", func(t *testing.T) {
		zone := &dns.Zone{
			Zone: "update.zone",
			TTL:  42,
		}

		require.Nil(t, mock.AddZoneUpdateTestCase(nil, nil, zone, zone))

		_, err := client.Zones.Update(zone)
		require.Nil(t, err)
	})

	t.Run("AddZoneDeleteTestCase", func(t *testing.T) {
		require.Nil(t, mock.AddZoneDeleteTestCase("delete.zone", nil, nil))

		_, err := client.Zones.Delete("delete.zone")
		require.Nil(t, err)
	})
}

func deepcopy(t *testing.T, source, target interface{}) {
	data, err := json.Marshal(source)
	require.Nil(t, err)
	require.Nil(t, json.Unmarshal(data, target))
}
