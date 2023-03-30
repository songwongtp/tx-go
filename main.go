package main

import (
	"fmt"
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joho/godotenv"

	"txgo/msgs/wasm"
	"txgo/tx"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	// Set prefix
	sdk.GetConfig().SetBech32PrefixForAccount(os.Getenv("PREFIX"), "")
}

func main() {
	signer := tx.NewSigner()
	wasm.Test(signer)
}
