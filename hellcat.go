package hellcat

import (
	"github.com/dustin/go-humanize"
	"github.com/meirf/gopart"

	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// Version is semver.
const Version = "0.0.2"

// Config parameterizes hellcat navigation.
type Config struct {
	// Working denotes the user's current working directory.
	Working string

	// Toplevels denotes a collection of file and/or directory paths to search.
	Toplevels []string

	// Examine enables additional metadata.
	// Default: false
	Examine bool

	// Recurse enables nested directory traversal.
	// Default: false
	Recurse bool
}

// Neighborhood calculates the number of nodes
// within 1 edge of a directed directory or file tree.
//
// Files count as one.
// Directories by themselves do not count,
// however self links (.) and parent links (..) each count as one.
func Neighborhood(pth string) (int, error) {
	fi, err := os.Stat(pth)

	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			return 1, nil
		}

		return -1, err
	}

	if fi.Mode().IsRegular() {
		return 1, nil
	}

	children, err := ioutil.ReadDir(pth)

	if err != nil {
		return -1, err
	}

	return 2 + len(children), nil
}

// Abbreviate truncates long strings.
func Abbreviate(s string, limit int) string {
	if len(s) <= limit {
		return s
	}

	return s[:(limit-2)/2] + "..." + s[len(s)-(limit-3)/2:]
}

// roamFile prints directory information.
func (o Config) roamDirectory(toplevel string, pth string) error {
	fis, err := ioutil.ReadDir(pth)

	if err != nil {
		if e, ok := err.(*os.PathError); ok {
			fmt.Fprintf(os.Stderr, "Error loading path: %v\n", e.Path)
			return nil
		}

		return err
	}

	for _, fi := range fis {
		pth2 := fi.Name()
		mode := fi.Mode()
		size := fi.Size()
		ts := fi.ModTime()
		sys := fi.Sys()

		if !strings.HasPrefix(pth2, "/") {
			pth2 = path.Join(pth, pth2)
		}

		pthRel, err2 := filepath.Rel(toplevel, pth2)

		if err2 != nil {
			return err2
		}

		pthRelClean := path.Clean(pthRel)

		var sysString string

		if sys != nil {
			ss, err3 := FileIDs(fi)

			if err3 != nil {
				return err3
			}

			sysString = ss
		}

		neighbors, err2 := Neighborhood(pthRelClean)

		if err2 != nil {
			fmt.Fprintf(os.Stderr, "Error loading path: %v\n", pthRelClean)
			neighbors = 1
		}

		neighborsSI, neighborsPrefix := humanize.ComputeSI(float64(neighbors))

		neighborsString := fmt.Sprintf("%v%v", neighborsSI, neighborsPrefix)

		sizeSI, sizePrefix := humanize.ComputeSI(float64(size))

		sizeString := fmt.Sprintf("%v%vB", int(sizeSI), sizePrefix)

		tsRFC3339 := ts.UTC().Format(time.RFC3339)

		var nameString string

		if mode&os.ModeSymlink != 0 {
			target, err2 := os.Readlink(pth2)

			if err2 != nil {
				return err2
			}

			nameString = fmt.Sprintf("%v -> %v", pthRelClean, target)
		} else {
			nameString = pthRelClean
		}

		fmt.Printf("%v %4v%v %4v %v %v\n", mode, neighborsString, sysString, sizeString, tsRFC3339, nameString)

		if o.Recurse {
			mode := fi.Mode()

			if mode.IsDir() {
				if err2 := o.roamDirectory(toplevel, pth2); err2 != nil {
					return err2
				}
			} else {
				if err2 := o.roamFile(pthRelClean); err2 != nil {
					return err2
				}
			}
		}
	}

	return nil
}

// roamFile prints file information.
func (o Config) roamFile(pth string) error {
	file, err := os.Open(pth)

	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			fmt.Println("(Missing)")
			return nil
		}

		return err
	}

	defer func() {
		if err2 := file.Close(); err2 != nil {
			fmt.Fprintln(os.Stderr, err2)
		}
	}()

	buf := make([]byte, 1024)

	count, err := file.Read(buf)

	if err != nil {
		if err != io.EOF {
			return err
		}

		return nil
	}

	mimetype := http.DetectContentType(buf[:count])
	isBinary := !strings.HasPrefix(mimetype, "text")

	for {
		if o.Examine || isBinary {
			for i := range gopart.Partition(count, 8) {
				var segments []string

				for _, b := range buf[i.Low:i.High] {
					segments = append(segments, fmt.Sprintf("%02x", b))
				}

				fmt.Printf("%08d %v\n", i.Low, strings.Join(segments, " "))
			}
		} else {
			if _, err2 := os.Stdout.Write(buf[:count]); err2 != nil {
				return err2
			}
		}

		count, err = file.Read(buf)

		if err != nil {
			if err != io.EOF {
				return err
			}

			break
		}
	}

	return nil
}

// Roam prints file system information.
func (o Config) Roam() error {
	for _, pth := range o.Toplevels {
		pth, err := filepath.Abs(pth)

		if err != nil {
			return err
		}

		var pth2 string

		if strings.HasPrefix(pth, "/") {
			pth2 = pth
		} else {
			pth2 = path.Join(o.Working, pth)
		}

		fi, err := os.Stat(pth2)

		if err != nil {
			if _, ok := err.(*os.PathError); ok {
				fmt.Println("(Missing)")
				continue
			}

			return err
		}

		mode := fi.Mode()

		if mode.IsDir() {
			if err2 := o.roamDirectory(pth, pth2); err2 != nil {
				return err2
			}
		} else if err2 := o.roamFile(pth2); err2 != nil {
			return err2
		}
	}

	return nil
}
