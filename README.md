# go-mentions [![Build Status](https://travis-ci.org/frozzare/go-mentions.svg?branch=master)](https://travis-ci.org/frozzare/go-mentions)

Get user mentions from string or replace user mentions with something else.

View the [docs](http://godoc.org/github.com/frozzare/go-mentions).

## Installation

```
$ go get github.com/frozzare/go-mentions
```

## Example

```go
package main

import "github.com/frozzare/go-mentions"

func main() {
	mentions.Get("hello @frozzare")
	//=> ["@frozzare"]
	
	mentions.Replace("hello @frozzare", func(mention string) string {
		return "<a href='http://twitter.com/" + mention[1:] + "'>" + mention + "</a>"
	})
	//=> hello <a href='http://twitter.com/frozzare'>@frozzare</a>
}
```

# License

MIT © [Fredrik Forsmo](https://github.com/frozzare)
