package types

const (
	// ModuleName defines the module name
	ModuleName = "flares"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"

	VoycherKey = "VOUCHER"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ContractKey         = "Contract"
	ContractReceiverKey = "Receiver"
	ContractTransferKey = "Transfer"
	SuccessContractKey  = "SuccessContract"
	FailContractKey     = "FailContract"
	BoardKey            = "Board"
	VoucherDenom        = "VOUCHER"
)

const (
	ContractTransferRecordKey = "ContractTransferRecord"
)
