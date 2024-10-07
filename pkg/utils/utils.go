package utils

import "strings"

// Bytes32ToString converts a byte32 Solidity array to a string.
func Bytes32ToString(b [32]byte) string {
	return strings.TrimRight(string(b[:]), "\x00")
}
