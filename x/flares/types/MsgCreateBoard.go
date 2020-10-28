package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgBoard{}

func NewMsgBoard(creator sdk.AccAddress,
	baseDenom string, acceptDenom string, source string) *MsgBoard {
	return &MsgBoard{
		Id:          uuid.New().String(),
		Creator:     creator,
		BaseDenom:   baseDenom,
		AcceptDenom: acceptDenom,
		Source:      source,
	}
}

func (msg *MsgBoard) Route() string {
	return RouterKey
}

func (msg *MsgBoard) Type() string {
	return "CreateBoard"
}

func (msg *MsgBoard) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgBoard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBoard) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
