package chatgpt

import (
	"errors"
	"fmt"
	"io"
	"strings"
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
			name: "Successful request to openai api and return list of at least 10 domains",
			args: args{
				keyword: "test.com",
			},
			wantDomains: 10,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDomains, err := GenerateDomainNames(tt.args.keyword)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateDomainNames() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("GenerateDomainNames() gotDomains = %v", gotDomains)
			gotDomainsLen := len(gotDomains)
			if gotDomainsLen < tt.wantDomains {
				t.Errorf("GenerateDomainNames() gotDomains = %v, want %v", gotDomainsLen, tt.wantDomains)
			}
		})
	}
}

func TestGetSimilarDomainStream(t *testing.T) {
	type args struct {
		keyword string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test stream",
			args:    args{"naver.com"},
			want:    50,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRecommendedDomainStream(tt.args.keyword)
			defer got.Close()
			fmt.Println("Stream start")
			var domains = make([]string, 50)
			var domain = ""
			for {
				recv, err := got.Recv()
				if errors.Is(err, io.EOF) {
					fmt.Println("\nStream finished")
					break
				}
				if err != nil {
					fmt.Errorf("Stream Recv failed: %w", err)
					break
				}
				fmt.Printf("Recv: %s\n", recv.Choices[0].Delta.Content)
				if strings.Contains(recv.Choices[0].Delta.Content, "\"") {
					if domain != "" {
						domains = append(domains, domain)
						fmt.Printf("Domain: %s\n", domain)
					}
					domain = ""
				} else {
					domain += recv.Choices[0].Delta.Content
				}
			}
			fmt.Println("Stream end")
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRecommendedDomainStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(domains) < tt.want {
				t.Errorf("GetRecommendedDomainStream() got = %v, want %v", got, tt.want)
			}
		})
	}
}
