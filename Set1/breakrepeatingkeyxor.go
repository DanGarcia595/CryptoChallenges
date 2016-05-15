package Set1

func Hammingdistance(firstBlock, secondBlock []byte) (numBits int) {
	var tmpByte byte
	for index, _ := range firstBlock {
		for tmpByte = firstBlock[index] ^ secondBlock[index]; tmpByte > 0; tmpByte >>= 1 {
			numBits += int(tmpByte & byte(1))
		}
	}
	return
}
