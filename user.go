package main

import (
	"fmt"
	"os/user"
)

func main() {
	user, err := user.Lookup("root")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*user, err)
}
