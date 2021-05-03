package main

import (
	"encoding/hex"
	"syscall/js"

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
		key, err := h2b(args[0].String())
		if err != nil {
			return err
		}

		fedPegScript, err := h2b(args[1].String())
		if err != nil {
			return err
		}

		net := pegin.NetworkType(args[2].Int())
		isDynaFed := args[3].Bool()

		contract, err := h2b(args[4].String())
		if err != nil {
			return err
		}

		addressInfo, err := pegin.GetAddressInfo(
			key,
			pegin.FedpegInfo{
				FedpegScript: fedPegScript,
			},
			net,
			isDynaFed,
			contract,
		)

		if err != nil {
			return err
		}

		return addressInfo.ClaimScript
	})
}

func b2h(buf []byte) string {
	return hex.EncodeToString(buf)
}

func h2b(str string) ([]byte, error) {
	return hex.DecodeString(str)
}
