// +build !windows
// +build !plan9

package hellcat

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

// FileIDs attempts to retrieve user and group IDs.
func FileIDs(fi os.FileInfo) (string, error) {
	sys := fi.Sys()

	t, ok := sys.(*syscall.Stat_t)

	if !ok {
		return "", nil
	}

	uid := t.Uid
	gid := t.Gid

	uidString := strconv.FormatUint(uint64(uid), 10)
	gidString := strconv.FormatUint(uint64(gid), 10)

	u, err2 := user.LookupId(uidString)

	if err2 == nil {
		uidString = u.Username
	}

	g, err2 := user.LookupGroupId(gidString)

	if err2 == nil {
		gidString = g.Name
	}

	return fmt.Sprintf(" %8v %8v", Abbreviate(uidString, 8), Abbreviate(gidString, 8)), nil
}
