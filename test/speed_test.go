package test

import (
	"github.com/dingyuqi/lmdb-storage"
	"os"
	"strconv"
	"testing"
)

func generateString(prefix string, count int) map[string]string {
	r := make(map[string]string)
	for i := 0; i < count; i++ {
		r[prefix+strconv.Itoa(i)] = "d"
	}
	return r
}

func TestPurePut(t *testing.T) {
	err := os.RemoveAll("D:/waste/lmdb")
	if err != nil {
		return
	}
	d := lmdb.NewDefaultLmdbDriver(TestDataPath)
	err = d.Put(generateString("孙悟空", 100000))
	if err != nil {
		return
	}
}
