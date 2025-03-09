package util

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateName(prefix ...string) string {
	n, _ := rand.Int(rand.Reader, big.NewInt(5))
	length := 6 + int(n.Int64())

	result := make([]byte, length)
	for i := range result {
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		result[i] = letters[idx.Int64()]
	}

	var sb strings.Builder
	for _, p := range prefix {
		sb.WriteString(p)
	}
	sb.Write(result)

	return sb.String()
}
