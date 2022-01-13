package readdir

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// WalkLineFunc is the function to get the line from reading.
type WalkLineFunc func(line string)

// ReadDir reads each text file in the directory and return the line.
func ReadDir(root string, fn WalkLineFunc) error {
	return filepath.Walk(root, func(path string, fi fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		return ReadFile(path, fn)
	})
}

// ReadFile returns each line in the file.
func ReadFile(path string, fn func(line string)) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		line := sc.Text()
		line = strings.TrimRight(line, "\n")
		fn(line)
	}
	return sc.Err()
}
