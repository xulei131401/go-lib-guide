package types

type GenContext struct {
	Dir          string
	FileName     string
	TplPath      string
	TplText      string
	TemplateName string
	Spec         interface{}
}
