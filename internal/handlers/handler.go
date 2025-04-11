package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/sarwilson417/bhe-code-exercise/go/pkg/sieve"
)

const (
	maxN = 100000000
	minN = 0
)

type Handlers struct {
	SieveService sieve.Sieve
}

func New(sieveService sieve.Sieve) Handlers {
	return Handlers{
		SieveService: sieveService,
	}
}

type PrimesResponse struct {
	Prime int64 `json:"nthPrime"`
}

func (h Handlers) GetNthPrime(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	n := params.Get("n")

	if r.Method != http.MethodGet {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)

		return
	}

	if n == "" {
		http.Error(w, "missing required query param 'n'", http.StatusBadRequest)

		return
	}

	num, err := convertToPositiveInt64(n)
	if err != nil {
		log.Printf("error converting query param 'n' to int64: %v", err)
		http.Error(w, "invalid value for query param 'n', must be in the range [0-100000000]", http.StatusBadRequest)

		return
	}

	nthPrime, err := h.SieveService.NthPrime(num)
	if err != nil {
		log.Printf("error finding nth prime: %v", err)
		http.Error(w, "error finding nth prime", http.StatusInternalServerError)
		
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(PrimesResponse{Prime: nthPrime})
}

func convertToPositiveInt64(n string) (int64, error) {
	num, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		return 0, err
	}
	if num < minN || num > maxN {
		return num, errors.New("value is outside allowed range")
	}

	return num, nil
}
