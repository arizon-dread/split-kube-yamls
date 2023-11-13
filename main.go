package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/arizon-dread/split-kube-yamls/helpers"
)

func main() {
	fn := ""
	o := ""
	tO := ""
	stdin := false
	var y []string
	flag.StringVar(&fn, "f", "", "filename")
	flag.StringVar(&tO, "o", "", "outputDir")

	flag.Parse()
	f := flag.Args()
	for _, flag := range f {
		if flag == "-" {
			stdin = true
		}
	}

	if tO == "" {
		fmt.Printf("Output dir not supplied, specify with `-o path/to/dir`, quitting.\n")
		os.Exit(2)
	} else {
		//read os-specific path to os-agnostic path
		o = filepath.ToSlash(tO)
		//create dirs with os-agnostic path
		err := os.MkdirAll(filepath.FromSlash(o), os.ModePerm)
		if err != nil {
			fmt.Printf("unable to create directory, %v\n", o)
			os.Exit(2)
		}
	}
	if fn != "" {
		c, err := helpers.ReadYamlFileToStringArr(fn)
		if err != nil {
			fmt.Printf("unable to read file %v, quitting\n", fn)
			os.Exit(2)
		}
		y = c
	}
	if stdin {

		strArr := helpers.ReadStdin()
		if len(strArr) > 1 {
			y = strArr
		}
	}

	if len(y) < 2 {
		fmt.Printf("no yaml supplied\n")
		os.Exit(2)
	}

	for _, str := range y {

		fn := ""

		kind, name, err := helpers.GetYamlKindName(str)
		if err != nil {
			continue
		}
		fn = fmt.Sprintf("%v-%v.yaml", name, kind)
		var filePath []string
		filePath = append(filePath, o)
		filePath = append(filePath, fn)
		file := strings.Join(filePath, string(os.PathSeparator))

		werr := helpers.WriteOutput(str, file)
		if werr != nil {
			fmt.Printf("failed to write file %v, error: %v\n", file, werr)
		}
	}

}
