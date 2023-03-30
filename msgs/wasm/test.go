package wasm

import (
	"strconv"

	"txgo/tx"
)

func Test(signer *tx.Signer) {
	accAddress := signer.GetAddress().String()

	// MsgStoreCode
	println("Wasm/MsgStoreCode - Invalid File")
	println(signer.SendTx(NewStoreCodeInvalidFile(accAddress)).RawLog, "\n")

	println("Wasm/MsgStoreCode - Success")
	res := signer.SendTx(NewStoreCodeSuccess(accAddress))
	codeId, err := strconv.ParseUint(res.Logs[0].Events[1].Attributes[1].Value, 10, 64)
	if err != nil {
		panic(err)
	}
	println("Code ID:", codeId, "\n")

	// MsgInstantiateContract
	println("Wasm/MsgInstantiateContract - Invalid Code")
	println(signer.SendTx(NewInstantiateInvalidCode(accAddress)).RawLog, "\n")

	println("Wasm/MsgInstantiateContract - Success")
	res = signer.SendTx(NewInstantiateSuccess(accAddress, codeId))
	contractAddress0 := res.Logs[0].Events[0].Attributes[0].Value
	println("Contract Address:", contractAddress0, "\n")
	res = signer.SendTx(NewInstantiateSuccess(accAddress, codeId))
	contractAddress1 := res.Logs[0].Events[0].Attributes[0].Value
	println("Contract Address:", contractAddress1, "\n")

	// MsgInstantiateContract2
	println("Wasm/MsgInstantiateContract2")
	res = signer.SendTx(NewInstantiate2(accAddress, codeId))
	contractAddress2 := res.Logs[0].Events[0].Attributes[0].Value
	println("Contract Address:", contractAddress2, "\n")

	// MsgExecuteContract
	println("Wasm/MsgExecuteContract - Invalid Field")
	println(signer.SendTx(NewExecuteInvalidField(accAddress, contractAddress0)).RawLog, "\n")

	println("Wasm/MsgExecuteContract - Success")
	signer.SendTx(NewExecuteSuccess(accAddress, contractAddress0))
	println()

	// MsgMigrateContract
	println("Wasm/MsgMigrateContract - Invalid Code ID")
	println(signer.SendTx(NewMigrateInvalidCodeId(accAddress, contractAddress0)).RawLog, "\n")

	println("Wasm/MsgMigrateContract - Success")
	signer.SendTx(NewMigrateSuccess(accAddress, contractAddress0))
	println()

	// MsgUpdateAdmin
	println("Wasm/MsgUpdateAdmin - Invalid Address")
	println(signer.SendTx(NewUpdateAdminNotNewAddress(accAddress, contractAddress2)).RawLog, "\n")

	println("Wasm/MsgUpdateAdmin - Success")
	signer.SendTx(NewUpdateAdminSuccess(accAddress, contractAddress0))
	println()

	// MsgClearAdmin
	println("Wasm/MsgClearAdmin - No Permission")
	println(signer.SendTx(NewClearAdmin(accAddress, contractAddress0)).RawLog, "\n")

	println("Wasm/MsgClearAdmin - Success")
	signer.SendTx(NewClearAdmin(accAddress, contractAddress1))
	println()
}
