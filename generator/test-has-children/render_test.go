package testhaschildren

import (
	"testing"

	_ "embed"

	"github.com/a-h/templ/generator/htmldiff"
)

//go:embed expected.html
var expected string

func Test(t *testing.T) {
	// fmt.Println("Epxected HTML")
	// fmt.Println(expected)
	// fmt.Println("-------------")
	render := render([]string{"Foo", "Bar", "Deez"})
	// fmt.Println("Actual HTML")
	// render.Render(context.Background(), os.Stdout)
	// fmt.Println("")
	// fmt.Println("-------------")

	diff, err := htmldiff.Diff(render, expected)
	if err != nil {
		t.Fatal(err)
	}
	if diff != "" {
		t.Error(diff)
	}
}
