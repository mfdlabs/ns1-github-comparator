package client_test

import (
	"net/http"
	"testing"

	"github.com/mfdlabs/ns1-github-comparator/ns1/client/model/dhcp"

	api "github.com/mfdlabs/ns1-github-comparator/ns1/client"
	"github.com/mfdlabs/ns1-github-comparator/ns1/client/mockns1"
)

func TestDHCPScopeGroup(t *testing.T) {
	mock, doer, err := mockns1.New(t)
	if err != nil {
		t.Fatalf("Error creating mock service: %v", err)
	}
	defer mock.Shutdown()

	client := api.NewClient(doer, api.SetEndpoint("https://"+mock.Address+"/v1/"))

	t.Run("List", func(t *testing.T) {
		defer mock.ClearTestCases()

		client.FollowPagination = true
		sgs := []dhcp.ScopeGroup{
			{Name: "a"},
			{Name: "b"},
			{Name: "c"},
			{Name: "d"},
		}
		err := mock.AddTestCase(http.MethodGet, "/dhcp/scopegroup", http.StatusOK, nil, nil, "", sgs)
		if err != nil {
			t.Fatalf("error adding test case: %v", err)
		}

		respSgs, _, err := client.ScopeGroup.List()
		if err != nil {
			t.Fatalf("error listing IPAM addresses: %v", err)
		}
		if len(respSgs) != len(sgs) {
			t.Errorf("wrong length: want=%d, got=%d", len(sgs), len(respSgs))
		}

		for i, sg := range respSgs {
			if sg.Name != sgs[i].Name {
				t.Errorf("Incorrect name for scope group %d: want=%q, got=%q", i, sgs[i].Name, sg.Name)
			}
		}
	})

	t.Run("Get", func(t *testing.T) {
		defer mock.ClearTestCases()

		client.FollowPagination = false
		sg := dhcp.ScopeGroup{Name: "a"}

		err := mock.AddTestCase(http.MethodGet, "/dhcp/scopegroup/1", http.StatusOK, nil, nil, "", sg)
		if err != nil {
			t.Fatalf("error adding test case: %v", err)
		}

		respAddr, _, err := client.ScopeGroup.Get(1)
		if err != nil {
			t.Fatalf("error getting scpe group: %v", err)
		}
		if respAddr.Name != sg.Name {
			t.Errorf("wrong address returned, want=%+v, got=%+v", sg, respAddr)
		}
	})

	t.Run("Create", func(t *testing.T) {
		defer mock.ClearTestCases()

		t.Run("RequiredParams", func(t *testing.T) {
			sg := &dhcp.ScopeGroup{}
			_, _, err = client.ScopeGroup.Create(sg)
			if err == nil {
				t.Errorf("expected a missing name to result in an error")
			}
		})

		validSecs := 123
		enabled := true
		sg := &dhcp.ScopeGroup{
			Name: "a",
			DHCP4: dhcp.SettingsV4{
				Settings: dhcp.Settings{
					Enabled:           &enabled,
					ValidLifetimeSecs: &validSecs,
				},
			},
		}
		err := mock.AddTestCase(http.MethodPut, "/dhcp/scopegroup", http.StatusCreated, nil, nil, sg, sg)
		if err != nil {
			t.Fatalf("error adding test case: %v", err)
		}
		respSG, _, err := client.ScopeGroup.Create(sg)
		if err != nil {
			t.Fatalf("error creating scope group: %v", err)
		}
		if respSG.Name != sg.Name || *respSG.DHCP4.ValidLifetimeSecs != *sg.DHCP4.ValidLifetimeSecs {
			t.Errorf("wrong scope group returned: want=%+v, got=%+v", sg, respSG)
		}
	})

	t.Run("Edit", func(t *testing.T) {
		t.Run("RequiredParams", func(t *testing.T) {
			sg := &dhcp.ScopeGroup{Name: "a"}
			_, _, err = client.ScopeGroup.Edit(sg)
			if err == nil {
				t.Errorf("expected a missing ID to result in an error")
			}
		})

		defer mock.ClearTestCases()

		enabled := true
		id := 1
		validSecs := 123
		sg := &dhcp.ScopeGroup{
			ID:   &id,
			Name: "a",
			DHCP4: dhcp.SettingsV4{
				Settings: dhcp.Settings{
					Enabled:           &enabled,
					ValidLifetimeSecs: &validSecs,
				},
			},
		}
		err := mock.AddTestCase(http.MethodPost, "/dhcp/scopegroup/1", http.StatusOK, nil, nil, sg, sg)
		if err != nil {
			t.Fatalf("error adding test case: %v", err)
		}

		respSG, _, err := client.ScopeGroup.Edit(sg)
		if err != nil {
			t.Fatalf("error editing scope group: %v", err)
		}
		if respSG.Name != sg.Name || *respSG.DHCP4.ValidLifetimeSecs != *sg.DHCP4.ValidLifetimeSecs {
			t.Errorf("wrong address returned: want=%+v, got=%+v", sg, respSG)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		defer mock.ClearTestCases()

		err := mock.AddTestCase(http.MethodDelete, "/dhcp/scopegroup/1", http.StatusNoContent, nil, nil, "", nil)
		if err != nil {
			t.Fatalf("error adding test case: %v", err)
		}
		_, err = client.ScopeGroup.Delete(1)
		if err != nil {
			t.Fatalf("error deleting scope group: %v", err)
		}
	})
}
