package testhaschildren

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	_ "embed"

	"github.com/a-h/templ"
	"github.com/a-h/templ/generator/htmldiff"
)

//go:embed expected.html
var expected string

func Test(t *testing.T) {
	fmt.Println(expected)
	contents := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, "<div>Inserted from Go code</div>")
		return err
	})
	render := render([]string{"Foo", "Bar", "Deez"})
	fmt.Println("--- pre testing ---")
	list().Render(templ.WithChildren(context.Background(), contents), os.Stdout)
	fmt.Println("")
	list().Render(context.Background(), os.Stdout)
	fmt.Println("")
	fmt.Println("--- post testing ---")

	diff, err := htmldiff.Diff(render, expected)
	if err != nil {
		t.Fatal(err)
	}
	if diff != "" {
		t.Error(diff)
	}
}
