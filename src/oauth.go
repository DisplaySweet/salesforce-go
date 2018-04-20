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

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	loginURI = "https://login.salesforce.com/services/oauth2/token"
)

// OAuthDetails stores the response from an OAuth login
type OAuthDetails struct {
	ID          string
	AccessToken string `json:"access_token"`
	InstanceURL string `json:"instance_url"`
	TokenType   string `json:"token_type"`
	IssuedAt    string `json:"issued_at"`
	Signature   string
}

// LoginByOAuth attempts to log in by oauth
func LoginByOAuth(creds Credentials) (*Session, error) {
	payload := url.Values{
		"grant_type":    {creds.GrantType},
		"client_id":     {creds.ClientID},
		"client_secret": {creds.ClientSecret},
		"username":      {creds.Username},
		"password":      {fmt.Sprintf("%v%v", creds.Password, creds.Token)},
		"code":          {creds.Code},
		"redirect_uri":  {creds.RedirectURI},
	}

	body := strings.NewReader(payload.Encode())

	req, err := http.NewRequest("POST", loginURI, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var oauthResp OAuthDetails
	err = json.Unmarshal(respBody, &oauthResp)
	if err != nil {
		return nil, err
	}

	session := &Session{
		OAuth:      oauthResp,
		APIVersion: creds.APIVersion,
	}

	return session, nil
}
