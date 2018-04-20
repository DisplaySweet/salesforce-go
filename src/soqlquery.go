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
	"log"
	"net/http"
)

// SOQLQuery executes an SOQL query and returns the response
func (s *Session) SOQLQuery(query string) (*QueryResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/services/data/v%s/query/", s.OAuth.InstanceURL, s.APIVersion), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("q", query)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", fmt.Sprintf("%s %s", s.OAuth.TokenType, s.OAuth.AccessToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var temp *QueryResponse
	err = json.Unmarshal(respBody, temp)
	if err != nil {
		log.Println(string(respBody))
		return nil, err
	}

	return temp, nil
}
