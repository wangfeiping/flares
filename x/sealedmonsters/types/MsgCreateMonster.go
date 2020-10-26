package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgMonster{}

func NewMsgMonster(creator sdk.AccAddress, description string, solutionHash string, reward string, solution string, scavenger string) *MsgMonster {
  return &MsgMonster{
    Id: uuid.New().String(),
		Creator: creator,
    Description: description,
    SolutionHash: solutionHash,
    Reward: reward,
    Solution: solution,
    Scavenger: scavenger,
	}
}

func (msg *MsgMonster) Route() string {
  return RouterKey
}

func (msg *MsgMonster) Type() string {
  return "CreateMonster"
}

func (msg *MsgMonster) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgMonster) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgMonster) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
