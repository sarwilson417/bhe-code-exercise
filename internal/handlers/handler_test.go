package handlers_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/sarwilson417/bhe-code-exercise/go/internal/handlers"
	"github.com/sarwilson417/bhe-code-exercise/go/pkg/sieve"
)

func TestGetNthPrime(t *testing.T) {
	t.Parallel()
	type args struct {
		n      string
		method string
	}
	type want struct {
		statusCode int
		body       string
	}
	type mocks struct {
		mockSieve *sieve.MockSieve
	}
	tests := []struct {
		args    *args
		want    *want
		prepare func(m *mocks)
		name    string
	}{
		{
			name: "wrong method",
			args: &args{
				n:      "abc",
				method: http.MethodPost,
			},
			prepare: func(m *mocks) {},
			want: &want{
				statusCode: http.StatusMethodNotAllowed,
				body:       `method not supported`,
			},
		},
		{
			name: "n not provided",
			args: &args{
				n:      "",
				method: http.MethodGet,
			},
			prepare: func(m *mocks) {},
			want: &want{
				statusCode: http.StatusBadRequest,
				body:       `missing required query param 'n'`,
			},
		},
		{
			name: "n is not a number",
			args: &args{
				n:      "nan",
				method: http.MethodGet,
			},
			prepare: func(m *mocks) {},
			want: &want{
				statusCode: http.StatusBadRequest,
				body:       `invalid value for query param 'n', must be in the range [0-100000000]`,
			},
		},
		{
			name: "n must not be negative",
			args: &args{
				n:      "-1",
				method: http.MethodGet,
			},
			prepare: func(m *mocks) {},
			want: &want{
				statusCode: http.StatusBadRequest,
				body:       `invalid value for query param 'n', must be in the range [0-100000000]`,
			},
		},
		{
			name: "n is too big for int64",
			args: &args{
				n:      "9223372036854775808",
				method: http.MethodGet,
			},
			prepare: func(m *mocks) {},
			want: &want{
				statusCode: http.StatusBadRequest,
				body:       `invalid value for query param 'n', must be in the range [0-100000000]`,
			},
		},
		{
			name: "error: SieveService.NthPrime",
			args: &args{
				n:      "0",
				method: http.MethodGet,
			},
			prepare: func(m *mocks) {
				m.mockSieve.EXPECT().NthPrime(int64(0)).Return(int64(0), errors.New("error"))
			},
			want: &want{
				statusCode: http.StatusInternalServerError,
				body:       `error finding nth prime`,
			},
		},
		{
			name: "success - min range",
			args: &args{
				n:      "0",
				method: http.MethodGet,
			},
			prepare: func(m *mocks) {
				m.mockSieve.EXPECT().NthPrime(int64(0)).Return(int64(2), nil)
			},
			want: &want{
				statusCode: http.StatusOK,
				body:       `{"nthPrime":2}`,
			},
		},
		{
			name: "success - mid range",
			args: &args{
				n:      "986",
				method: http.MethodGet,
			},
			prepare: func(m *mocks) {
				m.mockSieve.EXPECT().NthPrime(int64(986)).Return(int64(7793), nil)
			},
			want: &want{
				statusCode: http.StatusOK,
				body:       `{"nthPrime":7793}`,
			},
		},
		{
			name: "success - max range",
			args: &args{
				n:      "100000000",
				method: http.MethodGet,
			},
			prepare: func(m *mocks) {
				m.mockSieve.EXPECT().NthPrime(int64(100000000)).Return(int64(2038074751), nil)
			},
			want: &want{
				statusCode: http.StatusOK,
				body:       `{"nthPrime":2038074751}`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			m := mocks{
				mockSieve: sieve.NewMockSieve(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&m)
			}

			h := handlers.New(m.mockSieve)

			req := httptest.NewRequest(tt.args.method, "/primes?n="+tt.args.n, nil)
			w := httptest.NewRecorder()

			handler := http.HandlerFunc(h.GetNthPrime)
			handler.ServeHTTP(w, req)

			res := w.Result()
			defer res.Body.Close()

			body, _ := io.ReadAll(res.Body)

			if res.StatusCode != tt.want.statusCode {
				t.Errorf("did not receive expected status; want %d, got %d", tt.want.statusCode, res.StatusCode)
			}

			if strings.TrimSpace(string(body)) != tt.want.body {
				t.Errorf("did not receive expected body; want %q, got %q", tt.want.body, body)
			}
		})
	}
}
