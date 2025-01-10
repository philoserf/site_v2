package main

import (
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_processFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := processFile(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("processFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_stripFrontMatter(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stripFrontMatter(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stripFrontMatter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stripH1(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stripH1(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stripH1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stripFooter(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stripFooter(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stripFooter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_renameFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := renameFile(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("renameFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
