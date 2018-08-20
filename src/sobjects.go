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
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// GetSObject gets an SObject
func (s *Session) GetSObject(name string) (*QueryResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/services/data/v%s/sobjects/%s/", s.OAuth.InstanceURL, s.APIVersion, name), nil)
	if err != nil {
		return nil, err
	}

	s.AddSessionHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := &QueryResponse{}
	err = json.Unmarshal(respBody, response)

	return response, err
}

// CreateSObject creates the objname SObject defined by the payload
func (s *Session) CreateSObject(objname string, payload io.Reader) (*SObjectUpsertResponse, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/services/data/v42.0/sobjects/%s", s.OAuth.InstanceURL, objname), payload)
	if err != nil {
		return nil, err
	}

	s.AddSessionHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Println(string(respBody))

	response := &SObjectUpsertResponse{}
	err = json.Unmarshal(respBody, response)

	return response, err
}

// UpdateSObject PATCHes an SObject in salesforce
func (s *Session) UpdateSObject(objname string, objID string, payload io.Reader) (*SObjectUpsertResponse, error) {
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/services/data/v42.0/sobjects/%s/%s", s.OAuth.InstanceURL, objname, objID), payload)
	if err != nil {
		return nil, err
	}

	s.AddSessionHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Println(string(respBody))

	result := &SObjectUpsertResponse{}
	return result, nil
}
