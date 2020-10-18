package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgWhois{}

func NewMsgWhois(creator sdk.AccAddress, value string, owner string, price string) *MsgWhois {
  return &MsgWhois{
    Id: uuid.New().String(),
		Creator: creator,
    Value: value,
    Owner: owner,
    Price: price,
	}
}

func (msg *MsgWhois) Route() string {
  return RouterKey
}

func (msg *MsgWhois) Type() string {
  return "CreateWhois"
}

func (msg *MsgWhois) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgWhois) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgWhois) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
