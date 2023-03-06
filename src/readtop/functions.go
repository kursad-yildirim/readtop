package readtop

import (
	"bufio"
	"fmt"
	"strings"

	"redhat.com/ess/esstools"
)

func (top *myTop) NextTop(topLineScanner *bufio.Scanner) bool {
	topHeaderIdentify := "top - "
	top.HasFocus = false
	notEndOfFile := true
	top.text = nil
	notEndOfTop := true
	top.text = append(top.text, nextTopFirstLine)

	for notEndOfFile = true; notEndOfFile && notEndOfTop; notEndOfFile = topLineScanner.Scan() {
		notEndOfTop = !(strings.Contains(topLineScanner.Text(), topHeaderIdentify) && len(top.text) > 1)
		top.text = append(top.text, topLineScanner.Text())
		top.HasFocus = top.HasFocus || strings.Contains(topLineScanner.Text(), tokenToLook) && tokenToLook != ""
	}
	return notEndOfFile
}

func (top myTop) PrintTop(maxLines int) {
	for line := 0; line < esstools.MinInteger(len(top.text)-2, maxLines); line++ {
		fmt.Println(top.text[line])
	}
	nextTopFirstLine = top.text[len(top.text)-1]
	fmt.Println("=================================================================================")
}
