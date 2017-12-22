package core

import "testing"

func TestGuessDataType(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			name: "Bool",
			args: args{
				v: true,
			},
			want: TypeBool,
		},
		{
			name: "String",
			args: args{
				v: "Teste",
			},
			want: TypeString,
		},
		{
			name: "Int",
			args: args{
				v: 2,
			},
			want: TypeInt,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GuessDataType(tt.args.v); got != tt.want {
				t.Errorf("GuessDataType() = %v, want %v", got, tt.want)
			}
		})
	}
}
