package files

import (
	"testing"
)

func TestChange(t *testing.T) {
	dest := "/Users/mtorres/Pictures/Wedding Pics/final"
	path := "/Users/mtorres/Pictures/Wedding Pics/album-d63798248-downloads-pt1"
	Change(path, dest)
	path = "/Users/mtorres/Pictures/Wedding Pics/album-d63798248-downloads-pt2"
	Change(path, dest)
}
