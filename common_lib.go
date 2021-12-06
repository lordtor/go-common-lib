/*
 Base common func`s
*/

package go_common_lib

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// A function to check slice contain item string
// The return value is bool
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

// A function for update exist slice not contain item string
// The return value is new slice
func UpdateList(slice []string, item string) []string {
	if !SliceContain(slice, item) {
		slice = append(slice, item)
	}
	return slice
}

// A function convert slice to sting
// The return value is string
func StringFromList(list []string) (st string) {
	for i := 0; i < len(list); i++ {
		if st == "" {
			st = list[i]
		} else {
			st = fmt.Sprintf("%s,%s", st, list[i])
		}

	}
	return st
}

// A function for update exist slice values from anothe slice
// The return value is updated slice
func UpdateStructList(main []string, in []string) []string {
	for inID := range in {
		id := inID
		main = UpdateList(main, in[id])
	}
	return main
}

// A function for calculate string hash sum by sha256
// The return value is hash string
func CalcHash(aStr string) string {
	hash := sha256.New()
	hash.Write([]byte(aStr))
	return hex.EncodeToString(hash.Sum(nil))
}

// A function for sort slice map`s
// The return value is sorted slice
func SortedMap(m []map[string]string) (out []map[string]string) {
	var keys []string
	for i := range m {
		for k := range m[i] {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	for _, k := range keys {
		for i := range m {
			for kk, v := range m[i] {
				if k == kk {
					out = append(out, map[string]string{k: v})
					break
				}
			}
		}
	}
	return out
}

// A function for convert interface to slice strings
// The return value is slice strings
func ListConvert(in interface{}) ([]string, error) {
	var out []string
	r, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(r, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// A function for update slince map by not exist map
// The return value is updated slice
func UpdateMapList(sliceMaps []map[string]string, itemMap map[string]string) []map[string]string {
	update := false
	if len(sliceMaps) == 0 {
		sliceMaps = append(sliceMaps, itemMap)
	} else {
		for sliceID := range sliceMaps {
			for k, v := range sliceMaps[sliceID] {
				for kk, vv := range itemMap {
					if k != kk {
						sliceMaps = append(sliceMaps, itemMap)
						update = true
						break
					} else if k == kk && v != vv {
						sliceMaps[sliceID][k] = vv
						update = true
						break
					}
				}
				if update {
					break
				}
			}
			if update {
				break
			}
		}
	}
	return sliceMaps
}

// A function for trim first & last space
// The return value is trimed string
func StandardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
func UpdateListBySplit(slice []string, item string, spliters []string) []string {
	split := false
	char := ""
	if len(spliters) == 0 {
		spliters = []string{",", "\\", "/"}
	}

	for i := range spliters {
		if strings.Contains(item, spliters[i]) {
			split = true
			char = spliters[i]
		}
	}
	if split {
		words := strings.Split(item, char)
		if len(words) > 0 {
			for id := range words {
				w := StandardizeSpaces(words[id])
				if !SliceContain(slice, w) {
					slice = append(slice, w)
				}
			}
		}
	} else {
		if !SliceContain(slice, item) {
			w := StandardizeSpaces(item)
			slice = append(slice, w)
		}

	}
	return slice
}

func MapListConvert(in interface{}) ([]map[string]string, error) {
	var out []map[string]string
	r, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(r, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
