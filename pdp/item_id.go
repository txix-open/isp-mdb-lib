package pdp

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func GenerateItemIdForSingle(ssoId string, blockName string) string {
	byteValue := []byte(fmt.Sprintf("%s.%s", ssoId, blockName))
	hashValue := md5.Sum(byteValue)
	return strings.ToUpper(hex.EncodeToString(hashValue[:]))
}
