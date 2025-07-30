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

package gitapi

import "github.com/J-Siu/go-helper"

// Github repository action secret structure
type RepoEncryptedPair struct {
	Encrypted_value string `json:"encrypted_value"` // Encrypted value
	Key_id          string `json:"key_id"`          // Public key id
}

func (rEncryptedPair *RepoEncryptedPair) StringP() *string {
	var str string
	str += "Value:" + rEncryptedPair.Encrypted_value + "\n"
	str += "Key ID:" + rEncryptedPair.Key_id + "\n"
	return &str
}

func (rEncryptedPair *RepoEncryptedPair) String() string {
	return *rEncryptedPair.StringP()
}

// Github repository public key structure
type RepoPublicKey struct {
	Key_id string `json:"key_id"`
	Key    string `json:"key"`
}

func (rPKey *RepoPublicKey) StringP() *string {
	var str string
	str += "Key:" + rPKey.Key + "\n"
	str += "Key ID:" + rPKey.Key_id + "\n"
	return &str
}

func (rPKey *RepoPublicKey) String() string {
	return *rPKey.StringP()
}

// Github repository private structure
type RepoPrivate struct {
	Private bool `json:"private"`
}

func (rPrivate *RepoPrivate) StringP() *string {
	var str string
	str += helper.BoolString(rPrivate.Private)
	return &str
}

func (rPrivate *RepoPrivate) String() string {
	return *rPrivate.StringP()
}

// Github repository visibility structure
type RepoVisibility struct {
	Visibility string `json:"visibility"`
}

func (rVisibility *RepoVisibility) StringP() *string {
	return &rVisibility.Visibility
}

func (rVisibility *RepoVisibility) String() string {
	return *rVisibility.StringP()
}

// Github repository description structure
type RepoDescription struct {
	Description string `json:"description"`
}

func (rDesc *RepoDescription) StringP() *string {
	return &rDesc.Description
}

func (rDesc *RepoDescription) String() string {
	return *rDesc.StringP()
}

// Github repository topics structure
type RepoTopics struct {
	Topics *[]string `json:"topics"` // Github topics is "Topics"
	Names  *[]string `json:"names"`  // Gitea topics is "Names"
}

func (rTopics *RepoTopics) StringP() *string {
	var str string
	// Gitea
	if rTopics.Names != nil {
		for _, t := range *rTopics.Names {
			str += t + "\n"
		}
	}
	// Github
	if rTopics.Topics != nil {
		for _, t := range *rTopics.Topics {
			str += t + "\n"
		}
	}
	return &str
}

func (rTopics *RepoTopics) String() string {
	return *rTopics.StringP()
}

// Github repository(creation) info structure
type RepoInfo struct {
	Name    string `json:"name"`
	Private bool   `json:"private"`
}

func (rInfo *RepoInfo) StringP() *string {
	str := rInfo.Name + " (private:" + helper.BoolString(rInfo.Private) + ")"
	return &str
}

func (rInfo *RepoInfo) String() string {
	return *rInfo.StringP()
}

/*
Github uses message and status.
Gitea uses message and errors.
*/
type RepoError struct {
	Errors  string `json:"errors"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func (rError *RepoError) StringP() *string {
	// return helper.ReportSp(rError, "", true, false)
	return &rError.Message
}

func (rError *RepoError) String() string {
	return *rError.StringP()
}

// Github repository(creation) info array
type RepoInfoList []RepoInfo

func (rInfoList *RepoInfoList) StringP() *string {
	var str string
	for _, i := range *rInfoList {
		str += *i.StringP() + "\n"
	}
	return &str
}

func (rInfoList *RepoInfoList) String() string {
	return *rInfoList.StringP()
}

// This is used as a dummy type for nil GitApiInfo parameter
type NilType struct {
	Nil *byte // Dummy pointer
}

func (nilT *NilType) StringP() *string {
	return nil
}

func (nilT *NilType) String() string {
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
