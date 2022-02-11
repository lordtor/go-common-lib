/*
 Base common func`s
*/

package go_common_lib

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	logging "github.com/lordtor/go-logging"
	jarger "github.com/lordtor/go-trace-lib"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var (
	Log = logging.Log
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

// A function for update exist slice not contain item string
// The return value is new slice
func UpdateStingSlice(ctx context.Context, slice []string, item string) []string {
	_, span := jarger.NewSpan(ctx, "Lib.UpdateStingSlice", nil)
	defer span.End()
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

// A function for update exist slice values from anothe slice
// The return value is updated slice
func UpdateStingSliceFromStingSlice(ctx context.Context, main []string, in []string) []string {
	_, span := jarger.NewSpan(ctx, "Lib.GetStringHash", nil)
	defer span.End()
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

// A function for calculate string hash sum by sha256
// The return value is hash string
func GetStringHash(ctx context.Context, aStr string) string {
	_, span := jarger.NewSpan(ctx, "Lib.GetStringHash", nil)
	defer span.End()
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

func SortedSliceMapString(ctx context.Context, m []map[string]string) []map[string]string {
	_, span := jarger.NewSpan(ctx, "Lib.SortedSliceMapString", nil)
	defer span.End()
	var keys []string
	res := []map[string]string{}
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
					res = append(res, map[string]string{k: v})
					break
				}
			}

		}
	}
	return res
}

// A function for convert interface to slice strings
// The return value is slice strings
func ListConvert(in interface{}) ([]string, error) {
	var out []string
	r, err := json.Marshal(in)
	if err != nil {
		Log.Error(err)
		return nil, err
	}
	err = json.Unmarshal(r, &out)
	if err != nil {
		Log.Error(err)
		return nil, err
	}
	return out, nil
}

func ConvertInterfaceToSliceString(ctx context.Context, t interface{}) []string {
	_, span := jarger.NewSpan(ctx, "Lib.ConvertInterfaceToSliceString", nil)
	defer span.End()
	r, err := json.Marshal(t)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(1, err.Error())
		Log.Error(err)
	}
	gr := []string{}
	err = json.Unmarshal(r, &gr)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(1, err.Error())
		Log.Error(err)
	}
	if err != nil {
		Log.Error(err)
	}
	return gr
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

// A function for update slince map by not exist map
// The return value is updated slice
func UpdateSliceMapStringFromMapString(ctx context.Context, sliceMaps []map[string]string, itemMap map[string]string) []map[string]string {
	_, span := jarger.NewSpan(ctx, "Lib.UpdateSliceMapStringFromMapString", nil)
	defer span.End()
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

// A function for update slice string by splited values use slice spliters string
// The return value is slice strings
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

// A function for convert interface to slice map
// The return value is slice map or error
func MapListConvert(in interface{}) ([]map[string]string, error) {
	var out []map[string]string
	rawData, err := json.Marshal(in)
	if err != nil {
		Log.Error(err)
		return nil, err
	}
	err = json.Unmarshal(rawData, &out)
	if err != nil {
		Log.Error(err)
		return nil, err
	}
	return out, nil
}

func ConvertInterfaceToSliceMapString(ctx context.Context, in interface{}) ([]map[string]string, error) {
	_, span := jarger.NewSpan(ctx, "Lib.ConvertInterfaceToSliceMapString", nil)
	defer span.End()
	var out []map[string]string
	rawData, err := json.Marshal(in)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(1, err.Error())
		Log.Error(err)
		return nil, err
	}
	err = json.Unmarshal(rawData, &out)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(1, err.Error())
		Log.Error(err)
		return nil, err
	}
	return out, nil
}

func SpanSetAttribute(sp trace.Span, prefix string, attributs interface{}) {
	var inInterface map[string]interface{}
	rawData, err := json.Marshal(attributs)
	if err != nil {
		Log.Error(err)
		return
	}
	err = json.Unmarshal(rawData, &inInterface)
	if err != nil {
		Log.Error(err)
		return
	}
	for field, val := range inInterface {
		switch v := val.(type) {
		case int:
			sp.SetAttributes(attribute.Key(fmt.Sprintf("%s%s", prefix, field)).Int(v))
		case float64:
			sp.SetAttributes(attribute.Key(fmt.Sprintf("%s%s", prefix, field)).Float64(v))
		case string:
			sp.SetAttributes(attribute.Key(fmt.Sprintf("%s%s", prefix, field)).String(v))
		case bool:
			sp.SetAttributes(attribute.Key(fmt.Sprintf("%s%s", prefix, field)).Bool(v))
		case []interface{}:
			ss := []string{}
			for vv := range v {
				ss = append(ss, fmt.Sprint(v[vv]))
			}
			sp.SetAttributes(attribute.Key(fmt.Sprintf("%s%s", prefix, field)).StringSlice(ss))
		default:
			rawData, err = json.Marshal(v)
			if err != nil {
				Log.Error(err)
			} else {
				sp.SetAttributes(attribute.Key(fmt.Sprintf("%s%s", prefix, field)).String(string(rawData)))
			}
		}
	}
}

func DecodeInterfaceToInterface(ctx context.Context, data interface{}, out interface{}) error {
	_, span := jarger.NewSpan(ctx, "Lib.DecodeIn", nil)
	defer span.End()
	raw, err := json.Marshal(data)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(1, err.Error())
		return err
	}
	err = json.Unmarshal(raw, &out)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(1, err.Error())
		return err
	}
	return nil
}
