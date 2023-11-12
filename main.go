package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/arizon-dread/split-yaml/helpers"
)

func main() {
	fn := ""
	o := ""
	var y []string
	flag.StringVar(&fn, "f", "", "filename")
	flag.StringVar(&o, "o", "", "outputDir")

	flag.Parse()

	if o == "" {
		fmt.Printf("Output dir not supplied, specify with `-o path/to/dir`, quitting.\n")
		os.Exit(2)
	}
	if fn != "" {
		c, err := helpers.ReadYamlFileToStringArr(fn)
		if err != nil {
			fmt.Printf("unable to read file %v, quitting\n", fn)
			os.Exit(2)
		}
		y = c
	}

	strArr := helpers.ReadStdin()
	if len(strArr) > 1 {
		y = strArr
	}

	if len(y) < 2 {
		fmt.Printf("no yaml supplied\n")
		os.Exit(2)
	}

	for _, str := range y {

		fn := ""
		// kind, name := helpers.GetKindAndNameFromYaml(str)

		// if kind == "" || name == "" {
		// 	continue
		// }
		kind, name, err := helpers.GetYamlKindName(str)
		if err != nil {
			os.Exit(2)
		}
		fn = fmt.Sprintf("%v-%v.yaml", name, kind)
		var filePath []string
		filePath = append(filePath, o)
		filePath = append(filePath, fn)
		helpers.WriteOutput(str, strings.Join(filePath, "/"))
	}

}
