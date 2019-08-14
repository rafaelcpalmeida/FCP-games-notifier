package src

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type twilio struct {
	accountSid string
	authToken string
	urlStr string
}

type message struct {
	from string
	to string
	body string
}

func SendSMS(recipientNumber string, messageBody string) {
	smsSender, err := smsProvider()

	if err != nil {
		log.Fatal("something went wrong. Error: " + err.Error())
	}

	message := message{
		from: "+46790645794",
		to:   recipientNumber,
		body: messageBody,
	}

	prepareSMSDetails(message, smsSender)
}

func smsProvider() (twilio, error) {
	if len(os.Getenv("TWILIO_ACCOUNT_SID")) == 0 || len(os.Getenv("TWILIO_ACCOUNT_SID")) == 0 {
		return twilio{}, errors.New("one or more environment variables not found")
	}

	return twilio{
		accountSid: os.Getenv("TWILIO_ACCOUNT_SID"),
		authToken:  os.Getenv("TWILIO_ACCOUNT_AUTH_TOKEN"),
		urlStr:     fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", os.Getenv("TWILIO_ACCOUNT_SID")),
	}, nil
}

func prepareSMSDetails(messageData message, smsSender twilio) {
	// Pack up the data for our message
	msgData := url.Values{}
	msgData.Set("To", messageData.to)
	msgData.Set("From", messageData.from)
	msgData.Set("Body", messageData.body)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client, req := prepareTwilioRequest(smsSender, &msgDataReader)

	submitRequestToTwilio(client, req)
}

func prepareTwilioRequest(smsSender twilio, messageDetails io.Reader) (*http.Client, *http.Request) {
	// Create HTTP request client
	client := &http.Client{}
	req, _ := http.NewRequest("POST", smsSender.urlStr, messageDetails)
	req.SetBasicAuth(smsSender.accountSid, smsSender.authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return client, req
}

func submitRequestToTwilio(client *http.Client, req *http.Request) {
	// Make HTTP POST request and return message SID
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)

		if err == nil {
			log.Print(data["status"])
		}
	} else {
		body, _ := ioutil.ReadAll(resp.Body)

		log.Fatal(fmt.Sprintf("Error code: %s\nMessage: %s", resp.Status, string(body)))
	}
}