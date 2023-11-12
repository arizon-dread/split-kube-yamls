package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/arizon-dread/split-kube-yamls/models"
	"gopkg.in/yaml.v2"
)

func ReadYamlFileToStringArr(fn string) ([]string, error) {
	b, err := os.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	str := string(b)
	strArr := splitStr(str)
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
		result = append(result, "  apiVersion:"+str)
	}
	return result
}
func GetYamlKindName(y string) (string, string, error) {
	var r models.KubeResource
	err := yaml.Unmarshal([]byte(y), &r)
	if err != nil {
		fmt.Printf("couldn't unmarshal yaml, %v\n", err)
		return "", "", err
	}
	if r.Kind == "" {
		return "", "", fmt.Errorf("could not find 'kind' in yaml")
	}
	if r.Metadata.Name == "" {
		return "", "", fmt.Errorf("could not find 'metadata.name' in yaml")
	}
	return r.Kind, r.Metadata.Name, nil
}

func GetKindAndNameFromYaml(y string) (string, string) {
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

}

func ReadStdin() []string {
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

func WriteOutput(c string, fn string) error {
	return os.WriteFile(fn, []byte(c), 0644)
}
