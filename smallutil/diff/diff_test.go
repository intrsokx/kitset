package diff

import (
	"strings"
	"testing"
)

func TestCompareStr(t *testing.T) {
	oldLines := `
kangxi
wangfan
`
	newLines := `
kangxi
wangfan
chenpeng
`
	addLines, delLines := CompareStr(oldLines, newLines)
	t.Logf("addLines: %v\n", addLines)
	t.Logf("delLines: %v\n", delLines)

	addLines, delLines = CompareStr(newLines, oldLines)
	t.Logf("addLines: %v\n", addLines)
	t.Logf("delLines: %v\n", delLines)
}

func TestCompareFile(t *testing.T) {
	addLines, delLines := CompareFile("old.txt", "new.txt")
	t.Logf("addLines: \n%v\n", strings.Join(addLines, "\n"))
	t.Logf("delLines: \n%v\n", strings.Join(delLines, "\n"))
}
