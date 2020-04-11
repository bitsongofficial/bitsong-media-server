package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

const TypeMsgEditTrack = "edit_track"

var _ sdk.Msg = MsgEditTrack{}

type MsgEditTrack struct {
	FromAddress          sdk.AccAddress `json:"from_address"`
	TrackId              string         `json:"track_id"`
	Title                string         `json:"title"`
	Artists              string         `json:"artists"`
	Featurings           string         `json:"featurings,omitempty"`
	Producers            string         `json:"producers,omitempty"`
	Genre                string         `json:"genre"`
	Mood                 string         `json:"mood"`
	ReleaseDate          string         `json:"release_date"`
	ReleaseDatePrecision string         `json:"release_date_precision"`
	Tags                 string         `json:"tags,omitempty"`
	Explicit             bool           `json:"explicit"`
	Label                string         `json:"label,omitempty"`
	Isrc                 string         `json:"isrc,omitempty"`
	UpcEan               string         `json:"upc_ean,omitempty"`
	Iswc                 string         `json:"iswc,omitempty"`
	Credits              string         `json:"credits,omitempty"`
	Copyright            string         `json:"copyright"`
	RewardsUsers         string         `json:"rewards_users"`
	RewardsPlaylists     string         `json:"rewards_playlists"`
	RightsHolders        string         `json:"rights_holders"`
	Visibility           string         `json:"visibility"`
}

func (msg MsgEditTrack) Route() string { return TypeMsgEditTrack }
func (msg MsgEditTrack) Type() string  { return TypeMsgEditTrack }
func (msg MsgEditTrack) ValidateBasic() sdk.Error {
	if msg.FromAddress.Empty() {
		return sdk.NewError(DefaultCodespace, CodeInvalidFromAddress, "from address cannot be blank")
	}

	if msg.TrackId == "" {
		return sdk.NewError(DefaultCodespace, CodeInvalid, "track id cannot be blank")
	}

	if msg.Title == "" {
		return sdk.NewError(DefaultCodespace, CodeInvalid, "title cannot be blank")
	}

	if msg.Artists == "" {
		return sdk.NewError(DefaultCodespace, CodeInvalid, "artists cannot be blank")
	}

	if msg.Genre == "" {
		return sdk.NewError(DefaultCodespace, CodeInvalid, "genre cannot be blank")
	}

	if msg.Mood == "" {
		return sdk.NewError(DefaultCodespace, CodeInvalid, "mood cannot be blank")
	}

	if msg.ReleaseDate == "" {
		return sdk.NewError(DefaultCodespace, CodeInvalid, "release_date cannot be blank")
	}
	// TODO: check release date is as date format YYYY-MM-DD
	// date, err := time.Parse("2006-01-02", msg.ReleaseDate)

	if msg.ReleaseDatePrecision == "" || msg.ReleaseDatePrecision != "Day" && msg.ReleaseDatePrecision != "Month" && msg.ReleaseDatePrecision != "Year" {
		return sdk.NewError(DefaultCodespace, CodeInvalid, "release_date_precision must be Day, Month or Year")
	}

	if msg.Copyright == "" || msg.Copyright != "RR" && msg.Copyright != "CC" {
		return sdk.NewError(DefaultCodespace, CodeInvalid, "copyright must be RR for Rights Reserved and CC for Creative Commons")
	}

	if msg.Visibility == "" || msg.Visibility != "public" && msg.Visibility != "private" {
		return sdk.NewError(DefaultCodespace, CodeInvalid, "visibility must be public or private")
	}

	rewardUsersFloat, err := strconv.ParseFloat(msg.RewardsUsers, 64)
	if err != nil {
		return sdk.NewError(DefaultCodespace, CodeInvalidFromAddress, "rewards users error")
	}

	if rewardUsersFloat <= 0 {
		return sdk.NewError(DefaultCodespace, CodeInvalidFromAddress, "rewards users must be > 0.01")
	}

	rewardPlaylistsFloat, err := strconv.ParseFloat(msg.RewardsPlaylists, 64)
	if err != nil {
		return sdk.NewError(DefaultCodespace, CodeInvalidFromAddress, "rewards playlists error")
	}

	if rewardPlaylistsFloat <= 0 {
		return sdk.NewError(DefaultCodespace, CodeInvalidFromAddress, "rewards playlists must be > 0.01")
	}

	rewardSum := rewardUsersFloat + rewardPlaylistsFloat
	if rewardSum > 100 {
		return sdk.NewError(DefaultCodespace, CodeInvalidFromAddress, "rewards sum must be less than 100")
	}

	if msg.RightsHolders == "" {
		return sdk.NewError(DefaultCodespace, CodeInvalidFromAddress, "rights holders cannot be blank")
	}
	// TODO: check rights_holders format bitsong1dsde3........:80,bitsong1dsde3........:20
	// TODO: check rights_holders quota sum equal to 100

	return nil
}

func (msg MsgEditTrack) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgEditTrack) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress}
}
