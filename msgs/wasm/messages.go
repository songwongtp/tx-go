package wasm

import (
	"os"
	"time"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

func NewStoreCodeInvalidFile(sender string) *wasmtypes.MsgStoreCode {
	wasmFile, _ := os.ReadFile(TEXT_FILE_PATH)
	return &wasmtypes.MsgStoreCode{Sender: sender, WASMByteCode: wasmFile}
}

func NewStoreCodeSuccess(sender string) *wasmtypes.MsgStoreCode {
	wasmFile, _ := os.ReadFile(WASM_FILE_PATH)
	return &wasmtypes.MsgStoreCode{Sender: sender, WASMByteCode: wasmFile}
}

func NewInstantiateInvalidCode(sender string) *wasmtypes.MsgInstantiateContract {
	return &wasmtypes.MsgInstantiateContract{
		Sender: sender,
		CodeID: 1000000,
		Label:  "Contract Invalid Code",
		Msg:    []byte(InitMsg),
	}
}

func NewInstantiateSuccess(sender string, codeId uint64) *wasmtypes.MsgInstantiateContract {
	return &wasmtypes.MsgInstantiateContract{
		Sender: sender,
		Admin:  sender,
		CodeID: codeId,
		Label:  "Contract Instantiate Success",
		Msg:    []byte(InitMsg),
	}
}

func NewInstantiate2(sender string, codeId uint64) *wasmtypes.MsgInstantiateContract2 {
	return &wasmtypes.MsgInstantiateContract2{
		Sender: sender,
		CodeID: codeId,
		Label:  "Contract Instantiate2",
		Msg:    []byte(InitMsg),
		Salt:   []byte(time.Now().String()),
	}
}

func NewExecuteInvalidField(sender string, contractAddress string) *wasmtypes.MsgExecuteContract {
	return &wasmtypes.MsgExecuteContract{
		Sender:   sender,
		Contract: contractAddress,
		Msg:      []byte(ExecuteMsgInvalid),
	}
}

func NewExecuteSuccess(sender string, contractAddress string) *wasmtypes.MsgExecuteContract {
	return &wasmtypes.MsgExecuteContract{
		Sender:   sender,
		Contract: contractAddress,
		Msg:      []byte(ExecuteMsg),
	}
}

func NewMigrateInvalidCodeId(sender string, contractAddress string) *wasmtypes.MsgMigrateContract {
	return &wasmtypes.MsgMigrateContract{
		Sender:   sender,
		Contract: contractAddress,
		CodeID:   1,
		Msg:      []byte(MigrateMsg),
	}
}

func NewMigrateSuccess(sender string, contractAddress string) *wasmtypes.MsgMigrateContract {
	return &wasmtypes.MsgMigrateContract{
		Sender:   sender,
		Contract: contractAddress,
		CodeID:   MIGRATE_TO_CODE_ID,
		Msg:      []byte(MigrateMsg),
	}
}

func NewUpdateAdminNotNewAddress(sender string, contractAddress string) *wasmtypes.MsgUpdateAdmin {
	return &wasmtypes.MsgUpdateAdmin{
		Sender:   sender,
		NewAdmin: NEW_ADMIN_ADDRESS,
		Contract: contractAddress,
	}
}

func NewUpdateAdminSuccess(sender string, contractAddress string) *wasmtypes.MsgUpdateAdmin {
	return &wasmtypes.MsgUpdateAdmin{
		Sender:   sender,
		NewAdmin: NEW_ADMIN_ADDRESS,
		Contract: contractAddress,
	}
}

func NewClearAdmin(sender string, contractAddress string) *wasmtypes.MsgClearAdmin {
	return &wasmtypes.MsgClearAdmin{
		Sender:   sender,
		Contract: contractAddress,
	}
}
