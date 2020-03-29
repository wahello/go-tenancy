package services

import (
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/snowlyg/go-tenancy/models"
)

func TestNewPermService(t *testing.T) {
	type args struct {
		gdb *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want PermService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPermService(tt.args.gdb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPermService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_permService_Create(t *testing.T) {
	type fields struct {
		gdb *gorm.DB
	}
	type args struct {
		menu *models.Perm
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &permService{
				gdb: tt.fields.gdb,
			}
			if err := s.Create(tt.args.menu); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_permService_GetAll(t *testing.T) {
	type fields struct {
		gdb *gorm.DB
	}
	type args struct {
		args        map[string]interface{}
		typefilters []string
		ispreload   bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
		want1  []*models.Perm
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &permService{
				gdb: tt.fields.gdb,
			}
			got, got1 := s.GetAll(tt.args.args, tt.args.typefilters, tt.args.ispreload)
			if got != tt.want {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetAll() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_permService_GetPermissionByHrefMethod(t *testing.T) {
	type fields struct {
		gdb *gorm.DB
	}
	type args struct {
		href   string
		method string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   models.Perm
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &permService{
				gdb: tt.fields.gdb,
			}
			got, got1 := s.GetPermissionByHrefMethod(tt.args.href, tt.args.method)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPermissionByHrefMethod() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetPermissionByHrefMethod() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}