package services

import "math/rand"

const CHARSET = "abcdfghjklmnpqrstvwxyz" +
	"ABCDFGHJKLMNPQRSTVWXYZ" +
	"0123456789"

func generateCode(codeLen int8) string {
	b := make([]byte, codeLen)

	for i := int8(0); i < codeLen; i++ {
		b[i] = CHARSET[rand.Intn(len(CHARSET))]
	}

	return string(b)
}
