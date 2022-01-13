package readdir

import (
	"reflect"
	"testing"
)

func TestReadDir(t *testing.T) {
	type args struct {
		root string
	}
	tests := []struct {
		name string
		args args
		want <-chan string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadDir(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
