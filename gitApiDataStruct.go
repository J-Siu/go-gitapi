/*
Copyright © 2022 John, Sing Dao, Siu <john.sd.siu@gmail.com>

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

package gitapi

import "github.com/J-Siu/go-helper"

// Github repository action secret structure
type RepoEncryptedPair struct {
	Encrypted_value string `json:"encrypted_value"` // Encrypted value
	Key_id          string `json:"key_id"`          // Public key id
}

func (self *RepoEncryptedPair) StringP() *string {
	var str string
	str += "Value:" + self.Encrypted_value + "\n"
	str += "Key ID:" + self.Key_id + "\n"
	return &str
}

func (self *RepoEncryptedPair) String() string {
	return *self.StringP()
}

// Github repository public key structure
type RepoPublicKey struct {
	Key_id string `json:"key_id"`
	Key    string `json:"key"`
}

func (self *RepoPublicKey) StringP() *string {
	var str string
	str += "Key:" + self.Key + "\n"
	str += "Key ID:" + self.Key_id + "\n"
	return &str
}

func (self *RepoPublicKey) String() string {
	return *self.StringP()
}

// Github repository private structure
type RepoPrivate struct {
	Private bool `json:"private"`
}

func (self *RepoPrivate) StringP() *string {
	var str string
	str += helper.BoolString(self.Private)
	return &str
}

func (self *RepoPrivate) String() string {
	return *self.StringP()
}

// Github repository visibility structure
type RepoVisibility struct {
	Visibility string `json:"visibility"`
}

func (self *RepoVisibility) StringP() *string {
	return &self.Visibility
}

func (self *RepoVisibility) String() string {
	return *self.StringP()
}

// Github repository description structure
type RepoDescription struct {
	Description string `json:"description"`
}

func (self *RepoDescription) StringP() *string {
	return &self.Description
}

func (self *RepoDescription) String() string {
	return *self.StringP()
}

// Github repository topics structure
type RepoTopics struct {
	Topics *[]string `json:"topics"` // Github topics is "Topics"
	Names  *[]string `json:"names"`  // Gitea topics is "Names"
}

func (self *RepoTopics) StringP() *string {
	var str string
	// Gitea
	if self.Names != nil {
		for _, t := range *self.Names {
			str += t + "\n"
		}
	}
	// Github
	if self.Topics != nil {
		for _, t := range *self.Topics {
			str += t + "\n"
		}
	}
	return &str
}

func (self *RepoTopics) String() string {
	return *self.StringP()
}

// Github repository(creation) info structure
type RepoInfo struct {
	Name    string `json:"name"`
	Private bool   `json:"private"`
}

func (self *RepoInfo) StringP() *string {
	var str string
	str = self.Name + " (private:" + helper.BoolString(self.Private) + ")"
	return &str
}

func (self *RepoInfo) String() string {
	return *self.StringP()
}

// Github repository(creation) info array
type RepoInfoList []RepoInfo

func (self *RepoInfoList) StringP() *string {
	var str string
	for _, i := range *self {
		str += *i.StringP() + "\n"
	}
	return &str
}

func (self *RepoInfoList) String() string {
	return *self.StringP()
}

// This is used as a dummy type for nil GitApiInfo parameter
type NilType struct {
	Nil *byte // Dummy pointer
}

func (self *NilType) StringP() *string {
	return nil
}

func (self *NilType) String() string {
	return "Null"
}

// Return a *NilType(nil) dummy pointer
func Nil() *NilType {
	return nil
}

// GitApi structures interface
type GitApiInfo interface {
	// *RepoEncryptedPair |
	// *RepoPublicKey |
	// *RepoPrivate |
	// *RepoVisibility |
	// *RepoDescription |
	// *RepoTopics |
	// *RepoInfo |
	// *RepoInfoList |
	// *NilType
	StringP() *string
	String() string
}
