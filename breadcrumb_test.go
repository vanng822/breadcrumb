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

func TestBreadcrumb2(t *testing.T) {
  expected := "<ol itemscope itemtype=\"http://schema.org/BreadcrumbList\">\n\t\t<li itemprop=\"itemListElement\" itemscope itemtype=\"http://schema.org/ListItem\">\n\t\t\t<a itemprop=\"item\" href=\"url.tld/parent\"><span itemprop=\"name\">parent</span></a>\n\t\t\t<meta itemprop=\"position\" content=\"1\" />\n\t\t</li>\n\t\t<li itemprop=\"itemListElement\" itemscope itemtype=\"http://schema.org/ListItem\">\n\t\t\t><a itemprop=\"item\" href=\"url.tld/child\"><span itemprop=\"name\">child</span></a>\n\t\t\t<meta itemprop=\"position\" content=\"2\" />\n\t\t</li>\n\t\t<li itemprop=\"itemListElement\" itemscope itemtype=\"http://schema.org/ListItem\">\n\t\t\t><a itemprop=\"item\" href=\"url.tld/child2\"><span itemprop=\"name\">child2</span></a>\n\t\t\t<meta itemprop=\"position\" content=\"3\" />\n\t\t</li></ol>"
  bc := New("", ">", []BreadCrumbItem{
    NewItem("parent", "url.tld/parent"),
    NewItem("child", "url.tld/child"),
    NewItem("child2", "url.tld/child2"),})
  assert.Equal(t, expected, bc.Render())
}

func TestBreadcrumb3(t *testing.T) {
  expected := "<ol itemscope itemtype=\"http://schema.org/BreadcrumbList\">\n\t\t<li itemprop=\"itemListElement\" itemscope itemtype=\"http://schema.org/ListItem\">\n\t\t\t<a itemprop=\"item\" href=\"url.tld/parent\"><span itemprop=\"name\">parent</span></a>\n\t\t\t<meta itemprop=\"position\" content=\"1\" />\n\t\t</li>\n\t\t<li itemprop=\"itemListElement\" itemscope itemtype=\"http://schema.org/ListItem\">\n\t\t\t><a itemprop=\"item\" href=\"url.tld/child\"><span itemprop=\"name\">child</span></a>\n\t\t\t<meta itemprop=\"position\" content=\"2\" />\n\t\t</li>\n\t\t<li itemprop=\"itemListElement\" itemscope itemtype=\"http://schema.org/ListItem\">\n\t\t\t><a itemprop=\"item\" href=\"url.tld/child2\"><span itemprop=\"name\">child2</span></a>\n\t\t\t<meta itemprop=\"position\" content=\"3\" />\n\t\t</li>\n\t\t<li itemprop=\"itemListElement\" itemscope itemtype=\"http://schema.org/ListItem\">\n\t\t\t><a itemprop=\"item\" href=\"url.tld/child3\"><span itemprop=\"name\">child3</span></a>\n\t\t\t<meta itemprop=\"position\" content=\"4\" />\n\t\t</li></ol>"
  bc := New("", ">", []BreadCrumbItem{
    NewItem("parent", "url.tld/parent"),
    NewItem("child", "url.tld/child"),
    NewItem("child2", "url.tld/child2"),
		NewItem("child3", "url.tld/child3"),})
  assert.Equal(t, expected, bc.Render())
}

func TestBreadCrumbSingle(t *testing.T) {
  bc := New("", ">", []BreadCrumbItem{
    NewItem("parent", "url.tld/parent"),})
	expected := "<ol itemscope itemtype=\"http://schema.org/BreadcrumbList\">\n\t\t<li itemprop=\"itemListElement\" itemscope itemtype=\"http://schema.org/ListItem\">\n\t\t\t<a itemprop=\"item\" href=\"url.tld/parent\"><span itemprop=\"name\">parent</span></a>\n\t\t\t<meta itemprop=\"position\" content=\"1\" />\n\t\t</li></ol>"
  assert.Equal(t, expected, bc.Render())
}

func TestRenderBreadCrumb(t *testing.T) {
  bc := New("", ">", []BreadCrumbItem{
    NewItem("parent", "url.tld/parent"),})
  assert.IsType(t, template.HTML(""), RenderBreadCrumb(bc))
}
