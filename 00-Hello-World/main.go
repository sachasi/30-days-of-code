package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')

	fmt.Printf("Hello, World.\n%s", message)
}
