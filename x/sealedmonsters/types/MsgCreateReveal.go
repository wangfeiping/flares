package types

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgReveal{}

func NewMsgReveal(creator sdk.AccAddress, solution string) *MsgReveal {
	solutionHash := sha256.Sum256([]byte(solution))

	return &MsgReveal{
		Id:           uuid.New().String(),
		Creator:      creator,
		SolutionHash: hex.EncodeToString(solutionHash[:]),
		Solution:     solution,
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
	if msg.SolutionHash == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "solutionScavengerHash can't be empty")
	}
	if msg.Solution == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "solutionHash can't be empty")
	}

	var solutionHash = sha256.Sum256([]byte(msg.Solution))
	var solutionHashString = hex.EncodeToString(solutionHash[:])

	if msg.SolutionHash != solutionHashString {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Hash of solution (%s) doesn't equal solutionHash (%s)", msg.SolutionHash, solutionHashString))
	}

	return nil
}
