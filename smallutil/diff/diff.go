package diff

import (
	"io/ioutil"
	"strings"
)

func CompareStr(oldLines, newLines string) (addLines, delLines []string) {
	a1 := strings.Split(oldLines, "\n")
	a2 := strings.Split(newLines, "\n")

	inA1 := make(map[string]bool)
	for _, v := range a1 {
		inA1[v] = true
	}
	inA2 := make(map[string]bool)
	for _, v := range a2 {
		inA2[v] = true
	}

	for _, v := range a2 {
		if !inA1[v] {
			addLines = append(addLines, v)
		}
	}

	for _, v := range a1 {
		if !inA2[v] {
			delLines = append(delLines, v)
		}
	}

	return addLines, delLines
}

func CompareFile(oldFile, newFile string) (addLines, delLines []string) {
	buf1, err := ioutil.ReadFile(oldFile)
	if err != nil {
		panic(err)
	}

	buf2, err := ioutil.ReadFile(newFile)
	if err != nil {
		panic(err)
	}

	return CompareStr(string(buf1), string(buf2))
}
