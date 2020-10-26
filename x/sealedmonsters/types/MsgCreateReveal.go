package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgReveal{}

func NewMsgReveal(creator sdk.AccAddress, solutionHash string, solution string, scavenger string) *MsgReveal {
  return &MsgReveal{
    Id: uuid.New().String(),
		Creator: creator,
    SolutionHash: solutionHash,
    Solution: solution,
    Scavenger: scavenger,
	}
}

func (msg *MsgReveal) Route() string {
  return RouterKey
}

func (msg *MsgReveal) Type() string {
  return "CreateReveal"
}

func (msg *MsgReveal) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgReveal) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgReveal) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
