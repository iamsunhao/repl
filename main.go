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
		fmt.Println(
			`操作说明：
将通过管道传入的标准输入进行替换。最多接受两个参数，多余的参数将被丢弃。
		
repl [正则表达式] [substring]
	将通过正则表达式匹配到的内容替换为空。
		
repl [正则表达式] [substring]
	将通过正则表达式匹配到的内容替换为[substring]。

例：
	echo "abc123" | repl "[a-z]" "0"
	输出：000123`)
		os.Exit(0)
	}
	substring := ""
	if len(os.Args) > 2 {
		substring = os.Args[2]
	}
	re := regexp.MustCompile(os.Args[1])

	output := re.ReplaceAllLiteralString(GetStdio(), substring)
	fmt.Println(output)
}

// 将标准输入原样转为字符串（总会多一行空行）
func GetStdio() string {
	var sb strings.Builder
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		sb.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
	}
	// 换行符\n在字符串中只占一个长度
	return sb.String()[0:(sb.Len() - 1)]
}
