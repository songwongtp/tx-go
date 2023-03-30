package tx

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/osmosis-labs/osmosis/v15/app"
)

type Signer struct {
	ClientCtx   client.Context
	Txf         tx.Factory
	Key         keyring.Info
	AccSequence uint64
}

func NewSigner() *Signer {
	RPC_PATH := os.Getenv("RPC_PATH")
	CHAIN_ID := os.Getenv("CHAIN_ID")
	MNEMONIC := os.Getenv("MNEMONIC")
	GAS_PRICES := os.Getenv("GAS_PRICES")

	// Setup keyring
	kb := keyring.NewInMemory()
	path := sdk.GetConfig().GetFullBIP44Path()
	key, err := kb.NewAccount("alice", MNEMONIC, "", path, hd.Secp256k1)
	if err != nil {
		panic(err)
	}
	println("Account:", key.GetAddress().String())

	// Create client
	clientNode, err := client.NewClientFromNode(RPC_PATH)
	if err != nil {
		panic(err)
	}
	clientCtx := client.Context{
		Client:            clientNode,
		ChainID:           CHAIN_ID,
		NodeURI:           RPC_PATH,
		InterfaceRegistry: app.MakeEncodingConfig().InterfaceRegistry,
		TxConfig:          app.MakeEncodingConfig().TxConfig,
		Keyring:           kb,
	}

	// Retrieve account info
	accountRetriever := authtypes.AccountRetriever{}
	acc, err := accountRetriever.GetAccount(clientCtx, key.GetAddress())
	if err != nil {
		panic(err)
	}

	// Create transaction factory
	txf := tx.Factory{}.
		WithKeybase(kb).
		WithTxConfig(app.MakeEncodingConfig().TxConfig).
		WithAccountRetriever(clientCtx.AccountRetriever).
		WithAccountNumber(acc.GetAccountNumber()).
		WithGas(2000000).
		WithGasAdjustment(1.5).
		WithGasPrices(GAS_PRICES).
		WithChainID(CHAIN_ID).
		WithMemo("").
		WithSignMode(signing.SignMode_SIGN_MODE_DIRECT)

	return &Signer{
		ClientCtx:   clientCtx,
		Key:         key,
		Txf:         txf,
		AccSequence: acc.GetSequence(),
	}
}

func (signer *Signer) GetAddress() sdk.AccAddress {
	return signer.Key.GetAddress()
}

func (signer *Signer) IncrementSequence() {
	signer.AccSequence += 1
}

func (signer *Signer) SendTx(msg sdk.Msg) *sdk.TxResponse {
	txf := signer.Txf.WithSequence(signer.AccSequence)
	defer signer.IncrementSequence()

	_, adjustedGas, err := tx.CalculateGas(signer.ClientCtx, txf, msg)
	if err == nil {
		txf = txf.WithGas(adjustedGas)
	} else {
		println("simulate", err.Error())
	}

	txb, err := tx.BuildUnsignedTx(txf, msg)
	if err != nil {
		panic(err)
	}

	err = tx.Sign(txf, signer.Key.GetName(), txb, true)
	if err != nil {
		panic(err)
	}

	txBytes, err := signer.ClientCtx.TxConfig.TxEncoder()(txb.GetTx())
	if err != nil {
		panic(err)
	}

	res, err := signer.ClientCtx.BroadcastTxCommit(txBytes)
	if err != nil {
		panic(err)
	}

	println("\t", res.TxHash)
	return res
}
