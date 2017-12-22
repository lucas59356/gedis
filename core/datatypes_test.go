package core

import "testing"

func TestGuessDataType(t *testing.T) {
	type args struct {
		v interface{}
	}
	var i8 int8 = 2
	var i16 int16 = 2
	var i32 int32 = 2
	var i64 int64 = 2
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
		{
			name: "Int8",
			args: args{
				v: i8,
			},
			want: TypeInt8,
		},
		{
			name: "Int16",
			args: args{
				v: i16,
			},
			want: TypeInt16,
		},
		{
			name: "Int32",
			args: args{
				v: i32,
			},
			want: TypeInt32,
		},
		{
			name: "Int64",
			args: args{
				v: i64,
			},
			want: TypeInt64,
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
