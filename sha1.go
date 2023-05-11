package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"

	"bj58.com/wfs2/util"
)

func a(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func b(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func main() {
	var data string = "abcsddddddddddddddddddddddddddddddddd"
	fmt.Printf("MD5 : %s\n", a(data))
	fmt.Printf("SHA1 : %s\n", b(data))
	fmt.Printf("self: %s\n", util.CalculSha1(data))
}
