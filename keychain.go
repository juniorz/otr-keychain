package keychain

import (
	"bytes"

	"golang.org/x/crypto/otr"
)

// ImportFromLibOTR parses the contents of a libotr private key file and imports all the keys found.
func ImportFromLibOTR(in []byte) []otr.PrivateKey {
	acctStart := []byte("(account")
	ret := []otr.PrivateKey{}

	var i, p int
	for {
		i = bytes.Index(in[p:], acctStart)
		if i == -1 {
			break
		}

		p += i + len(acctStart)

		key := otr.PrivateKey{}
		if key.Import(in[p:]) {
			ret = append(ret, key)
		}
	}

	return ret
}