package Set1

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
