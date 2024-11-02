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
	var result []string
	if len(strArr) == 1 {
		if strings.Contains(s, "- apiVersion:") {
			strArr = strings.Split(s, "- apiVersion:")
		}
		for _, str := range strArr {
			var each string
			if strings.Contains(str, "kind: List") {
				continue
			}
			for _, s := range strings.Split(str, "\n") {
				each += strings.TrimPrefix(s, "  ") + "\n"
			}
			result = append(result, "apiVersion:"+each)
		}
	} else {
		result = strArr
	}

	return result
}
func GetYamlKindName(y string) (string, string, error) {
	var r models.KubeResource
	err := yaml.Unmarshal([]byte(y), &r)
	if err != nil {
		return "", "", err
	}
	if r.Kind == "" {
		return "", "", fmt.Errorf("could not find 'kind' in yaml")
	}
	if r.Metadata.Name == "" {
		return "", "", fmt.Errorf("could not find 'metadata.name' in yaml")
	}
	return strings.ToLower(r.Kind), r.Metadata.Name, nil
}

//func GetKindAndNameFromYaml(y string) (string, string) {
//	kind := ""
//	name := ""
//
//	s := bufio.NewScanner(strings.NewReader(y))
//	for s.Scan() {
//		if s.Text() == "" {
//			break
//		}
//		s := s.Text()
//		b, a, found := strings.Cut(s, ":")
//		if found && strings.Trim(strings.ToLower(b), " ") == "kind" {
//			kind = strings.Trim(a, " ")
//		}
//		if found && strings.Trim(strings.ToLower(b), " ") == "name" {
//			name = strings.Trim(a, " ")
//		}
//		if kind != "" && name != "" {
//			break
//		}
//	}
//
//	return kind, name
//
//}

func ReadStdin() []string {
	s := bufio.NewScanner(os.Stdin)
	var str string
	var l []string
	for s.Scan() {
		str += "\n" + s.Text()

	}

	l = splitStr(str)
	return l
}

func WriteOutput(c string, fn string) error {
	return os.WriteFile(fn, []byte(c), 0644)
}
