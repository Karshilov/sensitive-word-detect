package utils

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"unicode"
)

func GetKeywords() ([]string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	raw, err := ioutil.ReadFile(path.Join(pwd, `/keywords.txt`))
	if err != nil {
		return nil, err
	}
	keywords := strings.FieldsFunc(string(raw), func(r rune) bool { return unicode.IsControl(r) || unicode.IsSpace(r) })
	return keywords, nil
}
