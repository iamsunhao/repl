package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(`操作说明：
将通过管道传入的标准输入进行替换。最多接受两个参数，多余的参数将被丢弃。
		
repl [正则表达式]
	将通过正则表达式匹配到的内容替换为空。
		
repl [正则表达式] substring
	将通过正则表达式匹配到的内容替换为substring。`)
		os.Exit(0)
	}
	substring := ""
	if len(os.Args) > 2 {
		substring = os.Args[2]
	}
	re := regexp.MustCompile(os.Args[1])
	var raw string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		raw += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
	}
	output := re.ReplaceAllLiteralString(raw, substring)
	fmt.Println(output)
}
