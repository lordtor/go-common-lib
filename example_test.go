package go_common_lib_test

import (
	"fmt"

	. "github.com/lordtor/go-common-lib"
)

// Example check slice contain value
func ExampleSliceContain() {
	result := SliceContain([]string{"a", "b", "c"}, "b")
	fmt.Println("Contain:", result)
	// Output:
	// Contain: true
}

// Example update string slice add new value if not contain
func ExampleUpdateList() {
	result := UpdateList([]string{"a", "b", "c"}, "d")
	fmt.Println("New slice:", result)
	// Output:
	// New slice: [a b c d]
}

// Example convert slice to string
func ExampleStringFromList() {
	result := StringFromList([]string{"a", "b", "c"})
	fmt.Println("String:", result)
	// Output:
	// String: a,b,c
}

// Example update string slice by anothe string slice add not contain datas
func ExampleUpdateStructList() {
	result := UpdateStructList([]string{"a", "b"}, []string{"a", "c"})
	fmt.Println("Updated slice:", result)
	// Output:
	// Updated slice: [a b c]
}

// Example generate hash string for string data
func ExampleCalcHash() {
	result := CalcHash("testString")
	fmt.Println("Hash:", result)
	// Output:
	// Hash: 4acf0b39d9c4766709a3689f553ac01ab550545ffa4544dfc0b2cea82fba02a3
}

// Example sort slice map string by key
func ExampleSortedMap() {
	result := SortedMap([]map[string]string{{"b": "c", "c": "c", "a": "c"}})
	fmt.Println("Sorted map:", result)
	// Output:
	// Sorted map: [map[a:c] map[b:c] map[c:c]]
}

// Example convert interface to slice strings
func ExampleListConvert() {
	var in interface{}
	in = []string{"first", "second"}
	result, _ := ListConvert(in)
	fmt.Println("Slice:", result)
	// Output:
	// Slice: [first second]
}

// Example update map string slice by anothe add new data or value
func ExampleUpdateMapList() {
	result := UpdateMapList([]map[string]string{{"a": "b"}}, map[string]string{"b": "b"})
	fmt.Println("Slice:", result)
	// Output:
	// Slice: [map[a:b] map[b:b]]
}

// Example strip first and last spaces
func ExampleStandardizeSpaces() {
	result := StandardizeSpaces(" ggg ggg ggg ")
	fmt.Println("String:", result)
	// Output:
	// String: ggg ggg ggg
}

// Example update slice string values by splited data
func ExampleUpdateListBySplit() {
	result := UpdateListBySplit([]string{"A", "B"}, "B*C", []string{"*"})
	fmt.Println("String:", result)
	// Output:
	// String: [A B C]
}

// Example  convert slice map string to string
func ExampleMapListConvert() {
	var names interface{}
	names = []map[string]string{{"first": "second"}}
	result, _ := MapListConvert(names)
	fmt.Println("String:", result)
	// Output:
	// String: [map[first:second]]
}
