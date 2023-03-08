package api

import (
	"holy-cmd/hctl/tool"
	"holy-cmd/hctl/tool/pathx"
	"holy-cmd/hctl/types"
	"log"
)

func Execute() {
	var err error
	context := &types.GenContext{
		Dir:          "/Users/xulei/jungle/golangworkspace/holy-cmd/_tmp/",
		FileName:     "api.go",
		TemplateName: "api",
	}

	var spec = make(map[string]string)
	spec["Name"] = "Teacher"
	context.Spec = spec

	var tpl string
	tplPath := "/Users/xulei/jungle/golangworkspace/holy-cmd/hctl/api/apigen/api.tpl"
	tpl, err = pathx.LoadTemplate(tplPath)
	if err != nil {
		panic(err)
	}

	context.TplText = tpl
	err = tool.Generate(context)
	if err != nil {
		log.Println("err:", err)
	} else {
		log.Println("成功")
	}

}
