package function

import (
	"net/http"
	"strconv"

	"github.com/purnaresa/trustme/pkg/keygen"
)

func GenerateKeyPair(w http.ResponseWriter, r *http.Request) {
	keypair, err := keygen.GenerateKeyPair()
	if err != nil {
		http.Error(w,
			"Key pair failed to generated",
			http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=key-pair.zip")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", strconv.FormatInt(int64(len(keypair)), 10))
	w.Write(keypair)
}
