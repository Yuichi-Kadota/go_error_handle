package main

import (
	"errors"
	"reflect"
	"testing"

	httpErr "github.com/Yuichi-Kadota/go_error_handle/internal/error"
	"github.com/k0kubun/pp"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		isErr    bool
		wantErr  error
		wantCode int
	}{
		{
			name:    "err_internal",
			num:     0,
			isErr:   true,
			wantErr: httpErr.InternalServerErr{},
		},
		{
			name:    "err_unAuthorized",
			num:     1,
			isErr:   true,
			wantErr: httpErr.UnAuthorized{},
		},
		{
			name:    "err_forbidden",
			num:     2,
			isErr:   true,
			wantErr: httpErr.Forbidden{},
		},
		{
			name:    "err_uncontrolled",
			num:     3,
			isErr:   true,
			wantErr: errors.New(""),
		},
		{
			name:    "other case",
			num:     4,
			isErr:   false,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Process(tt.num)
			if (err != nil) == tt.isErr {
				t.Errorf("Process() unit test case error")
			}
			assertError(t, err, tt.wantErr)
		})
	}
}

// ユーザー定義のエラー型が一致するか検証する
func assertError(t *testing.T, err error, wantErr error) {
	target := reflect.TypeOf(err)
	want := reflect.TypeOf(wantErr)
	if target != want {
		pp.Printf("err : %s\n", target.String())
		pp.Printf("wantErr : %s\n", want.String())
		t.Error("err != wantErr")
	}
}
