/*
MIT License

Copyright (c) 2018 Display Sweet

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package salesforce

// OAuthDetails stores the response from an OAuth login
type OAuthDetails struct {
	ID          string
	AccessToken string `json:"access_token"`
	InstanceURL string `json:"instance_url"`
	TokenType   string `json:"token_type"`
	IssuedAt    string `json:"issued_at"`
	Signature   string
}

// QueryResponse holds a response from Salesfroce
type QueryResponse struct {
	TotalSize int                      `json:"totalSize"`
	Done      bool                     `json:"done"`
	Records   []map[string]interface{} `json:"records"`
}

// SObjectUpsertResponse is returned when a POST or PATCH is sent to Salesforce
type SObjectUpsertResponse struct {
	ID      string        `json:"id"`
	Success bool          `json:"success"`
	Errors  []interface{} `json:"errors"`
}
