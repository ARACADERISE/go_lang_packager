package main

import (
	"TFPackager/packager"
	"fmt"
)

func main() {
	//packager.Read_command_line()
	packager.Read_commands()
	info := packager.Read_info_package("lang_info.json")
	fmt.Println(info)
}
