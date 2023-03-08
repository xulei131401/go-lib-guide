package tool

import (
	"os"
	"text/template"

	"holy-cmd/hctl/types"
)

func Generate(genContext *types.GenContext) error {
	dir := genContext.Dir
	fileName := genContext.FileName
	templateName := genContext.TemplateName
	tpl := genContext.TplText
	spec := genContext.Spec

	err := os.MkdirAll(dir, 0o755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(dir+fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	t := template.New(templateName)
	t, err = t.Parse(tpl)
	if err != nil {
		return err
	}

	return t.Execute(file, spec)
}
