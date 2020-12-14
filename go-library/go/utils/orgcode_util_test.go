package utils

import (
	"fmt"
	"testing"
)

func TestIsValidOrgCode(t *testing.T) {
	orgCode := "SOB75M0Y-3"
	code := IsValidOrgCode(orgCode)
	fmt.Println(code)
}

func TestIsValidOrgCodeAndLicenseNumber(t *testing.T) {
	type args struct {
		orgCode       string
		licenseNumber string
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
			if err := IsValidOrgCodeAndLicenseNumber(tt.args.orgCode, tt.args.licenseNumber); (err != nil) != tt.wantErr {
				t.Errorf("IsValidOrgCodeAndLicenseNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
