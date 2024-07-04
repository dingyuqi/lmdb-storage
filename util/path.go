package util

import (
	"os"
	"path"
	"sort"
	"strconv"
)

// FetchAllParFiles 获取dbRoot下所有分区文件的路径, 文件名字按照数字大小顺序排序
func FetchAllParFiles(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	// 让数字的字符串按照数字含义的大小顺序输出
	sort.Slice(files, func(i, j int) bool {
		num1, _ := strconv.Atoi(files[i].Name())
		num2, _ := strconv.Atoi(files[j].Name())
		return num1 < num2
	})
	var fNames []string
	for _, f := range files {
		fNames = append(fNames, path.Join(dir, f.Name()))
	}
	return fNames, err
}
