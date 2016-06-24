package breadcrumb

import (
	"html/template"
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

func (bc *BreadCrumb) renderChild(index int) string {
	var html string
	html += "<span rel=\"v:child\">"
	html += "<span typeof=\"v:Breadcrumb\">"
	html += "<a href=\"" + bc.Items[index].Url() + "\" rel=\"v:url\" property=\"v:title\">"
	html += bc.Items[index].Text()
	html += "</a>"
	html += bc.Separator
	return html
}

func (bc *BreadCrumb) Render() string {
	var html string
	var endspans string
	noOfItems := len(bc.Items)

	if noOfItems > 0 {
		html += "<div xmlns:v=\"http://rdf.data-vocabulary.org/#\">"
		html += "<span typeof=\"v:Breadcrumb\">"
		html += "<a href=\"" + bc.Items[0].Url() + "\" rel=\"v:url\" property=\"v:title\">"
		html += bc.Items[0].Text()
		html += "</a>"

		html += bc.Separator
		var i int
		for i = 1; i < noOfItems-1; i++ {
			html += bc.renderChild(i)
			endspans += "</span></span>"
		}
		endspans += "</span>"
		if noOfItems > 1 {
			html += bc.Items[i].Text()
		}
		html += endspans
		html += "</div>"
	}
	return html
}

func RenderBreadCrumb(breabcrumb *BreadCrumb) template.HTML {
	return template.HTML(breabcrumb.Render())
}
