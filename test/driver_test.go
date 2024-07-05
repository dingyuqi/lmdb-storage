package test

import (
	"github.com/dingyuqi/lmdb-storage"
	"log"
	"strconv"
	"testing"
)

var TestDataPath = "D:/test_lmdb"

func TestDriver_Put(t *testing.T) {
	log.Println("开始检测Put========")
	driver, err := lmdb.NewDefaultLmdbDriver(TestDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	data := make(map[string]string)
	for i := 0; i < 100; i++ {
		data[strconv.Itoa(i)] = strconv.Itoa(i)
	}
	for i := 0; i < 10; i++ {
		data[strconv.Itoa(i)+"_time"] = strconv.Itoa(i)
	}
	err = driver.Put(data)
	if lmdb.IsKeyExits(err) {
		log.Println("already exist")
	}
	log.Println("Put 完成")

}

func TestDriver_Get(t *testing.T) {
	log.Println("开始检测Get========")
	driver, err := lmdb.NewDefaultLmdbDriver(TestDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	k := map[string]struct{}{"3": {}, "2": {}, "1": {}, "102": {}}
	log.Println("查询:", k)
	result, err := driver.Get(k)
	if err != nil {
		log.Println("get error", err)
		return
	}
	log.Println("最终查找结果为:", result)
}
