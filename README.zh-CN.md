# lmdb-storage

**其他语言版本: [English](./README.md)**

lmdb-storage是一个用于保存键值对数据的Go语言库, 它是[lmdb](http://www.lmdb.tech/doc/starting.html)数据库的一个封装.
该库在[lmdb-go](https://github.com/bmatsuo/lmdb-go)的基础上进一步地增加了批量读写和自动分区的功能, 更适合应对大数据量的数据读写.

## 安装
使用 `go get`安装lmdb-storage

```shell
go get github.com/dingyuqi/lmdb-storage
```

## 用法
### 新建Driver对象
调用`NewLmdbDriver()`可以获得一个Driver对象, 一共提供三个参数:
1. `root`: lmdb数据库存储的根目录
2. `mapSize`: 单个数据库DB的大小
3. `blockSize`: 批量读写时一个批次的key-value的个数

> [!TIP]  
> 如果不想设定mapSize和blockSize, 则可以调用`NewDefaultLmdbDriver()`, 该函数仅需要提供`root`路径参数, 默认mapSize为5MB, blockSize为10000.

### 批量写数据
在新建lmdbDriver后即可对数据进行批量读写. 写数据的操作主要使用`Put()`函数.

> [!IMPORTANT]  
> 为了避免数据冲突, `Put()`函数不接受重复的键进行保存. 不管是同一批次中还是不同批次之间的保存, 只要该键在历史中被重复保存则会立刻触发错误: `IsKeyExits`.


```go
package main

import (
	lmdb "github.com/dingyuqi/lmdb-storage"
	"log"
	"strconv"
)

// TestDataPath lmdb database root path
// please use your own local path
var TestDataPath = "D:/test_lmdb"

func main() {
	// init a default driver
	driver, err := lmdb.NewDefaultLmdbDriver(TestDataPath)
	if err != nil {
		log.Fatalln(err)
	} 
	// prepare key-value data
	data := make(map[string]string)
	for i := 0; i < 100; i++ {
		data[strconv.Itoa(i)] = strconv.Itoa(i)
	}
	// write key-value data into database
	err = driver.Put(data)
	// if one of key in data has been saved before, it will raise an error
	if lmdb.IsKeyExits(err) {
		log.Println("key already exist")
	}
}
```
### 批量读数据

> [!IMPORTANT]  
> 如果查询的某一个key在数据库中搜索不到, 则不会出现在返回结果`result`中.

```go
package main

import (
	lmdb "github.com/dingyuqi/lmdb-storage"
	"log"
)

// TestDataPath lmdb database root path
// please use your own local path
var TestDataPath = "D:/test_lmdb"

func main() {
	// init a default driver
	driver, err := lmdb.NewDefaultLmdbDriver(TestDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	data := map[string]struct{}{"3": {}, "2": {}, "1": {}, "102": {}}
	result, err := driver.Get(data)
	if err != nil {
		log.Println("get error", err)
		return
	}
	log.Println("result:", result)
}
```

## 认证
[MIT](https://choosealicense.com/licenses/mit/) License