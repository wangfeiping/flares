package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgContractTransferRecord{}

func NewMsgContractTransferRecord(creator sdk.AccAddress, hash string, from string, to string, amount string) *MsgContractTransferRecord {
  return &MsgContractTransferRecord{
    Id: uuid.New().String(),
		Creator: creator,
    Hash: hash,
    From: from,
    To: to,
    Amount: amount,
	}
}

func (msg *MsgContractTransferRecord) Route() string {
  return RouterKey
}

func (msg *MsgContractTransferRecord) Type() string {
  return "CreateContractTransferRecord"
}

func (msg *MsgContractTransferRecord) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgContractTransferRecord) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgContractTransferRecord) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
