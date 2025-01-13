package helpers

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
			r := regexp.MustCompile(`metadata:\n\s+resourceVersion: ""`)
			str = stripStr(str, "metadata:\nresourceVersion: \"\"", r)
			str = stripStr(str, "items:", nil)
			str = stripStr(str, "kind: List", nil)
			for _, s := range strings.Split(str, "\n") {

				if strings.Contains(s, "apiVersion:") {
					continue
				}

				each += strings.TrimPrefix(s, "  ") + "\n"
			}
			found := true
			beforeStr := ""
			for found {
				beforeStr, found = strings.CutSuffix(each, "\n")
				if found {
					each = beforeStr
				}
			}
			for {
				if strings.HasSuffix(each, "\n") {
					each = strings.TrimSuffix(each, "\n")
				} else if strings.HasPrefix(each, " ") {
					each = strings.TrimPrefix(each, " ")
				} else if strings.HasSuffix(each, " ") {
					each = strings.TrimSuffix(each, " ")
				} else {
					break
				}
			}
			if len(each) > 0 && !strings.HasPrefix(each, "apiVersion") {
				each = "apiVersion: " + each
			}
			if each != "" {
				result = append(result, each)
			}

		}
	} else {
		result = strArr
	}

	return result
}
func GetYamlKindName(y string) (string, string, string, error) {
	var r models.KubeResource
	err := yaml.Unmarshal([]byte(y), &r)
	if err != nil {
		return "", "", "", err
	}
	if r.Kind == "" {
		return "", "", "", fmt.Errorf("could not find 'kind' in yaml")
	}
	if r.Metadata.Name == "" {
		return "", "", "", fmt.Errorf("could not find 'metadata.name' in yaml")
	}
	return strings.ToLower(r.Kind), r.Metadata.Name, r.Metadata.Namespace, nil
}
func stripStr(str string, strip string, regex *regexp.Regexp) string {

	found := false
	if regex != nil {
		found = regex.Match([]byte(str))
		if found {
			s := regex.ReplaceAll([]byte(str), []byte{})
			str = string(s)
		}
	} else if strings.Contains(str, strip) { //if this is not prepended with indentation, it's a list outside of the actual resource yaml that we want
		startIndex := strings.Index(str, strip)
		prefix := str[:startIndex]
		suffix := str[len(prefix)+len(strip):]
		prefix = strings.TrimSuffix(prefix, " ")
		return prefix + suffix
	}
	return str
}

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
