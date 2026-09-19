package main

import (
	"context"
	"fmt"

	// "net/http"
	// "time"

	fetchers "arken/fetchers"
	// arken "arken/lexicons/api/arken"
	// "io"

	// "github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/atproto/atclient"
	"github.com/bluesky-social/indigo/atproto/identity"
	"github.com/bluesky-social/indigo/atproto/syntax"
	// "github.com/bluesky-social/indigo/lex/util"
)

func main() {
	ctx := context.Background()

	var handle string
	var password string

	fmt.Scan(&handle)
	fmt.Scan(&password)

	client, _ := atclient.LoginWithPassword(ctx,
		identity.DefaultDirectory(),
		syntax.AtIdentifier(handle),
		password,
		"",
		nil,
	)

	did := fmt.Sprintf("%s", client.AccountDID)
	fmt.Printf("ok! %s\n", did)

	tangledProfile := fetchers.FetchTangled(handle, client, did)
	if tangledProfile != "" {
		fmt.Printf("Tangled profile: %s", tangledProfile)
	}
	blueskyProfile := fetchers.FetchBluesky(client, did)
	if blueskyProfile != "" {
		fmt.Printf("Bluesky profile: %s", blueskyProfile)
	}
}
