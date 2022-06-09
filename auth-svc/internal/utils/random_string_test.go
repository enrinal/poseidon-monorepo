package utils

import "testing"

func TestGenerateRandomString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "generate random string with length 4",
			args: args{
				length: 4,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRandomString(tt.args.length); len(got) != tt.want {
				t.Errorf("GenerateRandomString() = %v, want %v", got, tt.want)
			}
		})
	}
}
