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

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/DisplaySweet/salesforce-go/src"
)

func main() {

	creds := salesforce.Credentials{
		Username:     "test",
		Password:     "secret_password",
		Token:        "10xnklsdf023jzxcsdfgjk",
		Code:         "code",
		GrantType:    "token",
		ClientID:     "3MVG9lKcPoNINVBJSoQsNCD.HHDdbugPsNXwwyFbgb47KWa_PTv",
		ClientSecret: "5678471853609579508",
		RedirectURI:  "https://localhost:8443/RestTest/oauth/_callback",
		APIVersion:   "42.0",
	}

	session, err := salesforce.LoginByOAuth(creds)
	if err != nil {
		log.Fatalln(err)
	}

	contact := Contact{
		FirstName: "What",
		LastName:  "Am",
		Phone:     "I",
		Email:     "Doing?",
	}

	b, err := json.Marshal(contact)
	if err != nil {
		log.Fatalln(err)
	}

	payload := bytes.NewReader(b)
	if err != nil {
		log.Fatalln(err)
	}

	result, err := session.CreateSObject("Contact", payload)
	if err != nil {
		log.Fatalln(err)
	}

	contactID := result.ID

	updateContact := Contact{
		FirstName: "Update",
		LastName:  "The",
		Phone:     "Previous",
		Email:     "Contact!",
	}

	b, err = json.Marshal(updateContact)
	if err != nil {
		log.Fatalln(err)
	}

	payload = bytes.NewReader(b)
	if err != nil {
		log.Fatalln(err)
	}

	contactEndpoint := fmt.Sprintf("Contact/%s", contactID)

	anotherResult, err := session.UpdateSObject(contactEndpoint, payload)

	response, err := session.GetSObject(contactEndpoint)
	if err != nil {
		log.Fatalln(err)
	}

	if response.TotalSize > 0 {
		contactData := response.Records[0]
		contactID := contactData["Id"].(string)
	}
}
