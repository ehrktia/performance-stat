package pgsql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validate_query_input(t *testing.T) {
	tt := []struct {
		desc string
		inp  int
		want bool
	}{
		{
			desc: "valid id",
			inp:  1,
			want: false,
		},
		{
			desc: "in valid id",
			inp:  0,
			want: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			got := isNotValidID(tc.inp)
			assert.Equal(t, got, tc.want)
		})

	}

}

func Test_build_stmt(t *testing.T) {
	got := buildStatement()
	assert.Contains(t, got, "SELECT")
	assert.NotEmpty(t, got)

}
