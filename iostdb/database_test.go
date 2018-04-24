package iostdb

import (
	"bytes"
	"io/ioutil"
	"os"

	"strconv"
	"sync"
	"testing"
)

var test_values = []string{"", "a", "123", "\x00"}

func TestMemDB_PutGet(t *testing.T) {
	db, _ := NewMemDatabase()
	testPutGet(db, t)
}

func TestLDB_PutGet(t *testing.T) {
	dirname, _ := ioutil.TempDir(os.TempDir(), "ethdb_test_")
	db, _ := NewLDBDatabase(dirname, 0, 0)
	testPutGet(db, t)
}
