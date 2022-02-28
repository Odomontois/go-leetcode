package amzn

import (
	"regexp"
	"strings"
)

func fruitCodeRegex(code [][]string, cart []string) (bool, error) {
	enc := makeEncoderRE()
	enc.encodeCode(code)
	codeString := enc.Buf.String()
	enc.Buf.Reset()
	enc.encodeSlice(cart)
	basketString := enc.Buf.String()
	codeRegex, err := regexp.Compile(codeString)
	if err != nil {
		return false, err
	}
	return codeRegex.Match([]byte(basketString)), nil
}

type encoderRE struct {
	mapping map[string]int
	next    int
	Buf     strings.Builder
}

func makeEncoderRE() encoderRE {
	return encoderRE{mapping: make(map[string]int), next: 1000}
}

func (enc *encoderRE) encodeSlice(slice []string) {
	for _, s := range slice {
		if s == "anything" {
			enc.Buf.WriteString("\\E.\\Q")
			continue
		}
		code := enc.next
		if c, ok := enc.mapping[s]; ok {
			code = c
		} else {
			enc.mapping[s] = code
			enc.next++

		}
		enc.Buf.WriteRune(rune(code))
	}
}

func (enc *encoderRE) encodeCode(code [][]string) {
	for _, ss := range code {
		enc.Buf.WriteString(".*")
		enc.Buf.WriteString("\\Q")
		enc.encodeSlice(ss)
		enc.Buf.WriteString("\\E")
	}
	enc.Buf.WriteString(".*")
}
