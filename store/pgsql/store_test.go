package pgsql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validation(t *testing.T) {
	testCases := []struct {
		desc string
		in   string
		exp  bool
	}{
		{
			desc: "valid string inp",
			in:   t.Name(),
			exp:  false,
		},
		{
			desc: "in valid string inp",
			in:   "",
			exp:  true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := isEmpty(tC.in)
			if got != tC.exp {
				t.Fatalf("got:%v\n,exp:%v\n", got, tC.exp)
			}

		})
	}
}

func Test_conn_string(t *testing.T) {
	got := getConnFromEnv()
	assert.NotEmpty(t, got)
}
