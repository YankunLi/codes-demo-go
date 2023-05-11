package main

import (
	"fmt"
	"os"
)

func coloringBlue(raw string) string {
	return fmt.Sprintf("%s %s %s", "\033[34m", raw, "\033[0m")
}

func coloringRed(raw string) string {
	return fmt.Sprintf("%s %s %s", "\033[31m", raw, "\033[0m")
}

func coloringNothing(raw string) string {
	return fmt.Sprintf("%s", raw)
}

func main() {
	fmt.Println("vim-go")
	fmt.Fprintf(os.Stdout, "\033[30m 黑色字 \033[0m")
	fmt.Fprintf(os.Stdout, "\033[31m 红色字 \033[0m")
	fmt.Fprintf(os.Stdout, "\033[32m 绿色字 \033[0m")
	fmt.Fprintf(os.Stdout, "\033[33m 黄色字 \033[0m")
	fmt.Fprintf(os.Stdout, "\033[34m 蓝色字 \033[0m")
	fmt.Fprintf(os.Stdout, "\033[35m 紫色字 \033[0m")
	fmt.Fprintf(os.Stdout, "\033[36m 天蓝字 \033[0m")
	fmt.Fprintf(os.Stdout, "\033[37m 白色字 \033[0m")
	fmt.Fprintf(os.Stdout, "%s ", coloringBlue("hello world"))
}
