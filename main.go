package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("通过给定的正则表达式和字符串，对标准输入的内容进行替换。\n")
		fmt.Printf("说明：%s regexp string\n", os.Args[0])
		fmt.Printf("  或：%s regexp\n", os.Args[0])
        fmt.Printf("\n")
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

	output := re.ReplaceAllLiteralString(GetStdio(), substring)
	fmt.Println(output)
}

// 将标准输入原样转为字符串
func GetStdio() string {
	var sb strings.Builder
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		sb.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "读取标准输入出错：", err)
	}
	// 去掉从标准输入读取字符串时多出的一个空行
	// 换行符\n在字符串中只占一个长度
	return sb.String()[0:(sb.Len() - 1)]
}
