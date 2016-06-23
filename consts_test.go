package breadcrumb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func TestSeparator(t *testing.T) {
  assert.Equal(t, "&nbsp;&#8250;&nbsp;", Separator)
}
