package profilefuncs

import (
	"strings"
)

func ParseDesc(text string) string {
	parts := strings.Split(text, `"description":`)
	if len(parts) < 2 {
		return ""
	}

	desc := strings.Split(parts[1], "}")
	if len(desc) < 1 {
		return ""
	}

	return desc[0]
}
