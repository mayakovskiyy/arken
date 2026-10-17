package profilefuncs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bluesky-social/indigo/api/agnostic"
	"github.com/bluesky-social/indigo/lex/util"
)

func FetchAvatar(did string, c util.LexClient, ctx context.Context) string {
	ans, err := agnostic.RepoGetRecord(ctx, c, "", "app.bsky.actor.profile", did, "self")
	if err != nil {
		fmt.Println(err)
	}

	var data map[string]any
	jsn, _ := json.Marshal(ans.Value)

	_ = json.Unmarshal(jsn, &data)

	avatarCid := data["avatar"].(map[string]any)["ref"].(map[string]any)["$link"].(string)

	return avatarCid
}
