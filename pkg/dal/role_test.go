package dal

import (
	"reflect"
	"testing"

	"github.com/rpinedafocus/u-library/pkg/model"
)

func TestCreateRole(t *testing.T) {
	type args struct {
		role *model.Role
	}
	tests := []struct {
		name    string
		args    args
		want    *model.RoleEntity
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateRole("root", tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRole() = %v, want %v", got, tt.want)
			}
		})
	}
}
