package lmdb

import "github.com/bmatsuo/lmdb-go/lmdb"

func IsKeyExits(err error) bool {
	return lmdb.IsErrno(err, lmdb.KeyExist)
}

func IsTxnFull(err error) bool {
	return lmdb.IsErrno(err, lmdb.TxnFull)
}
