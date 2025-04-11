package sieve

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNthPrime(t *testing.T) {
	t.Parallel()
	type args struct {
		n int64
	}
	type want struct {
		result int64
		err    error
	}
	tests := []struct {
		args *args
		want *want
		name string
	}{
		{
			name: "n less than 0",
			args: &args{
				n: -1,
			},
			want: &want{
				result: 0,
				err:    errors.New("invalid input"),
			},
		},
		{
			name: "n 0",
			args: &args{
				n: 0,
			},
			want: &want{
				result: 2,
				err:    nil,
			},
		},
		{
			name: "n 1",
			args: &args{
				n: 1,
			},
			want: &want{
				result: 3,
				err:    nil,
			},
		},
		{
			name: "n 6",
			args: &args{
				n: 6,
			},
			want: &want{
				result: 17,
				err:    nil,
			},
		},
		{
			name: "n 7",
			args: &args{
				n: 7,
			},
			want: &want{
				result: 19,
				err:    nil,
			},
		},
		{
			name: "n 19",
			args: &args{
				n: 19,
			},
			want: &want{
				result: 71,
				err:    nil,
			},
		},
		{
			name: "n 99",
			args: &args{
				n: 99,
			},
			want: &want{
				result: 541,
				err:    nil,
			},
		},
		{
			name: "n 500",
			args: &args{
				n: 500,
			},
			want: &want{
				result: 3581,
				err:    nil,
			},
		},
		{
			name: "n 986",
			args: &args{
				n: 986,
			},
			want: &want{
				result: 7793,
				err:    nil,
			},
		},
		{
			name: "n 2000",
			args: &args{
				n: 2000,
			},
			want: &want{
				result: 17393,
				err:    nil,
			},
		},
		{
			name: "n 1000000",
			args: &args{
				n: 1000000,
			},
			want: &want{
				result: 15485867,
				err:    nil,
			},
		},
		{
			name: "n 10000000",
			args: &args{
				n: 10000000,
			},
			want: &want{
				result: 179424691,
				err:    nil,
			},
		},
		{
			name: "n 100000000",
			args: &args{
				n: 100000000,
			},
			want: &want{
				result: 2038074751,
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()
			sieve := NewSieve()
			got, err := sieve.NthPrime(tt.args.n)
			require.Equal(t, tt.want.result, got)
			require.Equal(t, tt.want.err, err)
		})
	}
}
