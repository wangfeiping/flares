package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgContract{}

func NewMsgContract(creator sdk.AccAddress,
	module, key, receiver, accept string, durationHeight int32, bottom string) *MsgContract {
	return &MsgContract{
		Id:             uuid.New().String(),
		Creator:        creator,
		Module:         module,
		Key:            key,
		Receiver:       receiver,
		Accept:         accept,
		DurationHeight: durationHeight,
		Bottom:         bottom,
	}
}

func (msg *MsgContract) Route() string {
	return RouterKey
}

func (msg *MsgContract) Type() string {
	return "CreateContract"
}

func (msg *MsgContract) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgContract) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
