package model

import (
	"crypto/sha512"
	"encoding/hex"
)

func (ud *userDomain) EncryptPassword() {
	hash := sha512.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))	
}