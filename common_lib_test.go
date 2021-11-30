package go_common_lib

import (
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
		{"True", args{testSlice, "123"}, true},
		{"False", args{testSlice, "23"}, true},
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
		{"Not exist", args{testSlice, "23"}, testSlice2},
		{"Not Exist", args{testSlice, "123"}, testSlice},
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
	testSlice := []string{"123", "23"}
	testSlice2 := []string{}
	type args struct {
		list []string
	}
	tests := []struct {
		name   string
		args   args
		wantSt string
	}{
		{"Not exist", args{testSlice}, "123, 23"},
		{"Not Exist", args{testSlice2}, ""},
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
