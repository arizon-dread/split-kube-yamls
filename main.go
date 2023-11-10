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
	//c, err := readStdin()
	// if err != nil {
	// 	fmt.Printf("unable to read stdin properly. %v, quitting\n", err)
	// 	os.Exit(2)
	// }
	var strArr []string
	strArr = helpers.ReadStdin()
	if len(strArr) > 1 {
		y = strArr
	}

	if len(y) < 2 {
		fmt.Printf("no yaml supplied\n")
		os.Exit(2)
	}
	for _, str := range y {

		fn := ""
		kind, name := helpers.GetKindAndNameFromYaml(str)

		// for k, v := range m {
		// 	if k == "kind" {
		// 		kind = v
		// 	}
		// 	if k == "name" {
		// 		name = v
		// 		break
		// 	}
		// }
		if kind == "" || name == "" {
			continue
		}
		fn = fmt.Sprintf("%v-%v.yaml", name, kind)
		var filePath []string
		filePath = append(filePath, o)
		filePath = append(filePath, fn)
		helpers.WriteOutput(str, strings.Join(filePath, "/"))
	}

}
