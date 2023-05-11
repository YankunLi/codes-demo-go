package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path, _ := filepath.Abs("./sdf")
	fmt.Printf("%s\n", path)
}
