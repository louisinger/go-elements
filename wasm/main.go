package main

import (
	"encoding/hex"
	"syscall/js"

	"github.com/vulpemventures/go-elements/pegin"
)

// main binds go wrappers to js global scope functions
func main() {
	js.Global().Set("getPeginAddress", GetPeginAddressWrapper())

	select {} // prevents the function to stop
}

// GetPeginAddressWrapper returns the javascript bind function for pegin.GetAddressInfo
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

		return map[string]interface{}{
			"claimScript":      addressInfo.ClaimScript,
			"mainChainAddress": addressInfo.MainChainAddress,
		}
	})
}

// encodes bytes to hex
func b2h(buf []byte) string {
	return hex.EncodeToString(buf)
}

// decodes hex to bytes
func h2b(str string) ([]byte, error) {
	return hex.DecodeString(str)
}
