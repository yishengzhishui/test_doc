package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 指定目录路径
	//rootDirectory := "/Users/wangxing/test_doc/code_for_leetcode/go"
	rootDirectory := "/Users/wangxing/go/src/library"

	// 递归遍历目录
	err := filepath.Walk(rootDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 检查文件是否为 readme.md
		if info.IsDir() {
			return nil // 如果是目录，继续遍历
		}

		if strings.Contains(info.Name(), "_test") {
			// 删除文件
			err := os.Remove(path)
			//fmt.Println(path)
			//fmt.Println(info.Name())
			if err != nil {
				fmt.Printf("Error deleting file %s: %v\n", path, err)
			} else {
				fmt.Printf("File %s deleted successfully\n", path)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}
