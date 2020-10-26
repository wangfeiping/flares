package types

import (
	"crypto/sha256"
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgMonster{}

func NewMsgMonster(creator sdk.AccAddress,
	description string, solution string, reward string) *MsgMonster {
	var solutionHash = sha256.Sum256([]byte(solution))
	var solutionHashString = hex.EncodeToString(solutionHash[:])
	return &MsgMonster{
		Id:           uuid.New().String(),
		Creator:      creator,
		Description:  description,
		SolutionHash: solutionHashString,
		Reward:       reward,
		Solution:     "******",
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
