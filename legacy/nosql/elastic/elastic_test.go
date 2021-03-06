package elastic_test

import (
	"testing"

	"github.com/nwca/hidalgo/legacy/nosql/elastic"
	_ "github.com/nwca/hidalgo/legacy/nosql/elastic/test"
	"github.com/nwca/hidalgo/legacy/nosql/nosqltest"
)

func TestElastic(t *testing.T) {
	nosqltest.Test(t, elastic.Name)
}

func BenchmarkElastic(b *testing.B) {
	nosqltest.Benchmark(b, elastic.Name)
}
