package renderer

import (
	"fmt"
	"io"

	"github.com/nour_dev/blogposts/reader"
)

func Render(w io.Writer, post reader.Post) error {
	_, err := fmt.Fprintf(w, `<h1>%s</h1>`, post.Title)
	return err
}
