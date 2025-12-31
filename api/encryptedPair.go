/*
The MIT License (MIT)

Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package api

import (
	"path"

	"github.com/J-Siu/go-crypto/crypto"
	"github.com/J-Siu/go-gitapi/v3/base"
	"github.com/J-Siu/go-gitapi/v3/info"
)

// Github repository action secret structure
// Do() handles public key
type EncryptedPair struct {
	*base.Base
	Info  info.EncryptedPair
	name  string
	value string
}

func (t *EncryptedPair) New(property *base.Property) *EncryptedPair {
	property.Info = &t.Info
	t.Base = new(base.Base).New(property).EndpointReposSecrets()
	return t
}

func (t *EncryptedPair) Set(name, value string) *EncryptedPair {
	t.name = name
	t.value = value
	t.Base.Req.Endpoint = path.Join(t.Base.Req.Endpoint, t.name)
	t.SetPut()
	return t
}

// Do() handles public key
func (t *EncryptedPair) Do() *base.Base {
	// Get public key -- start
	var (
		publicKey = new(PublicKey).New(t.Property)
	)
	if !publicKey.Do().Ok() {
		return publicKey.Base
	}
	// Get public key -- end
	t.encrypt(&publicKey.Info)
	if *t.Err() == "" {
		t.Base.Do()
	}
	return t.Base
}

func (t *EncryptedPair) encrypt(pk *info.PublicKey) *EncryptedPair {
	t.Info.Key_id = pk.Key_id
	encrypted_value, e := crypto.BoxSealAnonymous(&pk.Key, &t.value)
	t.Info.Encrypted_value = *encrypted_value
	if e != nil {
		t.Res.Err = e.Error()
	}
	return t
}
