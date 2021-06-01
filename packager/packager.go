package packager

import (
	"os"
	"log"
	//"fmt"
	"encoding/json"
	"io/ioutil"
)

type DefaultJson struct {
	Author		string	 `json:"author"`
	Version		string	 `json:"version"`
	Desc		string	 `json:"desc"`
	Filename	string   `json:"file"`
	All_files	[]string `json:"all_files"`
}

type PreReq struct {
	file	string
}

func Read_commands() interface{} {
	var args interface{} = os.Args

	switch args.(type) {
		case []string: {
			if len(args.([]string)) > 1 {
				switch os.Args[1] {
					case "prereq": {
						if len(args.([]string)) > 2 {
							PR := PreReq{}

							PR.file = os.Args[2]

							return PR
						} else {
							log.Fatal("Expected ./main prereq filename")
						}
					}
				}
			}
		}
	}

	return Default{}
}

func Read_command_line() {
	info := DefaultJson{ Author: "Your Name", Version: "0.1.0", Desc: "My Language Description" }
	var args interface{} = os.Args

	// Error check it
	switch args.(type) {
		case []string: break
		default: log.Fatal("Expected more than one argument")
	}

	if len(args.([]string)) < 3 {
		log.Fatal("Expected 3 arguments: ./main.o init filename")
	}

	switch os.Args[1] {
		case "init": {
			if len(args.([]string)) == 3 {
				_, err := os.Stat(os.Getenv("HOME") + "/src")

				if err != nil {
					if  os.IsNotExist(err) {
						err = os.Mkdir(os.Getenv("HOME") + "/src", 0755)

						if err != nil {
							log.Fatal(err)
						}
					}
				}
				info.Filename = os.Getenv("HOME") + "/src/" + os.Args[2]
				info.All_files = append(info.All_files, info.Filename)

				file, _ := json.MarshalIndent(info, "", "\t")
				_ = ioutil.WriteFile(os.Getenv("HOME") + "/src/tfpackage.json", file, 0644)

				def := []byte("* DEFAULTLY RENDERED FROM PACKAGE *\n\n\n* INSERT YOUR CODE HERE: *")

				err = ioutil.WriteFile(info.Filename, def, 0644)

				if err != nil {
					log.Fatal(err)
				}
			} else {
				log.Fatal("Expected file argument: ./main.o init filename")
			}
		}
		default: log.Fatal("Unrecognized argument.")
	}
}
