package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
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
		c, err := readYamlFileToStringArr(fn)
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
	strArr = readStdin()
	if len(strArr) > 1 {
		y = strArr
	}

	if len(y) < 2 {
		fmt.Printf("no yaml supplied\n")
		os.Exit(2)
	}
	for _, str := range y {

		fn := ""
		kind, name := getKindAndNameFromYaml(str)

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
		writeOutput(str, strings.Join(filePath, "/"))
	}

}
func readYamlFileToStringArr(fn string) ([]string, error) {
	b, err := os.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	str := string(b)
	var strArr []string
	strArr = splitStr(str)
	return strArr, nil
}
func splitStr(s string) []string {
	var strArr []string
	strArr = strings.Split(s, "\n---\n")
	if len(strArr) == 1 {
		if strings.Contains(s, "- apiVersion:") {
			strArr = strings.Split(s, "- apiVersion:")
		}
	}
	var result []string
	for _, str := range strArr {
		result = append(result, "- apiVersion:"+str)
	}
	return result
}

func getKindAndNameFromYaml(y string) (string, string) {
	// var obj interface{}
	// err := yaml.Unmarshal([]byte(s), *obj)
	// if err != nil {
	// 	fmt.Printf("unable to unmarshal yaml, %v", err)
	// 	return "", "", err
	// }
	// name := obj.Metadata.Name
	// kind := obj.kind
	kind := ""
	name := ""

	scanner := bufio.NewScanner(strings.NewReader(y))
	for scanner.Scan() {
		s := scanner.Text()
		b, a, found := strings.Cut(s, ":")
		if found && strings.Trim(strings.ToLower(b), " ") == "kind" {
			kind = strings.Trim(a, " ")
		}
		if found && strings.Trim(strings.ToLower(b), " ") == "name" {
			name = strings.Trim(a, " ")
			break
		}

	}

	return kind, name
	// decode := scheme.Codecs.UniversalDeserializer().Decode
	// obj, install, err := decode([]byte(s), nil, nil)
	// if err != nil {
	// 	fmt.Printf("error decoding yaml, %v", err)
	// }

	// return install.Kind, obj.Metadata.name, nil
}

func readStdin() []string {
	s := bufio.NewScanner(os.Stdin)
	var str string
	var l []string
	for s.Scan() {
		str += "\n" + s.Text()

	}

	l = splitStr(str)
	fmt.Printf("str: %v", str)
	return l
}

func writeOutput(c string, fn string) error {
	return os.WriteFile(fn, []byte(c), 0644)
}
