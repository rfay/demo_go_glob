package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	matches, _ := checkForFileInSubdirs("./somedir", "junk.txt")
	fmt.Printf("Matches for 'junk.txt: %v\n", matches)

	matches, _ = checkForFileInSubdirs(".somedir", "lookingnotfinding.txt")
	fmt.Printf("Matches for 'lookingnotfinding.txt: %v\n", matches)

}

// checkForFileInSubdirs returns an array of found files in subdir
func checkForFileInSubdirs(pathWithSubdirs string, fileName string) ([]string, error) {
	glob := pathWithSubdirs + "/*/" + fileName
	matches, err := filepath.Glob(glob)
	return matches, err
}