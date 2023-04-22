package mockns1

import (
	"fmt"
	"net/http"

	"github.com/mfdlabs/ns1-github-comparator/ns1/client/model/pulsar"
)

// AddPulsarJobListTestCase sets up a test case for the api.Client.PulsarJobs.List()
// function
func (s *Service) AddPulsarJobListTestCase(
	appid string,
	requestHeaders, responseHeaders http.Header,
	response []*pulsar.PulsarJob,
) error {
	return s.AddTestCase(
		http.MethodGet, fmt.Sprintf("/pulsar/apps/%s/jobs", appid), http.StatusOK, requestHeaders,
		responseHeaders, "", response,
	)
}

// AddPulsarJobGetTestCase sets up a test case for the api.Client.PulsarJobs.Get()
// function
func (s *Service) AddPulsarJobGetTestCase(
	appid, jobid string,
	requestHeaders, responseHeaders http.Header,
	response *pulsar.PulsarJob,
) error {
	return s.AddTestCase(
		http.MethodGet, fmt.Sprintf("/pulsar/apps/%s/jobs/%s", appid, jobid), http.StatusOK, requestHeaders,
		responseHeaders, "", response,
	)
}

// AddPulsarJobGetTestCase sets up a test case for the api.Client.PulsarJobs.Create()
// function
func (s *Service) AddPulsarJobCreateTestCase(
	requestHeaders, responseHeaders http.Header,
	pulsarJob, response *pulsar.PulsarJob,
) error {
	return s.AddTestCase(
		http.MethodPut, fmt.Sprintf("pulsar/apps/%s/jobs", pulsarJob.AppID), http.StatusOK, requestHeaders,
		responseHeaders, pulsarJob, response,
	)
}

// AddPulsarJobGetTestCase sets up a test case for the api.Client.PulsarJobs.Update()
// function
func (s *Service) AddPulsarJobUpdateTestCase(
	requestHeaders, responseHeaders http.Header,
	pulsarJob, response *pulsar.PulsarJob,
) error {
	return s.AddTestCase(
		http.MethodPost, fmt.Sprintf("pulsar/apps/%s/jobs/%s", pulsarJob.AppID, pulsarJob.JobID), http.StatusOK, requestHeaders,
		responseHeaders, pulsarJob, response,
	)
}

// AddPulsarJobGetTestCase sets up a test case for the api.Client.PulsarJobs.Delete()
// function
func (s *Service) AddPulsarJobDeleteTestCase(
	requestHeaders, responseHeaders http.Header,
	pulsarJob, response *pulsar.PulsarJob,
) error {
	return s.AddTestCase(
		http.MethodDelete, fmt.Sprintf("pulsar/apps/%s/jobs/%s", pulsarJob.AppID, pulsarJob.JobID), http.StatusOK, requestHeaders,
		responseHeaders, "", "",
	)
}
