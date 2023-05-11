package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	new_bucket := "pic_bucket"
	new_file_name := "new_picname.jpg"
	wos_token := "xxxyyyjjj"
	ttl := 3000
	overwrite := 1
	data := fmt.Sprintf("%s:%s:%s:%d:%d", new_bucket, new_file_name, wos_token, ttl, overwrite)
	encodedEntryURI := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedEntryURI)
}
