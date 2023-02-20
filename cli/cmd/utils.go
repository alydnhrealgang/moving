package cmd

import (
	"github.com/alydnhrealgang/moving/common/utils"
	"strings"
)

func ExcelIndex2Column(n int) string {
	col := utils.EmptyString
	for n > 0 {
		m := n % 26
		if m == 0 {
			m = 26
		}
		col = string(rune(m+64)) + col
		n = (n - m) / 26
	}
	return col
}

func ExcelColumn2Index(col string) int {
	if utils.EmptyOrWhiteSpace(col) {
		return 0
	}
	n := 0
	j := 1
	col = strings.ToUpper(col)
	for i := len(col) - 1; i >= 0; i-- {
		c := col[i]
		if c < 'A' || c > 'Z' {
			return 0
		}
		n += (int(c) - 64) * j
		j *= 26
	}
	return n
}
