package profilefuncs

import (
	"encoding/json"
)

func ParseDesc(text string) string {
	var data struct {
		Description string `json:"description"`
	}

	if err := json.Unmarshal([]byte(text), &data); err != nil {
		return ""
	}
	/* parts := strings.Split(text, `"description":`)
	if len(parts) < 2 {
		return ""
	}

	desc := strings.Split(parts[1], "}")
	if len(desc) < 1 {
		return ""
	}

	return desc[0] */

	return data.Description
}
