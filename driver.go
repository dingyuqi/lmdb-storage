package lmdb

import (
	"lmdb-storage/util"
	"log"
	"os"
	"path"
	"strconv"
)

type Driver struct {
	dbRoot    string // 数据库文件路径
	mapSize   int64  // 单个数据库的大小
	blockSize int64  // 每次操作的数据块大小
}

// NewLmdbDriver 返回一个LmdbDriver对象, 使用传入的参数
func NewLmdbDriver(root string, mapSize, blockSize int64) Driver {
	err := util.CreateDir(root)
	if err != nil {
		log.Println("Create dir error:", err)
		return Driver{}
	}
	return Driver{dbRoot: root, mapSize: mapSize, blockSize: blockSize}
}

// NewDefaultLmdbDriver 返回一个LmdbDriver对象但是使用默认的相关常量
func NewDefaultLmdbDriver(root string) Driver {
	err := util.CreateDir(root)
	if err != nil {
		log.Println("Create dir error:", err)
		return Driver{}
	}
	return Driver{dbRoot: root, mapSize: defaultMapSize, blockSize: defaultBlockSize}
}

// tailDbPath 获取当前最大的分区绝对路径, 如果当前路径下没有任何分区则初始化分区0并返回
func (l *Driver) tailDbPath() string {
	files, _ := util.FetchAllParFiles(l.dbRoot)
	if len(files) == 0 {
		initParPath := path.Join(l.dbRoot, strconv.Itoa(InitialPar))
		_ = os.MkdirAll(initParPath, os.ModePerm)
		return initParPath
	}
	return files[len(files)-1]
}

// newTailDbPath 创建新的分区(老分区号+1)文件夹并返回新创建的路径
func (l *Driver) newTailDbPath() string {
	curPar, _ := strconv.Atoi(path.Base(l.tailDbPath()))
	newPath := path.Join(l.dbRoot, strconv.Itoa(curPar+1))
	_ = os.MkdirAll(newPath, os.ModePerm)
	return newPath
}
