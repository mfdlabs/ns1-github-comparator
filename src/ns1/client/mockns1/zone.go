package mockns1

import (
	"net/http"

	"github.com/mfdlabs/ns1-github-comparator/ns1/client/model/dns"
)

// AddZoneListTestCase sets up a test case for the api.Client.Zones.List()
// function
func (s *Service) AddZoneListTestCase(
	requestHeaders, responseHeaders http.Header,
	response []*dns.Zone,
) error {
	return s.AddTestCase(
		http.MethodGet, "/zones", http.StatusOK, requestHeaders,
		responseHeaders, "", response,
	)
}

// AddZoneGetTestCase sets up a test case for the api.Client.Zones.Get()
// function
func (s *Service) AddZoneGetTestCase(
	name string,
	requestHeaders, responseHeaders http.Header,
	response *dns.Zone,
) error {
	return s.AddTestCase(
		http.MethodGet, "/zones/"+name, http.StatusOK, requestHeaders,
		responseHeaders, "", response,
	)
}

// AddZoneCreateTestCase sets up a test case for the api.Client.Zones.Create()
// function
func (s *Service) AddZoneCreateTestCase(
	requestHeaders, responseHeaders http.Header,
	zone, response *dns.Zone,
) error {
	return s.AddTestCase(
		http.MethodPut, "/zones/"+zone.Zone, http.StatusCreated, requestHeaders,
		responseHeaders, zone, response,
	)
}

// AddZoneUpdateTestCase sets up a test case for the api.Client.Zones.Update()
// function
func (s *Service) AddZoneUpdateTestCase(
	requestHeaders, responseHeaders http.Header,
	zone, response *dns.Zone,
) error {
	return s.AddTestCase(
		http.MethodPost, "/zones/"+zone.Zone, http.StatusOK, requestHeaders,
		responseHeaders, zone, response,
	)
}

// AddZoneDeleteTestCase sets up a test case for the api.Client.Zones.Delete()
// function
func (s *Service) AddZoneDeleteTestCase(
	name string, requestHeaders, responseHeaders http.Header,
) error {
	return s.AddTestCase(
		http.MethodDelete, "/zones/"+name, http.StatusNoContent, requestHeaders,
		responseHeaders, "", "",
	)
}
