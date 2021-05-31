package packager

import (
	"os"
	"log"
	"encoding/json"
	"io/ioutil"
	"strings"
)

type PackageInfo struct {
	Author		string	 `json:"author"`
	Version		string	 `json:"version"`
	Desc		string	 `json:"desc"`
	File		string	 `json:"file"`
	All_files	[]string `json:"all_files"`
}

func exists() bool {
	_, err := os.Stat("tfpackage.json")

	if err != nil {
		return false
	}
	return true
}

func Package(filename string) *PackageInfo {
	if !exists() {
		info := PackageInfo{ Author: "Your Name", Version: "0.1.0", Desc: "Your language description", File: os.Getenv("HOME") + "/src/" + filename }

		files, err := iouti.Read(os.Getenv("HOME") + "/src")

		if err != nil {
			log.Fatal(err)
		}

		for _,f := range files {
			if strings.Contains(files.Name(), ".tf") {
				info.All_files = append(info.all_files, os.Getenv("HOME") + "/src/" + f.Name())
			}
		}

		file, _ := json.MarshalIndent(info, "", "\t")
		_ = ioutil.WriteFile("tfpackage.json", file, 0644)

		return &info
	} else {
		file, err := os.Open("tfpackage.json")

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		decode := json.NewDecoder(file)
		PI := PackageInfo{}
		err = decode.Decode(&PI)

		if err != nil {
			log.Fatal(err)
		}

		return &info
	}
}
