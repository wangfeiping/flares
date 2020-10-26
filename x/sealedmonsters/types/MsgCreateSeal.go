package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgSeal{}

func NewMsgSeal(creator sdk.AccAddress, solutionHash string, solutionScavengerHash string, scavenger string) *MsgSeal {
  return &MsgSeal{
    Id: uuid.New().String(),
		Creator: creator,
    SolutionHash: solutionHash,
    SolutionScavengerHash: solutionScavengerHash,
    Scavenger: scavenger,
	}
}

func (msg *MsgSeal) Route() string {
  return RouterKey
}

func (msg *MsgSeal) Type() string {
  return "CreateSeal"
}

func (msg *MsgSeal) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgSeal) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgSeal) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
