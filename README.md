# lmdb-storage

**Read this in other languages: [简体中文](./README.zh-CN.md)**


lmdb-storage is a Go language library used to save key-value data. It is a further packaging of the [lmdb](http://www.lmdb.tech/doc/starting.html) database.
This library further adds the functions of batch reading and writing and automatic partitioning based on [lmdb-go](https://github.com/bmatsuo/lmdb-go), which is more suitable for dealing with large amounts of data reading and writing.

## Installation
Use the `go get` to install lmdb-storage.
```shell
go get github.com/dingyuqi/lmdb-storage
```
## Usage
### New lmdb driver

Call `NewLmdbDriver()` to get a Driver object, providing a total of three parameters:
1. `root`: the root path to save data
2. `mapSize`: size of one DB file
3. `blockSize`: batch size of reading and writing

> [!TIP]  
>If you do not want to set mapSize and blockSize, you can call `NewDefaultLmdbDriver()`. 
This function only needs to provide the `root` path parameter. The default mapSize is 5MB and blockSize is 10000.

```go
package main

import (
	"github.com/dingyuqi/lmdb-storage"
	"log"
)

// TestDataPath lmdb database root path
var TestDataPath = "D:/test_lmdb"

func main() {
	// init a default driver
	defaultDriver, err := lmdb.NewDefaultLmdbDriver(TestDataPath)
	if err != nil {
		log.Fatalln(err)
	}
	// init a user-defined driver
	definedDriver, err := lmdb.NewLmdbDriver(TestDataPath, 5*1024*1024, 10000)
	if err != nil {
		log.Fatalln(err)
	}
}
```
### Read data in batches

> [!IMPORTANT]  
> If a certain key in the query cannot be searched in the database, it will not appear in the returned result.

```go
package main

import (
	lmdb "github.com/dingyuqi/lmdb-storage"
	"log"
)

func main() {
	var TestDataPath = "D:/test_lmdb"
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

## License
[MIT](https://choosealicense.com/licenses/mit/)