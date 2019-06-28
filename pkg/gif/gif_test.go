package gif

import "testing"

func TestGif(t *testing.T) {
	ReverseFrames("cd.gif", "cd_rev.gif")
}
