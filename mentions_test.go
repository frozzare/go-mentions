package mentions

import "github.com/bmizerany/assert"
import "strings"
import "testing"

func TestGet(t *testing.T) {
	res := Get("hello @frozzare @mariastarck")
	assert.Equal(t, res[0], "@frozzare")
	assert.Equal(t, res[1], "@mariastarck")
}

func TestReplace(t *testing.T) {
	res := Replace("hello @frozzare @mariastarck", func(mention string) string {
		return "<a href='http://twitter.com/" + mention[1:] + "'>" + mention + "</a>"
	})
	assert.Equal(t, true, strings.Contains(res, "<a href='http://twitter.com/frozzare'>@frozzare</a>"))
	assert.Equal(t, true, strings.Contains(res, "<a href='http://twitter.com/mariastarck'>@mariastarck</a>"))
}
