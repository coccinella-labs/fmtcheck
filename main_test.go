package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCheckDirFindsUnformatted(t *testing.T) {
	dir := t.TempDir()
	good := filepath.Join(dir, "good.go")
	if err := os.WriteFile(good, []byte("package main\n\nfunc main() {}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	bad := filepath.Join(dir, "bad.go")
	if err := os.WriteFile(bad, []byte("package main\nfunc  main(  ){}\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	unformatted, err := checkDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(unformatted) != 1 || unformatted[0] != bad {
		t.Fatalf("expected %s, got %v", bad, unformatted)
	}
}

func TestCheckDirIgnoresNonGoFiles(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "README.md"), []byte("# hi\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	unformatted, err := checkDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(unformatted) != 0 {
		t.Fatalf("expected no files, got %v", unformatted)
	}
}
