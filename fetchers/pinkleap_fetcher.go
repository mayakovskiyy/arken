package fetchers

import (
	"context"
	"fmt"
	"strings"

	// atproto "github.com/bluesky-social/indigo/api/atproto"
	// appbsky "github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/api/agnostic"
	"github.com/bluesky-social/indigo/lex/util"
)

func FetchPinkleap(handle string, c util.LexClient, did string) string {
	ctx := context.Background()

	var accountUrl string

	ans, err := agnostic.RepoGetRecord(ctx, c, "", "social.pinkleap.declaration", did, "self")
	if err != nil {
		fmt.Println(err)

		if strings.Contains(err.Error(), "RecordNotFound") {
			fmt.Printf("404 arken: not found\n")
		}
	}
	/* if ans == nil {
	fmt.Println("arken: empty tangled record")
	return ""
	} */

	if ans != nil {
		accountUrl = fmt.Sprintf("https://pinkleap.social/@%s\n", handle)
	}
	return accountUrl
}
