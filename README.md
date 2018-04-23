# Salesforce-Go

A library to help consume Salesforce objects, useful when creating external integrations to push to or pull data from Salesforce.

To download the package:
```
> go get github.com/DisplaySweet/salesforce-go/src
```


Usage:
```
func main() {

    //Enter your salesforce credentials
	creds := salesforce.Credentials{
		Username:     "test",
		Password:     "secret_password",
		Token:        "10xnklsdf023jzxcsdfgjk",
		Code:         "code",
		GrantType:    "token",
		ClientID:     "3MVG9lKcPoNINVBJSoQsNCD.HHDdbugPsNXwwyFbgb47KWa_PTv",
		ClientSecret: "5678471853609579508",
		RedirectURI:  "https://localhost:8443/RestTest/oauth/_callback",
		APIVersion:   "2",
	}

    //Authorise with salesforce
	session, err := salesforce.LoginByOAuth(creds)
	if err != nil {
		log.Fatalln(err)
	}

    //Initialise an object for creation
	contact := Contact{
		FirstName: "Create",
		LastName:  "A",
		Phone:     "New",
		Email:     "Contact",
	}

    //The CreateSObject method's payload is an io.Reader type
    //so marshal the object into bytes
	b, err := json.Marshal(contact)
	if err != nil {
		log.Fatalln(err)
	}

    //And then create a new io.Reader
	payload := bytes.NewReader(b)
	if err != nil {
		log.Fatalln(err)
	}

    //Push the Contact object and it's payload to salesforce
	result, err := session.CreateSObject("Contact", payload)
	if err != nil {
		log.Fatalln(err)
	}

    //We're going to update this record, so fetch it's ID from the response
	contactID := result.ID

    //Recreate the object data to update with
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

    //We need to push the data to the individual object so we use it's ID
	contactEndpoint := fmt.Sprintf("Contact/%s", contactID)

    //Push the payload to update the object
	anotherResult, err := session.UpdateSObject(contactEndpoint, payload)

    //Let's get the object we just updated, and view it's contents
	response, err := session.GetSObject(contactEndpoint)
	if err != nil {
		log.Fatalln(err)
	}

	if response.TotalSize > 0 {
		contactData := response.Records[0]
		contactID := contactData["Id"].(string)
	}
}
```
