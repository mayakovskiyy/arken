// com.whtwnd.blog.entry
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

func FetchWhiteWind(handle string, c util.LexClient, did string) string {
	ctx := context.Background()

	var accountUrl string

	ans, err := agnostic.RepoGetRecord(ctx, c, "", "com.whtwnd.blog.entry", did, "self")
	if err != nil {
		fmt.Println(err)

		if strings.Contains(err.Error(), "RecordNotFound") {
			fmt.Printf("404 arken: not found\n%s", err)
		}
	}
	if ans == nil {
		fmt.Println("arken: empty whitewind record\n")
		return ""
	}

	if ans != nil {
		accountUrl = fmt.Sprintf("https://whtwnd.com/%s\n", handle)
	}
	/* pData, ok := ans.Value.Val.(*appbsky.ActorProfile)
	if !ok {
		fmt.Println("arken: record is not a profile")
		} */

	return accountUrl
}
