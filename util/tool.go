package util

import (
	"log"
	"os"
)

func Convert2Bytes(data string) []byte {
	return []byte(data)
}

// MergeMap 合并两个map, 但是不处理其中key的冲突.
func MergeMap(x, y map[string]string) map[string]string {
	if len(x) == 0 {
		return y
	}
	if len(y) == 0 {
		return x
	}
	for k, v := range y {
		x[k] = v
	}
	return x
}

func CreateDir(dir string) error {
	_, err := os.Stat(dir)
	if err != nil {
		// 如果文件夹不存在，则创建它
		if os.IsNotExist(err) {
			err = os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				panic(err)
			}
			log.Println("Create dir:", dir)
		} else {
			return err
		}
	}
	return nil
}

func CollectRes(resCh chan map[string]string, retCh chan map[string]string) {
	ret := make(map[string]string)
	for m := range resCh {
		ret = MergeMap(ret, m)
	}
	retCh <- ret
}
