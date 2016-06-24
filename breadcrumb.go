package breadcrumb

import (
	"html/template"
	"fmt"
)

type BreadCrumbItem interface {
	Text() string
	Url()  string
}

func NewItem(text, url string) BreadCrumbItem {
  return &breadCrumbItem{text: text, url: url}
}

type breadCrumbItem struct {
	text string
	url  string
}

func (bci *breadCrumbItem) Text() string {
  return bci.text
}

func (bci *breadCrumbItem) Url() string {
  return bci.url
}

func New(baseUrl, separator string, items []BreadCrumbItem) *BreadCrumb {
  return &BreadCrumb{
    BaseUrl: baseUrl,
  	Items: items,
  	Separator: separator,
  }
}

type BreadCrumb struct {
	BaseUrl   string
	Items     []BreadCrumbItem
	Separator string
}

func (bc *BreadCrumb) renderItem(index int, separator string) string {
	return fmt.Sprintf(`
		<li itemprop="itemListElement" itemscope itemtype="http://schema.org/ListItem">
			%s<a itemprop="item" href="%s"><span itemprop="name">%s</span></a>
			<meta itemprop="position" content="%d" />
		</li>`, separator, bc.Items[index].Url(), bc.Items[index].Text(), index+1)
}

func (bc *BreadCrumb) Render() string {
	var html string
	noOfItems := len(bc.Items)
	if noOfItems > 0 {
		html += "<ol itemscope itemtype=\"http://schema.org/BreadcrumbList\">"
		html += bc.renderItem(0, "")
		var i int
		for i = 1; i < noOfItems; i++ {
			html += bc.renderItem(i, bc.Separator)
		}
		html += "</ol>"
	}
	return html
}

func RenderBreadCrumb(breabcrumb *BreadCrumb) template.HTML {
	return template.HTML(breabcrumb.Render())
}
