package go_common_lib

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func SliceContain(slice []string, item string) bool {
	exist := false
	for sliceID := range slice {
		if slice[sliceID] == item {
			exist = true
			break
		}
	}
	return exist
}
func UpdateList(slice []string, item string) []string {
	if !SliceContain(slice, item) {
		slice = append(slice, item)
	}
	return slice
}

func StringFromList(list []string) (st string) {
	for i := 0; i < len(list); i++ {
		if st == "" {
			st = fmt.Sprintf("%s", list[i])
		} else {
			st = fmt.Sprintf("%s,%s", st, list[i])
		}

	}
	return st
}
func UpdateStructList(main []string, in []string) []string {
	for inID := range in {
		id := inID
		main = UpdateList(main, in[id])
	}
	return main
}

func CalcHash(aStr string) string {
	hash := sha256.New()
	hash.Write([]byte(aStr))
	return hex.EncodeToString(hash.Sum(nil))
}
