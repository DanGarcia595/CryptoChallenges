package Set1

import "errors"

func PKCS7padding(data []byte, blocksize int) (finalData []byte) {
	remainingBytes := len(data) % blocksize
	if remainingBytes == 0 {
		finalData = data
		return
	}
	bytesToPad := blocksize - remainingBytes
	paddedData := make([]byte, len(data)+bytesToPad)
	for index, element := range data {
		paddedData[index] = element
	}
	for i := 0; i < bytesToPad; i++ {
		paddedData[len(data)+i] = 0x04 //pad
	}
	finalData = paddedData
	return
}

func PKCS7paddingStrip(data []byte, blocksize int) (string, error) {
	strippedData := ""
	if len(data)%blocksize != 0 {
		return "", errors.New("bad pkcs7 padding")
	}
	for i := 0; i < len(data); i++ {
		element := data[len(data)-i-1]
		if element >= 0x20 && element <= 0x7E { //found the text
			return string(data[:len(data)-i]), nil
		} else if element != 0x04 {
			return "", errors.New("bad pkcs7 padding")
		}
	}
	//should not happen, but error anyway
	return string(strippedData), errors.New("bad pkcs7 padding")
}
