package helpers

import "testing"

func TestSanitizer(t *testing.T) {

	cases := []struct {
		arg      string
		expected string
	}{
		{"content/posts.md", "Posts"},
		{"content/posts-1.md", "Posts 1"},
		{"content/posts_2.md", "Posts 2"},
		{"content/posts@3.md", "Posts 3"},
		{"content/posts_blog@4.md", "Posts blog 4"},
		{"content/posts_blog@/4.md", "4"},
	}

    for _, tc := range cases {
        actual,_ := Sanitize(tc.arg)
        if actual != tc.expected {
            t.Errorf("Expected %s, got %s", tc.expected, actual)
        }
    }

}
