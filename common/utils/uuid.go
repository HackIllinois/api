package slice_utils

import (
	"math/rand"
	"encoding/hex"
)

func GenerateUniqueID() string {
	id := make([]byte, 16)
	rand.Read(id)
	return hex.EncodeToString(id)
}