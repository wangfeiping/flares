package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgBoard{}

func NewMsgBoard(creator sdk.AccAddress, base string, baseDenom string, baseAddress string, accept string, acceptDenom string, acceptAddress string, source string) *MsgBoard {
  return &MsgBoard{
    Id: uuid.New().String(),
		Creator: creator,
    Base: base,
    BaseDenom: baseDenom,
    BaseAddress: baseAddress,
    Accept: accept,
    AcceptDenom: acceptDenom,
    AcceptAddress: acceptAddress,
    Source: source,
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
