package ed25519

import "crypto"

func init() {
	crypto.RegisterHash(crypto.SHA512, _Newi_)
}
