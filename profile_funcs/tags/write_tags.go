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

type Tags struct {
	*arken.ArkenTags
}

func (a *Tags) MarshalCBOR(w io.Writer) error {
	data, err := a.ArkenTags.MarshalCBOR()
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

func WriteTags(ctx context.Context, did string, c lexutil.LexClient, tag1 string, tag2 string, tag3 string, tag4 string, tag5 string) {
	tagsData := &arken.ArkenTags{
		Tag_main:           tag1,
		Tag_hobby:          tag2,
		Tag_fandom1:        tag3,
		Tag_fandom2:        tag4,
		Tag_minor_activity: tag5,
	}
	tDataAdapted := &Tags{ArkenTags: tagsData}
	decodedData := &lexutil.LexiconTypeDecoder{Val: tDataAdapted}
	_, err := atproto.RepoPutRecord(ctx, c, &atproto.RepoPutRecord_Input{
		Collection: "fun.arken.tags",
		Repo:       did,
		Rkey:       "self",
		Record:     decodedData,
	})
	if err != nil {
		fmt.Println(err)
	}
}
