package go_common_lib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCalcHash(t *testing.T) {
	type args struct {
		aStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Hash", args{"testString"}, "4acf0b39d9c4766709a3689f553ac01ab550545ffa4544dfc0b2cea82fba02a3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcHash(tt.args.aStr); got != tt.want {
				t.Errorf("CalcHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceContain(t *testing.T) {
	testSlice := []string{"123"}
	type args struct {
		slice []string
		item  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Contain", args{testSlice, "123"}, true},
		{"Not contain", args{testSlice, "23"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceContain(tt.args.slice, tt.args.item); got != tt.want {
				t.Errorf("SliceContain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateList(t *testing.T) {
	testSlice := []string{"123"}
	testSlice2 := []string{"123", "23"}
	type args struct {
		slice []string
		item  string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Updated list", args{testSlice, "23"}, testSlice2},
		{"Not updated list", args{testSlice, "123"}, testSlice},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateList(tt.args.slice, tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringFromList(t *testing.T) {
	testSlice := []string{"123"}
	testSlice2 := []string{}
	testSlice3 := []string{"123", "23"}
	type args struct {
		list []string
	}
	tests := []struct {
		name   string
		args   args
		wantSt string
	}{
		{"Not empty list", args{testSlice}, "123"},
		{"Empty list", args{testSlice2}, ""},
		{"Not empty list2", args{testSlice3}, "123,23"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSt := StringFromList(tt.args.list); gotSt != tt.wantSt {
				t.Errorf("StringFromList() = %v, want %v", gotSt, tt.wantSt)
			}
		})
	}
}

func TestUpdateStructList(t *testing.T) {
	testSlice := []string{"123", "23"}
	testSlice2 := []string{"321"}
	testSlice3 := []string{"123", "23", "321"}
	testSlice4 := []string{}
	type args struct {
		main []string
		in   []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Not exist", args{testSlice, testSlice2}, testSlice3},
		{"Not Exist", args{testSlice2, testSlice4}, testSlice2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateStructList(tt.args.main, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateStructList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortedMap(t *testing.T) {
	type args struct {
		m []map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantOut []map[string]string
	}{
		{"Un Sort", args{[]map[string]string{{"b": "b"}, {"a": "a"}}}, []map[string]string{{"a": "a"}, {"b": "b"}}},
		{"Sort", args{[]map[string]string{{"a": "a"}, {"b": "b"}}}, []map[string]string{{"a": "a"}, {"b": "b"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := SortedMap(tt.args.m); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("SortedMap() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestListConvert(t *testing.T) {
	var names, names2, names3 interface{}
	names = []string{"first", "second"}
	names2 = []map[string]string{{"first": "second"}}
	names3 = make(chan int)
	type args struct {
		in interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"Not err", args{names}, []string{"first", "second"}, false},
		{"Err", args{names2}, nil, true},
		{"Err2", args{names3}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListConvert(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListConvert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateMapList(t *testing.T) {
	type args struct {
		sliceMaps []map[string]string
		itemMap   map[string]string
	}
	tests := []struct {
		name string
		args args
		want []map[string]string
	}{
		{"Update", args{[]map[string]string{{"a": "b"}}, map[string]string{"b": "b"}}, []map[string]string{{"a": "b"}, {"b": "b"}}},
		{"Update 2", args{[]map[string]string{}, map[string]string{"b": "b"}}, []map[string]string{{"b": "b"}}},
		{"Update 3", args{[]map[string]string{{"a": "b"}, {"b": "b"}}, map[string]string{"a": "c"}}, []map[string]string{{"a": "c"}, {"b": "b"}}},
		{"Not Update", args{[]map[string]string{{"a": "b"}}, map[string]string{"a": "b"}}, []map[string]string{{"a": "b"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateMapList(tt.args.sliceMaps, tt.args.itemMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateMapList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStandardizeSpaces(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Trim first", args{" A A"}, "A A"},
		{"Trim last", args{"A A "}, "A A"},
		{"Trim all", args{" A A "}, "A A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StandardizeSpaces(tt.args.s); got != tt.want {
				t.Errorf("StandardizeSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateListBySplit(t *testing.T) {
	type args struct {
		slice    []string
		item     string
		spliters []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Test1", args{[]string{}, "A,A", nil}, []string{"A"}},
		{"Test2", args{[]string{}, "A,B", nil}, []string{"A", "B"}},
		{"Test3", args{[]string{"A", "B"}, "A,B", nil}, []string{"A", "B"}},
		{"Test4", args{[]string{"A", "B"}, "B\\C", nil}, []string{"A", "B", "C"}},
		{"Test5", args{[]string{"A", "B"}, "B*C", []string{"*"}}, []string{"A", "B", "C"}},
		{"Test5", args{[]string{"A", "B"}, "BC", []string{"*"}}, []string{"A", "B", "BC"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateListBySplit(tt.args.slice, tt.args.item, tt.args.spliters); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateListBySplit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapListConvert(t *testing.T) {
	var names, names2, names3 interface{}
	names = []string{"first", "second"}
	names2 = []map[string]string{{"first": "second"}}
	names3 = make(chan int)
	type args struct {
		in interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []map[string]string
		wantErr bool
	}{
		{"Not err", args{names2}, []map[string]string{{"first": "second"}}, false},
		{"Err1", args{names}, nil, true},
		{"Err2", args{names3}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapListConvert(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapListConvert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapListConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleSliceContain() {
	result := SliceContain([]string{"a", "b", "c"}, "b")
	fmt.Println("Contain:", result)
	// Output:
	// Contain: true
}
func ExampleUpdateList() {
	result := UpdateList([]string{"a", "b", "c"}, "d")
	fmt.Println("New slice:", result)
	// Output:
	// New slice: [a b c d]
}
func ExampleStringFromList() {
	result := StringFromList([]string{"a", "b", "c"})
	fmt.Println("String:", result)
	// Output:
	// String: a,b,c
}
func ExampleUpdateStructList() {
	result := UpdateStructList([]string{"a", "b"}, []string{"a", "c"})
	fmt.Println("Updated slice:", result)
	// Output:
	// Updated slice: [a b c]
}
func ExampleCalcHash() {
	result := CalcHash("testString")
	fmt.Println("Hash:", result)
	// Output:
	// Hash: 4acf0b39d9c4766709a3689f553ac01ab550545ffa4544dfc0b2cea82fba02a3
}
func ExampleSortedMap() {
	result := SortedMap([]map[string]string{{"b": "c", "c": "c", "a": "c"}})
	fmt.Println("Sorted map:", result)
	// Output:
	// Sorted map: [map[a:c] map[b:c] map[c:c]]
}
func ExampleListConvert() {
	var in interface{}
	in = []string{"first", "second"}
	result, _ := ListConvert(in)
	fmt.Println("Slice:", result)
	// Output:
	// Slice: [first second]
}
func ExampleUpdateMapList() {
	result := UpdateMapList([]map[string]string{{"a": "b"}}, map[string]string{"b": "b"})
	fmt.Println("Slice:", result)
	// Output:
	// Slice: [map[a:b] map[b:b]]
}
func ExampleStandardizeSpaces() {
	result := StandardizeSpaces(" ggg ggg ggg ")
	fmt.Println("String:", result)
	// Output:
	// String: ggg ggg ggg
}
func ExampleUpdateListBySplit() {
	result := UpdateListBySplit([]string{"A", "B"}, "B*C", []string{"*"})
	fmt.Println("String:", result)
	// Output:
	// String: [A B C]
}
func ExampleMapListConvert() {
	var names interface{}
	names = []map[string]string{{"first": "second"}}
	result, _ := MapListConvert(names)
	fmt.Println("String:", result)
	// Output:
	// String: [map[first:second]]
}
