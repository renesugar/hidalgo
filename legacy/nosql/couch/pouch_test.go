// +build js

// To run "gopherjs test -v" you must "npm install pouchdb" and "npm install pouchdb-find" for queries.

package couch_test

import (
	"testing"

	"github.com/nwca/hidalgo/legacy/nosql/couch"
	_ "github.com/nwca/hidalgo/legacy/nosql/couch/test"
	"github.com/nwca/hidalgo/legacy/nosql/nosqltest"
)

func TestPouch(t *testing.T) {
	nosqltest.Test(t, couch.NamePouch)
}

func BenchmarkPouch(b *testing.B) {
	nosqltest.Benchmark(b, couch.NamePouch)
}
