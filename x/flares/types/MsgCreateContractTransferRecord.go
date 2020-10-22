package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgContractTransferRecord{}

func NewMsgContractTransferRecord(from string, to string, amount string) *MsgContractTransferRecord {
	return &MsgContractTransferRecord{
		From:   from,
		To:     to,
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
	return []sdk.AccAddress{}
}

func (msg *MsgContractTransferRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgContractTransferRecord) ValidateBasic() error {
	return nil
}
