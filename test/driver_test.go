package test

import (
	"github.com/dingyuqi/lmdb-storage"
	"log"
	"strconv"
	"testing"
)

var TestDataPath = "D:/test_lmdb"

func TestDriver_Put(t *testing.T) {
	log.Println("Start testing Put() with default driver========")
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
	log.Println("Put finished")

}

func TestDriver_Get(t *testing.T) {
	log.Println("Start testing Get() with default driver========")
	driver, err := lmdb.NewDefaultLmdbDriver(TestDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	k := map[string]struct{}{"3": {}, "2": {}, "1": {}, "102": {}}
	log.Println("start to find keys:", k)
	result, err := driver.Get(k)
	if err != nil {
		log.Println("get error", err)
		return
	}
	log.Println("result is:", result)
}

func TestDefinedDriver_Put(t *testing.T) {
	log.Println("Start testing Put() with defined driver========")
	driver, err := lmdb.NewLmdbDriver(TestDataPath, 1024*1024*5, 1000)
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
	log.Println("Put finished")

}

func TestDefinedDriver_Get(t *testing.T) {
	log.Println("Start testing Get() with defined driver========")
	driver, err := lmdb.NewLmdbDriver(TestDataPath, 1024*1024*5, 1000)
	if err != nil {
		log.Fatalln(err)
	}
	k := map[string]struct{}{"3": {}, "2": {}, "1": {}, "102": {}}
	log.Println("start to find keys:", k)
	result, err := driver.Get(k)
	if err != nil {
		log.Println("get error", err)
		return
	}
	log.Println("result is:", result)
}
