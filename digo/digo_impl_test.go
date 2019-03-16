package digo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDiGoImpl_SingletonFunc(t *testing.T) {
	type fields struct {
		DiGoStub      stub
		binding       map[Type]interface{}
		store         map[Type]reflect.Value
		share         map[Type]bool
		startBuilding map[Type]struct{}
		finishBuilt   map[Type]struct{}
	}
	type args struct {
		fn interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{func() int { return 316 }},
		},
		{
			args: args{func(i int) string { return fmt.Sprint(i) }},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			di := NewDiGoImpl()
			if err := di.SingletonFunc(tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("DiGoImpl.SingletonFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
