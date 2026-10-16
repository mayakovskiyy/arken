package profilefuncs

import (
	"context"
	"fmt"
	// "strings"

	// atproto "github.com/bluesky-social/indigo/api/atproto"
	// appbsky "github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/api/agnostic"
	"github.com/bluesky-social/indigo/lex/util"
)

func GetDescription(did string, c util.LexClient, ctx context.Context) *agnostic.RepoGetRecord_Output {
	ans, err := agnostic.RepoGetRecord(ctx, c, "", "fun.arken.profile", did, "self")
	if err != nil {
		fmt.Println(err)
	}

	return ans
}
