package breadcrumb

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
  "html/template"
  "os"
)


func TestMain(m *testing.M) {
	fmt.Println("Test starting")
	retCode := m.Run()
	fmt.Println("Test ending")
	os.Exit(retCode)
}


func TestNewItem(t *testing.T) {
  item := NewItem("text", "url.tld")
  assert.Equal(t, "text", item.Text())
  assert.Equal(t, "url.tld", item.Url())
}

func TestBreadcrumb(t *testing.T) {
  expected := "<div xmlns:v=\"http://rdf.data-vocabulary.org/#\"><span typeof=\"v:Breadcrumb\"><a href=\"url.tld/parent\" rel=\"v:url\" property=\"v:title\">parent</a>><span rel=\"v:child\"><span typeof=\"v:Breadcrumb\"><a href=\"url.tld/child\" rel=\"v:url\" property=\"v:title\">child</a>>child2</span></span></div>"
  bc := New("", ">", []BreadCrumbItem{
    NewItem("parent", "url.tld/parent"),
    NewItem("child", "url.tld/child"),
    NewItem("child2", "url.tld/child2"),})
  assert.Equal(t, expected, bc.Render())
}

func TestRenderBreadCrumb(t *testing.T) {
  bc := New("", ">", []BreadCrumbItem{
    NewItem("parent", "url.tld/parent"),})
  assert.IsType(t, template.HTML(""), RenderBreadCrumb(bc))
}
