// +build windows plan9

package hellcat

import (
	"os"
)

// FileIDs attempts to retrieve user and group IDs.
func FileIDs(fi os.FileInfo) (string, error) {
	return "", nil
}
