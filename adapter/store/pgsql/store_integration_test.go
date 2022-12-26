//go:build integration

package pgsql

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_new_connection(t *testing.T) {
	testCases := []struct {
		desc     string
		setupEnv bool
		expErr   bool
	}{
		{
			desc:     "in-valid conn string",
			setupEnv: true,
			expErr:   true,
		},
		{
			desc:     "valid connection",
			setupEnv: false,
			expErr:   false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.setupEnv {
				setupEnv(t)
				defer unsetEnv(t)
			}
			got, err := New()
			assert.NotNil(t, got)
			if tC.expErr {
				t.Logf("err received:%v\n", err)
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

		})
	}
}

func unsetEnv(t *testing.T) {
	if err := os.Unsetenv(USERNAME); err != nil {
		t.Fatal(err)
	}
	if err := os.Unsetenv(PASSWORD); err != nil {
		t.Fatal(err)
	}
	if err := os.Unsetenv(HOST); err != nil {
		t.Fatal(err)
	}
	if err := os.Unsetenv(DBNAME); err != nil {
		t.Fatal(err)
	}

}

func setupEnv(t *testing.T) {
	os.Setenv(USERNAME, t.Name())
	os.Setenv(PASSWORD, t.Name())
	os.Setenv(HOST, t.Name())
	os.Setenv(DBNAME, t.Name())

}

func cleanupData(newTestConn *pgStore) {
	cleanData := fmt.Sprintf("truncate table %s;", "public.test")
	if _, err := newTestConn.connection.Exec(cleanData); err != nil {
		panic(err)
	}

}

func Test_store_operation(t *testing.T) {
	newTestConn, err := New()
	assert.Nil(t, err)
	t.Run("insert data in to store", func(t *testing.T) {
		cleanupData(newTestConn)
		err := newTestConn.PutData()
		assert.Nil(t, err)
	})
	t.Run("get by id", func(t *testing.T) {
		testID := 1
		dbytes, err := newTestConn.GetByID(testID)
		assert.Nil(t, err)
		assert.NotEmpty(t, string(dbytes))
	})
	t.Run("get all", func(t *testing.T) {
		err := newTestConn.GetAll()
		assert.Nil(t, err)
	})
}

// unlogged table bench mark
// goos: linux
// goarch: amd64
// pkg: github.com/ehrktia/performance-stats/adapter/store/pgsql
// cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
// Benchmark_get-8              182           6441938 ns/op
// Benchmark_get-8              210           6317881 ns/op
// Benchmark_get-8              196           6253561 ns/op
// Benchmark_get-8              186           6593169 ns/op
// Benchmark_get-8              169           6401667 ns/op
// Benchmark_get-8              175           6146071 ns/op
// Benchmark_get-8              198           6301205 ns/op
// Benchmark_get-8              177           6885528 ns/op
// Benchmark_get-8              194           6264441 ns/op
// Benchmark_get-8              182           6635519 ns/op
// PASS
// ok      github.com/ehrktia/performance-stats/adapter/store/pgsql        17.931s

// benchmark for logged table
// goos: linux
// goarch: amd64
// pkg: github.com/ehrktia/performance-stats/adapter/store/pgsql
// cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
// Benchmark_get-8              176           6643421 ns/op
// Benchmark_get-8              174           7344905 ns/op
// Benchmark_get-8              186           6461425 ns/op
// Benchmark_get-8              193           7046342 ns/op
// Benchmark_get-8              170           6284813 ns/op
// Benchmark_get-8              188           7900479 ns/op
// Benchmark_get-8              188           6287658 ns/op
// Benchmark_get-8              180           6766485 ns/op
// Benchmark_get-8              196           6868959 ns/op
// Benchmark_get-8              188           6288429 ns/op
// PASS
// ok      github.com/ehrktia/performance-stats/adapter/store/pgsql        17.627s
func Benchmark_get(b *testing.B) {
	newTestConn, err := New()
	assert.Nil(b, err)
	for i := 0; i < b.N; i++ {
		if err := newTestConn.GetAll(); err != nil {
			b.Fatal(err)
		}

	}
}
