/*
 * @Author: rover lon
 * @Date: 2024-02-19 11:35:52
 * @LastEditors: xiaolong.wu xiaolong.wu@mabotech
 * @LastEditTime: 2024-02-28 10:22:37
 * @FilePath: \gotools\utils\publicFunc.go
 * @Description:
 *
 * Copyright (c) 2024 by ${git_name_email}, All Rights Reserved.
 */
package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CreateFileIfNotExists(filePath string) error {
	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// 文件不存在，创建目录
		dir := filepath.Dir(filePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}

		// 创建文件
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		fmt.Println("File created:", filePath)
		return nil
	} else if err != nil {
		// 其他错误
		return err
	}

	// 文件已经存在
	return nil
}

func ReplaceSpecialChars(input string, specialChars []string, replacement string) string {
	for _, char := range specialChars {
		input = strings.ReplaceAll(input, string(char), replacement)
	}

	return input
}
