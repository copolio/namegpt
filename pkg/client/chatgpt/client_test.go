package chatgpt

import (
	"testing"
)

func TestGetSimilarDomains(t *testing.T) {
	type args struct {
		keyword string
	}
	tests := []struct {
		name        string
		args        args
		wantDomains int
		wantErr     bool
	}{
		{
			name: "Successful request to openai api and return list of at least 50 domains",
			args: args{
				keyword: "test.com",
			},
			wantDomains: 50,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDomains, err := GetSimilarDomains(tt.args.keyword)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSimilarDomains() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("GetSimilarDomains() gotDomains = %v", gotDomains)
			gotDomainsLen := len(gotDomains)
			if gotDomainsLen < tt.wantDomains {
				t.Errorf("GetSimilarDomains() gotDomains = %v, want %v", gotDomainsLen, tt.wantDomains)
			}
		})
	}
}
