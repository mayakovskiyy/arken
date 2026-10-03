package profilefuncs

import (
	"context"
	"fmt"
	"strings"
	// atproto "github.com/bluesky-social/indigo/api/atproto"
	// appbsky "github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/api/agnostic"
	"github.com/bluesky-social/indigo/lex/util"
)

func FetchProfile(c util.LexClient, did string) bool {
	ctx := context.Background()

	ans, err := agnostic.RepoGetRecord(ctx, c, "", "fun.arken.profile", did, "self")
	if err != nil {
		fmt.Println(err)

		if strings.Contains(err.Error(), "RecordNotFound") {
			fmt.Printf("404 arken: fun.arken.profile not found\n%s\n", err)
			return false
		}
	}

	if ans != nil {
		fmt.Println("fun.arken.profile found")
	}
	return true
}
