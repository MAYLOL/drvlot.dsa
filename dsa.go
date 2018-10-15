// h 20181015
//
// DrvLot.DSA

package dsa

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

func sig(pri, sha string) string {
	ret := ""
	for {
		_pri, err := crypto.HexToECDSA(pri)
		if err != nil {
			break
		}
		_sha, err := hex.DecodeString(sha)
		if err != nil {
			break
		}
		_sig, err := crypto.Sign(_sha, _pri)
		if err != nil {
			break
		}
		ret = hex.EncodeToString(_sig)
		//
		// Finally
		if true {
			break
		}
	}
	return ret
}
