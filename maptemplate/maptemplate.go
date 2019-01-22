package maptemplate

import (
	"bufio"
	"bytes"
	"strings"
	"text/template"

	"github.com/drgrib/maps/maptemplate/mastertemplate"
)

type MapTemplate struct {
	Package,
	KeySuffix,
	ValueSuffix,
	KeyType,
	ValueType string
}

func NewMapTemplate(packageStr, keytype, valuetype string) MapTemplate {
	mapTemplate := MapTemplate{
		Package:     packageStr,
		KeyType:     keytype,
		ValueType:   valuetype,
		KeySuffix:   strings.Title(keytype),
		ValueSuffix: strings.Title(valuetype),
	}
	return mapTemplate
}

func (mapTemplate MapTemplate) Parse() (string, error) {
	masterTextTemplate, err := template.New("master").Parse(mastertemplate.String)
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	err = masterTextTemplate.Execute(writer, mapTemplate)
	if err != nil {
		return "", err
	}

	err = writer.Flush()
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
