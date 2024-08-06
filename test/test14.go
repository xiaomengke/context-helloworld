package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// 命令行参数指定目录名
// 遍历读取目录下的文件
func main14() {

	if len(os.Args) != 2 {
		os.Exit(1)
	}

	dir := os.Args[1]
	start, err := os.Stat(dir)
	if err != nil || !start.IsDir() {
		os.Exit(2)
	}

	var targets []string
	filepath.Walk(dir, func(fPath string, fInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fInfo.Mode().IsRegular() {
			return nil
		}

		targets = append(targets, fPath)
		return nil
	})

	for _, target := range targets {
		func() {}() //放到匿名函数内 break改为return就行
		f, err := os.Open(target)
		if err != nil {
			fmt.Println("bad target:", target, "error:", err) //error:too many open files
			break
		}
		defer f.Close() // 在每次 for 语句块结束时，不会关闭文件资源

		// 使用 f 资源
	}
}
