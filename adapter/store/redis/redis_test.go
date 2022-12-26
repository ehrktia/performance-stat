package redis

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setEnvVar(t *testing.T) {
	os.Setenv(HOST, t.Name())
	os.Setenv(PORT, t.Name())

}

func unSetEnv(t *testing.T) {
	assert.NoError(t, os.Unsetenv(HOST))
	assert.NoError(t, os.Unsetenv(PORT))

}

func Test_url(t *testing.T) {
	testCases := []struct {
		desc           string
		wantCustomHost bool
		customHost     string
	}{
		{
			desc: "default values",
		},
		{
			desc:           "custom values",
			wantCustomHost: true,
			customHost:     t.Name(),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.wantCustomHost {
				setEnvVar(t)
			}
			got := getURLFromEnv()
			if tC.wantCustomHost {
				assert.Contains(t, got, tC.customHost)
			}

		})
	}
	t.Cleanup(func() {
		unSetEnv(t)
	})
}
