package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("通过给定的正则表达式和字符串，对标准输入的内容进行替换。\n")
		fmt.Printf("说明：%s regexp string\n", os.Args[0])
		fmt.Printf("  或：%s regexp\n", os.Args[0])
		fmt.Printf("示例一：\n")
		fmt.Printf("  标准输入：Hello,123\n")
		fmt.Printf("  %s [0-9] A\n", os.Args[0])
		fmt.Printf("  标准输出：Hello,AAA\n")
		fmt.Printf("示例二：\n")
		fmt.Printf("  标准输入：Hello,123\n")
		fmt.Printf("  %s [0-9]\n", os.Args[0])
		fmt.Printf("  标准输出：Hello,\n")
		os.Exit(0)
	}
	substring := ""
	if len(os.Args) > 2 {
		substring = os.Args[2]
	}
	re, err := regexp.Compile(os.Args[1])
	if err != nil {
		fmt.Printf("%s 是一个非法的正则表达式。\n", os.Args[1])
		os.Exit(1)
	}

	output := re.ReplaceAllLiteralString(stdin2string(), substring)
	fmt.Println(output)
}

// stdin2string 将标准输入转为字符串（支持多行，支持超长的单行）
func stdin2string() string {
	var sb strings.Builder
	buff := bufio.NewReader(os.Stdin)

	for {
		line, isPrefix, err := buff.ReadLine()
		// 如果读取到文件结束符，则跳出循环
		if err == io.EOF {
			break
		}
		//  读取到错误则直接退出程序
		if err != nil {
			fmt.Printf("读取标准输入出错：%s", err)
			os.Exit(1)
		}
		// 判断是否一次就将整行读取，如果是，则跳出当前循环读取下一行
		if !isPrefix {
			sb.WriteString(string(line) + "\n")
			continue
		}
		// 读取一行中剩余的部分
		tmp := []byte{}
		for isPrefix {
			tmp, isPrefix, err = buff.ReadLine()
			if err != nil {
				fmt.Printf("读取标准输入出错：%s", err)
				os.Exit(1)
			}
			line = append(line, tmp...)
		}
		sb.WriteString(string(line) + "\n")
	}

	return sb.String()[0:(sb.Len() - 1)]
}
