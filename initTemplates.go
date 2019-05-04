package cmgen

import "github.com/gobuffalo/packr/v2"

var TemplateBoxes *packr.Box

func init()  {
	TemplateBoxes = packr.New("tmplFiles", "./template")
}
