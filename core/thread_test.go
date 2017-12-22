package core

import (
	"reflect"
	"testing"
)

func TestThread_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		t       Thread
		args    args
		want    interface{}
		want1   int8
		wantErr bool
	}{
		{
			name: "Tá getável?",
			t: Thread{
				Types: map[string]int8{
					"teste": TypeString,
				},
				Values: map[string]interface{}{
					"teste": "Teste",
				},
			},
			args: args{
				key: "Teste",
			},
			want:    "Teste",
			want1:   TypeString,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.t.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Thread.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Thread.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Thread.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestThread_Set(t *testing.T) {
	type args struct {
		key string
		v   interface{}
	}
	tr := NewThread()
	tests := []struct {
		name    string
		t       *Thread
		args    args
		want    interface{}
		want1   int8
		wantErr bool
	}{
		{
			name: "Basic set",
			t:    tr,
			args: args{
				key: "Teste",
				v:   "EOQ",
			},
			want:    "EOQ",
			want1:   TypeString,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.t.Set(tt.args.key, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Thread.Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Thread.Set() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Thread.Set() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRW(t *testing.T) {
	frase := "Isto é um teste"
	chave := "teste"
	th := NewThread()
	v, tp, err := th.Set(chave, frase)
	if v != frase {
		t.Fail()
	}
	if tp != TypeString {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}

	v, tp, err = th.Get(chave)
	if v != frase {
		t.Fail()
	}
	if tp != TypeString {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
}
