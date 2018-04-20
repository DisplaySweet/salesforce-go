# Salesforce-Go

A library to help consume Salesforce objects, useful when creating external integrations to push to or pull data from Salesforce.

To download the package:
```
> go get github.com/DisplaySweet/salesforce-go/src
```


Usage:
```
func main() {
    creds := salesforce.Credentials{
        Username: "test",
        Password: "secret_password",
    }

    session, err := salesforce.LoginByOAuth(creds)
    if err != nil {
        log.Fatalln(err)
    }

    response, err := session.GetSObject("Contact/AAAABBBCCCCqazc")
    if err != nil {
        log.Fatalln(err)
    }

    if response.TotalSize > 0 {
        contactData := response.Records[0]
        contactID := contactData["Id"].(string)
    }
}
```