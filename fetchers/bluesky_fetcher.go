package fetchers

import (
	"context"
	"fmt"

	atproto "github.com/bluesky-social/indigo/api/atproto"
	// appbsky "github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/lex/util"
)

func FetchBluesky(c util.LexClient, did string) string {
	ctx := context.Background()

	util.RegisterType("app.bsky.actor.profile", &util.RawExtension{})

	ans, err := atproto.RepoGetRecord(ctx, c, "", "app.bsky.actor.profile", did, "self")
	if err != nil {
		fmt.Println(err)
	}

	if ans == nil {
		fmt.Println("arken: empty bluesky record")
		return ""
	}
	accountUrl := ans.Uri

	return accountUrl
}
