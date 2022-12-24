package redis

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_connection(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "valid connection",
			wantErr: false,
		},
		{
			name:    "valid connection-with custom buffer",
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := New(ctx)
			if test.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.NotNil(t, got)
			}
		})
	}
}
