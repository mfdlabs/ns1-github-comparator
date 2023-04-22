package mockns1_test

import (
	"net/http"
	"testing"

	"github.com/mfdlabs/ns1-github-comparator/ns1/client/mockns1"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		mock, doer, err := mockns1.New(t)
		require.Nil(t, err)
		require.IsType(t, new(http.Client), doer)
		require.NotNil(t, mock)
		require.NotEmpty(t, mock.Address)
		mock.Shutdown()
	})

	t.Run("Shutdown", func(t *testing.T) {
		mock, _, err := mockns1.New(t)
		require.Nil(t, err)
		require.NotPanics(t, mock.Shutdown)
	})
}
