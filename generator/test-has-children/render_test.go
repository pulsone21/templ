package testhaschildren

import (
	"fmt"
	"testing"

	_ "embed"

	"github.com/a-h/templ/generator/htmldiff"
)

//go:embed expected.html
var expected string

func Test(t *testing.T) {
	fmt.Println(expected)
	render := render([]string{"Foo", "Bar", "Deez"})

	diff, err := htmldiff.Diff(render, expected)
	if err != nil {
		t.Fatal(err)
	}
	if diff != "" {
		t.Error(diff)
	}
}
