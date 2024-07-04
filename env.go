// Package lmdb -----------------------------
// @file      : env.go
// @author    : dingyq
// @time      : 2024/7/4 下午3:42
// -------------------------------------------
package lmdb

import (
	"github.com/bmatsuo/lmdb-go/lmdb"
	"log"
	"os"
)

func (l *Driver) newEnv(path string) (*lmdb.Env, error) {
	env, err := lmdb.NewEnv()
	if err != nil {
		log.Println("Error creating env:", err)
		return nil, err
	}
	err = env.SetMapSize(l.mapSize)
	if err != nil {
		log.Println("Error setting map size:", err)
		return nil, err
	}
	err = env.SetMaxDBs(maxDBNum)
	if err != nil {
		log.Println("Error setting max dbs:", err)
		return nil, err
	}
	err = env.Open(path, lmdb.NoMetaSync|lmdb.NoLock, os.ModePerm) //每个事务只将系统缓冲区刷新到磁盘一次，忽略元数据刷新
	if err != nil {
		log.Println("Error opening env:", err)
		return nil, err
	}
	return env, nil
}
