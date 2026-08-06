// Command fmtcheck checks Go source files against gofmt formatting.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path := flag.String("path", ".", "directory to check")
	flag.Parse()

	unformatted, err := checkDir(*path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fmtcheck: %v\n", err)
		os.Exit(1)
	}
	if len(unformatted) > 0 {
		fmt.Fprintf(os.Stderr, "fmtcheck: %d file(s) not formatted:\n", len(unformatted))
		for _, file := range unformatted {
			fmt.Fprintf(os.Stderr, "  %s\n", file)
		}
		os.Exit(1)
	}
	fmt.Println("fmtcheck: all files are formatted")
}

// checkDir returns the Go files under root that are not gofmt-formatted.
func checkDir(root string) ([]string, error) {
	var unformatted []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if d.Name() == ".git" || d.Name() == "vendor" {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		differs, err := isUnformatted(path)
		if err != nil {
			return err
		}
		if differs {
			unformatted = append(unformatted, path)
		}
		return nil
	})
	return unformatted, err
}

func isUnformatted(path string) (bool, error) {
	src, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}
	formatted, err := format.Source(src)
	if err != nil {
		return false, fmt.Errorf("%s: %w", path, err)
	}
	return !bytes.Equal(formatted, src), nil
}
