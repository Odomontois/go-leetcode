package hard

type Encrypter struct {
	encryption [26][2]byte
	decryption map[string]int
}

const aCode byte = byte('a')

func Constructor(keys []byte, values []string, dictionary []string) (enc Encrypter) {
	weights := make(map[[2]byte]int)

	for i, c := range keys {
		value := (*[2]byte)([]byte(values[i]))
		enc.encryption[c-aCode] = *value
		weights[*value]++
	}
	enc.decryption = make(map[string]int)
	for _, w := range dictionary {
		enc.decryption[enc.Encrypt(w)]++
	}

	return
}

func (enc *Encrypter) Encrypt(w string) string {
	res := make([]byte, 0, len(w)*2)
	for _, c := range []byte(w) {
		res = append(res, enc.encryption[c-aCode][:]...)
	}
	return string(res)
}

func (enc *Encrypter) Decrypt(w string) int {
	return enc.decryption[w]
}
