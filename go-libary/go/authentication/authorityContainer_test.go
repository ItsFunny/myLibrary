package authentication

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewAuthority(t *testing.T) {
	authority := NewAuthority()
	// bytes, e := json.Marshal(authority)
	// if nil != e {
	// 	panic(e)
	// }
	fmt.Println(authority)
	fmt.Println(len(authority))
	// fmt.Println(len(bytes))
	// fmt.Println(bytes)
	fmt.Println("====")

	authority.AddAuthentication(AuthValue(1))
	fmt.Println(authority)
	fmt.Println(len(authority))
	fmt.Println("====")
	authority.AddAuthentication(AuthValue(2))
	fmt.Println(authority)
	fmt.Println(len(authority))

	// tests := []struct {
	// 	name string
	// 	want Authority
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := NewAuthority(); !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("NewAuthority() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}

func TestAuthority_AddAuthentication(t *testing.T) {
	type args struct {
		value AuthValue
	}
	tests := []struct {
		name     string
		receiver Authority
		args     args
		want     Authority
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.receiver.AddAuthentication(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authority.AddAuthentication() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthority_CheckAuthentication(t *testing.T) {
	type args struct {
		authValues []AuthValue
	}
	tests := []struct {
		name     string
		receiver Authority
		args     args
		want     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.receiver.CheckAuthentication(tt.args.authValues...); got != tt.want {
				t.Errorf("Authority.CheckAuthentication() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthority_CheckAuthority(t *testing.T) {
	type args struct {
		authority Authority
	}
	tests := []struct {
		name     string
		receiver Authority
		args     args
		want     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.receiver.CheckAuthority(tt.args.authority); got != tt.want {
				t.Errorf("Authority.CheckAuthority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthority_DeleteAuthentication(t *testing.T) {
	type args struct {
		authValue AuthValue
	}
	tests := []struct {
		name     string
		receiver Authority
		args     args
		want     Authority
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.receiver.DeleteAuthentication(tt.args.authValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authority.DeleteAuthentication() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthority_CreateSuperSuperAdmin(t *testing.T) {
	tests := []struct {
		name     string
		receiver Authority
		want     Authority
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.receiver.CreateSuperSuperAdmin(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authority.CreateSuperSuperAdmin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBigEndianConvtBytes2Authority(t *testing.T) {
	authority := NewAuthority()
	authority = authority.AddAuthentication(AuthValue(64))
	bytes := authority.BigEndianConvt2Bytes()
	fmt.Println(authority)
	fmt.Println(len(bytes))
	fmt.Println(bytes)
	fmt.Println("===")
	bytes2Authority, e := BigEndianConvtBytes2Authority(bytes)
	if nil != e {
		panic(e)
	}
	fmt.Println(bytes2Authority)
}
