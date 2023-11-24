package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

func main() {
	input := flag.String("input", "", "要解码的 base64 字符串")
	flag.Parse()

	if *input == "" {
		fmt.Println("请提供要解码的 base64 字符串，使用 -input 参数")
		os.Exit(1)
	}

	decodeBytes, err := base64.StdEncoding.DecodeString(*input)
	if err != nil {
		fmt.Println("解码失败：", err)
		os.Exit(1)
	}

	fmt.Println("解码结果：", string(decodeBytes))
}
