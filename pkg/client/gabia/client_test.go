package gabia

import (
	"reflect"
	"testing"
)

func TestCheckDomainRegist(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name    string
		args    args
		want    *RegistCheckResult
		wantErr bool
	}{
		{
			name:    "Return Normal Response",
			args:    struct{ domain string }{domain: "test123123123.com"},
			want:    &RegistCheckResult{Domain: "test123123123.com"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckDomainRegist(tt.args.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckDomainRegist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Domain, tt.want.Domain) {
				t.Errorf("CheckDomainRegist() got = %v, want %v", got, tt.want)
			}
		})
	}
}
