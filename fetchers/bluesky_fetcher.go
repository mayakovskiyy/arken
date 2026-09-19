package fetchers

import (
	"context"
	"fmt"

	atproto "github.com/bluesky-social/indigo/api/atproto"
	_ "github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/lex/util"
)

func FetchBluesky(handle string, c util.LexClient, did string) string {
	ctx := context.Background()

	ans, err := atproto.RepoGetRecord(ctx, c, "", "app.bsky.actor.profile", did, "self")
	if err != nil {
		fmt.Println(err)
	}

	if ans == nil {
		fmt.Println("arken: empty bluesky record")
		return ""
	}
	accountUrl := fmt.Sprintf("https://bsky.app/profile/%s\n", handle)

	return accountUrl
}
