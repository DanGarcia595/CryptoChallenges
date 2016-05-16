package Set1

import (
	"encoding/base64"
	"encoding/hex"
)

func Hextobase64(s string) (ret string) {
	bytes, _ := hex.DecodeString(s)                //Decode the hex to bytes
	ret = base64.StdEncoding.EncodeToString(bytes) //Encode the bytes to base 64 string
	return
}
