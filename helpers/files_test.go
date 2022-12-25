package helpers

import "testing"

func TestSanitizer(t *testing.T) {

	cases := []struct {
		arg      string
		expected string
	}{
		{"pages/posts.md", "Posts"},
		{"pages/posts-1.md", "Posts 1"},
		{"pages/posts_2.md", "Posts 2"},
		{"pages/posts@3.md", "Posts 3"},
		{"pages/posts_blog@4.md", "Posts blog 4"},
		{"pages/posts_blog@/4.md", "4"},
	}

    for _, tc := range cases {
        actual,_ := Sanitize(tc.arg)
        if actual != tc.expected {
            t.Errorf("Expected %s, got %s", tc.expected, actual)
        }
    }

}
