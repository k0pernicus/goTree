package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func main() {
	maxDepth := flag.Int("maxDepth", 3, "Set the depth to plot the tree")
	dir := flag.String("dir", "./", "Set the repository to start the search")
	dirOnly := flag.Bool("dirOnly", false, "List directories only")
	flag.Parse()
	if err := list(*dir, *maxDepth, *dirOnly); err != nil {
		fmt.Println("Impossible to print the informations...")
	}
}

func splitFn(c rune) bool {
	return c == '/'
}

func depthOf(s string) int {
	return len(strings.FieldsFunc(s, splitFn))
}

func list(dir string, maxDepth int, dirOnly bool) error {
	dirDepth := depthOf(dir)
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		cDepth := depthOf(path) - (dirDepth + 1)
		if path == dir || cDepth > maxDepth {
			return nil
		}
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path: %q: %v\n", dir, err)
			return err
		}
		for i := 1; i <= cDepth; i++ {
			fmt.Printf("\t")
		}
		if info.IsDir() {
			color.Red("|> %+v\n", info.Name())
			return nil
		}
		if !dirOnly {
			color.Green("|-> %+v\n", info.Name())
		}
		return nil
	})
}
