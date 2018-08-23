package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"hash"
)

// HMAC is a wrapper around the crypto/hmac package
// making it a little easier to use in th code
type HMAC struct {
	hmac hash.Hash
}

// NewHMAC returns a new HMAC object
func NewHMAC(key string) HMAC {
	h := hmac.New(sha256.New, []byte(key))
	return HMAC{
		hmac: h,
	}
}

// Hash will hash the provides input string using HMAC with
// the secret key provided when HMAC object is created
func (h HMAC) Hash(input string) string {
	h.hmac.Reset()
	h.hmac.Write([]byte(input))
	return base64.StdEncoding.EncodeToString(h.hmac.Sum(nil))
}
