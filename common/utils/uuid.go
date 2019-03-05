package slice_utils

import (
	"math/rand"
	"hex"
)

func GenerateUniqueID() string {
	id := make([]byte, 16)
	rand.Read(id)
	return hex.EncodeToString(id)
}