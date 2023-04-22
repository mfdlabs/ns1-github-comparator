package mockns1

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stretchr/testify/assert"
)

// ServeHTTP is the request  handler for the mock service. This
// should not be called directly in unit tests.
func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.stopTimer()
	defer s.startTimer()

	if _, exists := s.tests[r.Method]; !exists {
		notFoundResponse(w, "method")
		return
	}

	tests, exists := s.tests[r.Method][r.RequestURI]
	if !exists {
		notFoundResponse(w, "uri")
		return
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte( // nolint: errcheck
			fmt.Sprintf(`{"message": "unable to read request body: %s`, err),
		))
		return
	}

	var test *testCase
	for _, t := range tests {
		if compareBody(t, body) && compareHeaders(t.request.headers, r.Header) {
			test = t
			break
		}
	}

	if test == nil {
		notFoundResponse(w, "no test")
		return
	}

	for k, vals := range test.response.headers {
		w.Header().Set(k, vals[0])
		for _, v := range vals[1:] {
			w.Header().Add(k, v)
		}
	}

	w.WriteHeader(test.status)
	w.Write(test.response.body) // nolint: errcheck
}

func compareBody(test *testCase, body []byte) bool {
	if !test.request.json {
		return assert.Equal(new(testifyT), test.request.body, body)
	}

	return assert.JSONEq(new(testifyT), string(test.request.body), string(body))
}

func notFoundResponse(w http.ResponseWriter, reason string) {
	msg := fmt.Sprintf(`{"message": "request not found: %s"}`, reason)
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(msg)) // nolint: errcheck
}

func compareHeaders(a, b http.Header) bool {
	for key := range a {
		if !compareHeader(key, a, b) {
			return false
		}
	}

	return true
}

func compareHeader(key string, a, b http.Header) bool {
	if _, exists := b[key]; !exists {
		return false
	}

	for _, v := range a[key] {
		if !inList(v, b[key]) {
			return false
		}
	}

	return true
}

func inList(needle string, haystack []string) bool {
	for _, straw := range haystack {
		if straw == needle {
			return true
		}
	}

	return false
}

// Stub "T" for use with github.com/stretchr/testify/assert tests
type testifyT struct{}

func (t *testifyT) Errorf(format string, args ...interface{}) {}
