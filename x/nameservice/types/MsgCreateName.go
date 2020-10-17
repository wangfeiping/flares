package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgName{}

func NewMsgName(creator sdk.AccAddress, value string, price string) *MsgName {
  return &MsgName{
    Id: uuid.New().String(),
		Creator: creator,
    Value: value,
    Price: price,
	}
}

func (msg *MsgName) Route() string {
  return RouterKey
}

func (msg *MsgName) Type() string {
  return "CreateName"
}

func (msg *MsgName) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgName) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgName) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
