package september

func shiftingLetters(s string, shifts []int) string {
	var b []byte
	ab := byte('a')
	k := 0
	var qs []byte
	for i := len(shifts) - 1; i >= 0; i-- {
		k = (k + shifts[i]) % 26
		qs = append(qs, byte(k))
	}

	for i, c := range s {
		b = append(b, (byte(c)-ab+qs[len(s)-i-1])%26+ab)
	}
	return string(b)
}
