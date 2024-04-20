package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/nour_dev/blogposts"
)

type StubFailingFS struct{}

func (fs StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("always fail!!")
}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: title
Description: desc`
		secondBody = `Title: title
Description: desc`
	)
	fs := fstest.MapFS{
		"hello world.md": {Data: []byte(firstBody)},
		"world2.md":      {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	assertNoError(t, err)

	assertSameLength(t, posts, fs)

	got := posts[0]

	assertPost(t, got, blogposts.Post{
		Title:       "title",
		Description: "desc",
	})
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func assertSameLength(t *testing.T, posts []blogposts.Post, fs fstest.MapFS) {
	t.Helper()
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted  %d posts", len(posts), len(fs))
	}
}
