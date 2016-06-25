package breadcrumb

import (
	"html/template"
)

var TemplateFuncs = template.FuncMap {
  "breadcrumb": RenderBreadCrumb,
}
