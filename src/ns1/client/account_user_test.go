package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mfdlabs/ns1-github-comparator/ns1/client/model/account"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		require.NoError(t, err)

		var u account.User
		require.NoError(t, json.Unmarshal(b, &u))
		assert.Nil(t, u.Permissions.Security)
		assert.Nil(t, u.Permissions.DHCP)
		assert.Nil(t, u.Permissions.IPAM)

		w.Write(b)
	}))
	defer ts.Close()
	c := NewClient(nil, SetEndpoint(ts.URL))

	u := &account.User{
		Name:        "name-1",
		Username:    "user-1",
		Email:       "email-1",
		Permissions: account.PermissionsMap{},
	}

	_, err := c.Users.Create(u)
	require.NoError(t, err)
}

func TestCreateDDIUser(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		require.NoError(t, err)

		var u account.User
		require.NoError(t, json.Unmarshal(b, &u))
		assert.NotNil(t, u.Permissions.Security)
		assert.NotNil(t, u.Permissions.DHCP)
		assert.NotNil(t, u.Permissions.IPAM)
		assert.NotNil(t, u.IPWhitelist)
		assert.True(t, u.IPWhitelistStrict)

		w.Write(b)
	}))
	defer ts.Close()
	c := NewClient(nil, SetEndpoint(ts.URL), SetDDIAPI())

	u := &account.User{
		Name:              "name-1",
		Username:          "user-1",
		Email:             "email-1",
		IPWhitelist:       []string{"1.1.1.1"},
		IPWhitelistStrict: true,
		Permissions:       account.PermissionsMap{},
	}

	_, err := c.Users.Create(u)
	require.NoError(t, err)
}
