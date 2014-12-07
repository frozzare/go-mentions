package mentions

import "regexp"

var r, _ = regexp.Compile(`(\@\w+\b(?:\/\w+|))`)

func Get(str string) []string {
	return r.FindAllString(str, -1)
}

type replace func(string) string

func Replace(str string, fn replace) string {
	return r.ReplaceAllStringFunc(str, fn)
}
