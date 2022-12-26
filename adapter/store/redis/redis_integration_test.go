//go:build integration

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
		name            string
		wantErr         bool
		wantCustomSetup bool
	}{
		{
			name:    "valid connection",
			wantErr: false,
		},
		{
			name:            "in valid connection",
			wantErr:         true,
			wantCustomSetup: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.wantCustomSetup {
				setEnvVar(t)
			}
			got, err := New(ctx)
			if test.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, got)
			}
		})
	}
	t.Cleanup(func() {
		unSetEnv(t)
	})
}

func Test_operations(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	testRedisStore, err := New(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, testRedisStore)
	t.Run("put in store", func(t *testing.T) {
		err := testRedisStore.PutData()
		assert.Nil(t, err)
	})
	t.Run("get by id from store", func(t *testing.T) {
		dataBytes, err := testRedisStore.GetByID(1)
		t.Logf("data:%v\n", string(dataBytes))
		assert.Nil(t, err)
		assert.NotEmpty(t, string(dataBytes))
	})

}

func Benchmark_redis_put(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	newTestConn, err := New(ctx)
	assert.Nil(b, err)
	for i := 0; i < b.N; i++ {
		if err := newTestConn.PutData(); err != nil {
			b.Fatal(err)
		}

	}
}
