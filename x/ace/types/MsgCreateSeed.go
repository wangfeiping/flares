package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgSeed{}

func NewMsgSeed(creator sdk.AccAddress, secret string) *MsgSeed {
  return &MsgSeed{
    Id: uuid.New().String(),
		Creator: creator,
    Secret: secret,
	}
}

func (msg *MsgSeed) Route() string {
  return RouterKey
}

func (msg *MsgSeed) Type() string {
  return "CreateSeed"
}

func (msg *MsgSeed) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgSeed) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgSeed) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
