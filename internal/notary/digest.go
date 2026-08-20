package notary

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strings"
)

func NormalizeDigest(value, algorithm string) (string, error) {
	value = strings.ToLower(strings.TrimSpace(value))
	algorithm = strings.ToUpper(strings.TrimSpace(algorithm))
	if algorithm != "SHA-256" && algorithm != "SHA-512" {
		return "", ErrInvalidAlgorithm
	}
	size := 64
	if algorithm == "SHA-512" {
		size = 128
	}
	if len(value) != size {
		return "", ErrInvalidDigest
	}
	if _, err := hex.DecodeString(value); err != nil {
		return "", ErrInvalidDigest
	}
	return value, nil
}
func HashBytes(data []byte, algorithm string) string {
	if strings.EqualFold(algorithm, "SHA-512") {
		h := sha512.Sum512(data)
		return hex.EncodeToString(h[:])
	}
	h := sha256.Sum256(data)
	return hex.EncodeToString(h[:])
}
func HashStrings(values ...string) string {
	joined := strings.Join(values, "|")
	return HashBytes([]byte(joined), "SHA-256")
}
func DigestLabel(algorithm string) string {
	return fmt.Sprintf("%s:%s", strings.ToLower(algorithm), algorithm)
}
