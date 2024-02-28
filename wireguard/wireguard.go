package wireguard

import (
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func GenerateWireguardKey() wgtypes.Key {
	key, error := wgtypes.GeneratePrivateKey()
	if error != nil {
		panic(error)
	}
	return key
}
