package encoding

import (
	"strings"
)

const allowedChars string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var base62Map []string

func init() {
	base62Map = strings.Split(allowedChars, "")
}

func Base62Encode(inputNum uint64) string {
	var numsToChars []string
	for inputNum > 0 {
		numsToChars = append(
			numsToChars,
			base62Map[inputNum%62],
		)
		inputNum /= 62
	}
	return strings.Join(numsToChars, "")
}
