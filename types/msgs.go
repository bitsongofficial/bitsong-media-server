package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace sdk.CodespaceType = "bitsong-media-server"

	CodeInvalidFromAddress sdk.CodeType = 0
	CodeInvalidFileHash    sdk.CodeType = 1
	CodeInvalid            sdk.CodeType = 2

	TypeMsgUpload    = "upload"
	TypeMsgGetTrack  = "get_track"
	TypeMsgGetTracks = "get_tracks"
	TypeMsgEditTrack = "edit_track"
)

var _ sdk.Msg = MsgUpload{}

type MsgUpload struct {
	FromAddress sdk.AccAddress `json:"from_address"`
	FileHash    string         `json:"file_hash"`
	TrackId     string         `json:"track_id"`
}

func (msg MsgUpload) Route() string { return TypeMsgUpload }
func (msg MsgUpload) Type() string  { return TypeMsgUpload }
func (msg MsgUpload) ValidateBasic() sdk.Error {
	if msg.FromAddress.Empty() {
		return sdk.NewError(DefaultCodespace, CodeInvalidFromAddress, "from address cannot be blank")
	}

	if msg.FileHash == "" {
		return sdk.NewError(DefaultCodespace, CodeInvalidFileHash, "file hash cannot be blank")
	}

	return nil
}

func (msg MsgUpload) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgUpload) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress}
}

var _ sdk.Msg = MsgEditTrack{}

type MsgEditTrack struct {
	FromAddress          sdk.AccAddress `json:"from_address"`
	TrackId              string         `json:"track_id"`
	Title                string         `json:"title"`
	Artists              string         `json:"artists"`
	Featurings           *string        `json:"featurings,omitempty"`
	Producers            *string        `json:"producers,omitempty"`
	Genre                string         `json:"genre"`
	Mood                 string         `json:"mood"`
	ReleaseDate          string         `json:"release_date"`
	ReleaseDatePrecision string         `json:"release_date_precision"`
	Tags                 *string        `json:"tags,omitempty"`
	Explicit             bool           `json:"explicit"`
	Label                *string        `json:"label,omitempty"`
	Isrc                 *string        `json:"isrc,omitempty"`
	UpcEan               *string        `json:"upc_ean,omitempty"`
	Iswc                 *string        `json:"iswc,omitempty"`
	Credits              *string        `json:"credits,omitempty"`
	Copyright            string         `json:"copyright"`
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

	return nil
}

func (msg MsgEditTrack) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgEditTrack) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress}
}

var _ sdk.Msg = MsgGetTracks{}

type MsgGetTracks struct {
	FromAddress sdk.AccAddress `json:"from_address"`
}

func (msg MsgGetTracks) Route() string { return TypeMsgGetTracks }
func (msg MsgGetTracks) Type() string  { return TypeMsgGetTracks }
func (msg MsgGetTracks) ValidateBasic() sdk.Error {
	if msg.FromAddress.Empty() {
		return sdk.NewError(DefaultCodespace, CodeInvalidFromAddress, "from address cannot be blank")
	}

	return nil
}

func (msg MsgGetTracks) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgGetTracks) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress}
}

var _ sdk.Msg = MsgGetTrack{}

type MsgGetTrack struct {
	TrackId string `json:"track_id"`
}

func (msg MsgGetTrack) Route() string { return TypeMsgGetTrack }
func (msg MsgGetTrack) Type() string  { return TypeMsgGetTrack }
func (msg MsgGetTrack) ValidateBasic() sdk.Error {
	if msg.TrackId == "" {
		return sdk.NewError(DefaultCodespace, CodeInvalidFromAddress, "track id cannot be blank")
	}

	return nil
}

func (msg MsgGetTrack) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgGetTrack) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}
