package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 指定目录路径
	rootDirectory := "/Users/wangxing/test_doc/code_for_leetcode/go"

	// 递归遍历目录
	err := filepath.Walk(rootDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 检查文件是否为 readme.md
		if info.IsDir() {
			return nil // 如果是目录，继续遍历
		}

		if info.Name() == "README.md" {
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
