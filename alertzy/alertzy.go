package alertzy

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
)

const url = "https://alertzy.app/send"

func SendNotification(accountKey string, title string, message string, priority int, group string) error {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	writer.WriteField("accountKey", accountKey)
	writer.WriteField("priority", fmt.Sprintf("%d", priority))
	writer.WriteField("group", group)
	writer.WriteField("title", title)
	writer.WriteField("message", message)
	err := writer.Close()
	if err != nil {
		return err
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	return nil
}
