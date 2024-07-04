package lmdb

import (
	"github.com/bmatsuo/lmdb-go/lmdb"
	"lmdb-storage/util"
	"log"
	"sync"
)

// Get 并发地批量获取数据. 如果查询的string不存在, 则返回的结果中没有该string
func (l *Driver) Get(data map[string]struct{}) (map[string]string, error) {
	var wg sync.WaitGroup
	retMap := make(map[string]string)
	files, err := util.FetchAllParFiles(l.dbRoot)
	if err != nil {
		return retMap, err
	}
	ch := make(chan map[string]string, ChLen)
	fileCh := make(chan string)
	retCh := make(chan map[string]string)
	go util.CollectRes(ch, retCh)
	for i := 0; i <= 30; i++ {
		wg.Add(1)
		go l.getSingleFileData(fileCh, data, ch, &wg)
	}
	for _, file := range files {
		fileCh <- file
	}
	close(fileCh)
	wg.Wait()
	close(ch)
	retMap = <-retCh
	return retMap, err
}

// GetAllReverse 遍历获取分区下所有数据并反转key,val后返回.
func (l *Driver) GetAllReverse() (map[string]string, error) {
	var wg sync.WaitGroup
	retMap := make(map[string]string)
	files, err := util.FetchAllParFiles(l.dbRoot)
	if err != nil {
		return retMap, err
	}
	ch := make(chan map[string]string, ChLen)
	for _, file := range files {
		wg.Add(1)
		go l.cursorAllReverse(file, ch, &wg)
	}
	retCh := make(chan map[string]string)
	go util.CollectRes(ch, retCh)
	wg.Wait()
	close(ch)
	retMap = <-retCh
	return retMap, err
}

func (l *Driver) cursorAllReverse(path string, resCh chan map[string]string, group *sync.WaitGroup) {
	env, err := l.newEnv(path)
	if err != nil {
		group.Done()
		return
	}
	defer func(env *lmdb.Env) {
		err := env.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(env)
	ret := make(map[string]string)
	err = env.View(func(txn *lmdb.Txn) (err error) {
		dbi, err := txn.OpenRoot(0)
		if err != nil {
			group.Done()
			return
		}
		cur, err := txn.OpenCursor(dbi)
		if err != nil {
			group.Done()
			return
		}
		defer cur.Close()
		for {
			k, v, err := cur.Get(nil, nil, lmdb.Next)
			if lmdb.IsNotFound(err) {
				return nil
			}
			ret[string(v)] = string(k)
		}
	})
	if err != nil {
		log.Println(err)
		group.Done()
		return
	}
	resCh <- ret
	group.Done()
}

// 如果没有找到相应的string则retMap里面没有该string
func (l *Driver) getSingleFileData(fileCh chan string, blockData map[string]struct{}, resCh chan map[string]string, group *sync.WaitGroup) {
	for path := range fileCh {
		env, err := l.newEnv(path)
		if err != nil {
			group.Done()
			return
		}
		retMap := make(map[string]string)
		err = env.View(func(txn *lmdb.Txn) (err error) {
			dbi, err := txn.OpenRoot(0)
			if err != nil {
				return err
			}
			for k := range blockData {
				v, err := txn.Get(dbi, util.Convert2Bytes(k))
				if !lmdb.IsNotFound(err) {
					retMap[k] = string(v)
				}
			}
			return nil
		})
		if err != nil {
			log.Println("getAFileData出现错误: ", err)
			group.Done()
			return
		}
		err = env.Close()
		if err != nil {
			log.Println("关闭env出现错误: ", err)
			return
		}
		resCh <- retMap
	}
	group.Done()
}
