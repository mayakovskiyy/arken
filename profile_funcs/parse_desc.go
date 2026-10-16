package profilefuncs

import (
	"encoding/json"
	"github.com/bluesky-social/indigo/api/agnostic"
)

func ParseDesc(text *agnostic.RepoGetRecord_Output) string {
	var data map[string]any
	jsn, _ := json.Marshal(text.Value)

	_ = json.Unmarshal(jsn, &data)

	desc := data["description"].(string)

	/* parts := strings.Split(text, `"description":`)
	if len(parts) < 2 {
		return ""
	}

	desc := strings.Split(parts[1], "}")
	if len(desc) < 1 {
		return ""
	}

	return desc[0] */

	return desc
}
