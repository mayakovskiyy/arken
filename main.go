package main

import (
	"context"
	"fmt"

	// "net/http"
	// "time"

	fetchers "arken/fetchers"
	// arken "arken/lexicons/api/arken"
	pfuncs "arken/profile_funcs"

	// "io"

	// "github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/atproto/atclient"
	"github.com/bluesky-social/indigo/atproto/identity"
	"github.com/bluesky-social/indigo/atproto/syntax"
	// lexutil "github.com/bluesky-social/indigo/lex/util"
	// "github.com/bluesky-social/indigo/lex/util"
)

func main() {
	ctx := context.Background()

	var handle string
	var password string
	var description string

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

	profile := pfuncs.FetchProfile(client, did)

	if profile != true {
		fmt.Scan(&description)
		pfuncs.WriteDescription(ctx, description, handle, did, client)
	}

	profileData := pfuncs.GetDescription(did, client, ctx)
	fmt.Println(profileData)

	tangledProfile := fetchers.FetchTangled(handle, client, did)
	if tangledProfile != "" {
		fmt.Printf("Tangled profile: %s", tangledProfile)
	}
	blueskyProfile := fetchers.FetchBluesky(handle, client, did)
	if blueskyProfile != "" {
		fmt.Printf("Bluesky profile: %s", blueskyProfile)
	}
	whitewindProfile := fetchers.FetchWhiteWind(handle, client, did)
	if whitewindProfile != "" {
		fmt.Printf("WhiteWind profile: %s", whitewindProfile)
	}
	streamplaceProfile := fetchers.FetchStreamplace(handle, client, did)
	if streamplaceProfile != "" {
		fmt.Printf("Streamplace profile: %s", streamplaceProfile)
	}
}
