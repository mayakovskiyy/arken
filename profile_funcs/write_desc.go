package profilefuncs

import (
	"context"
	"fmt"

	// "net/http"
	// "time"

	arken "arken/lexicons/api/arken"

	"io"

	// "github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/api/atproto"
	lexutil "github.com/bluesky-social/indigo/lex/util"
	// "github.com/bluesky-social/indigo/lex/util"
)

type Profile struct {
	*arken.ArkenProfile
}

func (a *Profile) MarshalCBOR(w io.Writer) error {
	data, err := a.ArkenProfile.MarshalCBOR()
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

func WriteDescription(ctx context.Context, description string, handle string, did string, c lexutil.LexClient) {
	profileData := &arken.ArkenProfile{
		Description: description,
		Handle:      handle,
	}
	pDataAdapted := &Profile{ArkenProfile: profileData}
	decodedData := &lexutil.LexiconTypeDecoder{Val: pDataAdapted}
	_, err := atproto.RepoPutRecord(ctx, c, &atproto.RepoPutRecord_Input{
		Collection: "fun.arken.profile",
		Repo:       did,
		Rkey:       "self",
		Record:     decodedData,
	})
	if err != nil {
		fmt.Println(err)
	}
}
