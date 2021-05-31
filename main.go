package main

import (
	"TFPackager/packager"
	"fmt"
)

func main() {
	packager.Read_command_line()
	info := packager.Read_lang_info_package()
	fmt.Println(info)
}
