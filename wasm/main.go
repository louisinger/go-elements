package main

import (
	"encoding/hex"
	"syscall/js"

	"github.com/btcsuite/btcd/btcec"
	"github.com/vulpemventures/go-elements/pegin"
)

// Main function: it sets up our Wasm application
func main() {
	// Define the function in the JavaScript scope
	js.Global().Set("getPeginAddress", GetPeginAddressWrapper())
	// Prevent the function from returning, which is required in a wasm module
	select {}
}

// GetPeginAddressWrapper returns the javascript bind for pegin/GetAddressInfo function
func GetPeginAddressWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), h2b(args[0].String()))
		fedPegInfo := args[1].String()
		net := pegin.NetworkType(args[2].Int())
		isDynaFed := args[3].Bool()
		contract := h2b(args[4].String())
		pegin.GetAddressInfo(
			*privKey,
			fedPegInfo,
			net,
			isDynaFed,
			contract,
		)
		return map[string]interface{}{
			"hello":  "world",
			"answer": 42,
		}
	})
}

func b2h(buf []byte) string {
	return hex.EncodeToString(buf)
}

func h2b(str string) []byte {
	buf, _ := hex.DecodeString(str)
	return buf
}
