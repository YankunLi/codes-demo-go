package main

import (
	"fmt"
)

func show() error {
	return fmt.Errorf("NULL")
}
func main() {
	var err error
	if true {
		err := show()
		if err != nil {
			fmt.Println("inter")
		}
	}

	if err == nil {
		fmt.Println(err)
	}
}
