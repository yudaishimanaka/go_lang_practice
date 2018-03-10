package main

import (
	"fmt"
	"os"
	"bufio"
)

func hello(username string) string{
	return "Hello " + username
}

func main() {
	fmt.Println("名前を入力してください")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	fmt.Println(hello(text))
}
