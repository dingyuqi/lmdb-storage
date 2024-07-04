package lmdb

const (
	mb = 1024 * 1024
)
const (
	defaultMapSize   = 5 * mb //映射的大小
	defaultBlockSize = 10000  //每批put或get的数据块大小
	maxDBNum         = 1      //每个地址下最多有多少个DB
	InitialPar       = 0      //初始化的分区号
	ChLen            = 100
)
