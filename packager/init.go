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
	Main		string	 `json:"file"`
	All_files	[]string `json:"all_files"`
}

type ExportAs struct {
	ExportName	string	`json:"export_mod_name"`
	ExportVersion	string	`json:"export_mod_version"`
	Path		string	`json:"mod_path"`
	Required_Export bool	`json:"required"`
}

type LangInfo struct {
	LangName	string		`json:"lang_name"`
	LangVersion	string		`json:"lang_version"`
	EA		[]ExportAs	`json:"exports"`

}

func Read_lang_info_package() *LangInfo {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat(dir)

	if err != nil {
		log.Fatal(err)
	}

	file, _err := os.Open(dir + "/lang_info.json")

	if _err != nil {
		log.Fatal(_err)
	}

	decode := json.NewDecoder(file)
	LI := LangInfo{}
	err = decode.Decode(&LI)

	if err != nil {
		log.Fatal(err)
	}

	if len(LI.EA) == 0 {
		dir, _Err := os.Getwd()

		if _Err != nil {
			log.Fatal(_Err)
		}

		LI.EA = append(LI.EA, ExportAs { ExportName: "Language Info", ExportVersion: "0.1.0", Path: dir + "/lang_info.json", Required_Export: true })

		file, e := json.MarshalIndent(LI, "", "\t")

		if e != nil {
			log.Fatal(e)
		}

		err = ioutil.WriteFile("lang_info.json", file, 0644)

		if err != nil {
			log.Fatal(err)
		}
	}

	return &LI
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
		dir, _err := os.Getwd()

		if _err != nil {
			log.Fatal(_err)
		}

		info := PackageInfo{ Author: "Your Name", Version: "0.1.0", Desc: "Your language description" }

		info.Main = dir + "/" + "main.tf"

		_, e := os.Stat(info.Main)

		if e != nil {
			log.Fatal(e)
		}

		info.All_files = append(info.All_files, info.Main)

		files, err := ioutil.ReadDir(dir)

		if err != nil {
			log.Fatal(err)
		}

		for _,f := range files {
			if strings.Contains(f.Name(), ".tf") {
				info.All_files = append(info.All_files, os.Getenv("HOME") + "/src/" + f.Name())
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

		return &PI
	}
}
