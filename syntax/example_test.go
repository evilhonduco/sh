// Copyright (c) 2016, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package syntax_test

import (
	"os"
	"strings"

	"github.com/mvdan/sh/syntax"
)

func Example() {
	in := strings.NewReader("{ foo; bar; }")
	f, err := syntax.Parse(in, "", 0)
	if err != nil {
		return
	}
	syntax.Fprint(os.Stdout, f)
	// Output:
	// {
	//	foo
	//	bar
	// }
}
