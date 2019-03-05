package slice_utils

import (
	"encoding/hex"
	"math/rand"
)

func GenerateUniqueID() string {
	id := make([]byte, 16)
	rand.Read(id)
	return hex.EncodeToString(id)
}
