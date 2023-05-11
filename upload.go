package main

import (
	"fmt"
	"mime"
)

func init() {
	mime.AddExtensionType(".apk", "application/vnd.android.package-archive")

	mime.AddExtensionType(".flv", "video/x-flv")
	mime.AddExtensionType(".mov", "video/quicktime")
	mime.AddExtensionType(".mp4", "video/mp4")
	mime.AddExtensionType(".rm", "audio/x-pn-realaudio")
	mime.AddExtensionType(".wmv", "video/x-ms-wmv")
	mime.AddExtensionType(".mpg", "video/mpeg")
	mime.AddExtensionType(".avi", "video/x-msvideo")
	mime.AddExtensionType(".3gp", "video/3gpp")
	//	mime.AddExtensionType(".xlsx", "video/3gpp")

}
func main() {
	mtype := mime.TypeByExtension(".xlsx")
	fmt.Println(mtype)
}
