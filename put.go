package lmdb

import (
	"github.com/bmatsuo/lmdb-go/lmdb"
	"lmdb-storage/util"
	"log"
)

// Put 存入新的k/v对.
// 不允许重复存入k/v对, 如果重复则返回报错 MDB_STRING_EXIST
func (l *Driver) Put(data map[string]string) error {
	for _, block := range l.splitData2Blocks(data) {
		for {
			err := l.putBlockData(block)
			if lmdb.IsMapFull(err) {
				l.newTailDbPath()
				continue
			} else if err != nil {
				return err
			}
			break
		}
	}
	return nil
}

func (l *Driver) splitData2Blocks(data map[string]string) []map[string]string {
	var ret []map[string]string
	block := make(map[string]string)
	for k, v := range data {
		if int64(len(block)) < l.blockSize {
			block[k] = v
		} else {
			ret = append(ret, block)
			block = make(map[string]string)
			block[k] = v
		}
	}
	ret = append(ret, block)
	return ret
}

func (l *Driver) putBlockData(blockData map[string]string) error {
	env, err := l.newEnv(l.tailDbPath())
	if err != nil {
		return err
	}
	defer func(env *lmdb.Env) {
		err := env.Close()
		if err != nil {
			return
		}
	}(env)
	updateErr := env.Update(func(txn *lmdb.Txn) error {
		dbi, err := txn.OpenRoot(0)
		if err != nil {
			return err
		}
		for k, v := range blockData {
			err = txn.Put(dbi, util.Convert2Bytes(k), util.Convert2Bytes(v), lmdb.NoOverwrite)
			if IsTxnFull(err) {
				err := env.Close()
				if err != nil {
					return err
				}
				env, err = l.newEnv(l.tailDbPath())
				if err != nil {
					return err
				}
				continue
			}
			if err != nil && !lmdb.IsMapFull(err) {
				log.Println(err, "相关key value: ", k, v)
				return err
			}
			if lmdb.IsMapFull(err) {
				return err
			}
		}
		return nil
	})
	return updateErr
}
