package pathx

import (
	"io/ioutil"
)

func LoadTemplate(file string) (string, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
