package fetchers

import (
	"context"
	"fmt"
	"strings"

	atproto "github.com/bluesky-social/indigo/api/atproto"
	// appbsky "github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/lex/util"
)

func FetchTangled(handle string, c util.LexClient, did string) string {
	ctx := context.Background()

	var accountUrl string

	ans, err := atproto.RepoGetRecord(ctx, c, "", "sh.tangled.actor.profile", did, "self")
	if err != nil {
		fmt.Println(err)
	}

	if strings.Contains(err.Error(), "not found") {
		fmt.Printf("404 arken: not found\n %s", err)
	}

	if ans == nil {
		fmt.Println("arken: empty tangled record")
		return ""
	}

	if ans != nil {
		accountUrl = fmt.Sprintf("https://tangled.sh/%s", handle)
	}
	/* pData, ok := ans.Value.Val.(*appbsky.ActorProfile)
	if !ok {
		fmt.Println("arken: record is not a profile")
		} */

	return accountUrl
}
