package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"
)

type Entry struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	URL     string `json:"url"`
	Content string `json:"content"`
}

type EntriesResponse struct {
	Entries []Entry `json:"entries"`
}

func main() {
	minifluxURL := os.Getenv("MINIFLUX_URL")
	minifluxUser := os.Getenv("MINIFLUX_USER")
	minifluxPass := os.Getenv("MINIFLUX_PASS")
	receiverEmail := os.Getenv("RECEIVER_EMAIL")
	gmailEmail := os.Getenv("GMAIL_EMAIL")
	gmailPassword := os.Getenv("GMAIL_PASSWORD")
	category := os.Getenv("CATEGORY")

	retrieveCategoryID(minifluxURL, minifluxUser, minifluxPass, category)

	unreadEntries := make([]Entry, 0)
	categoryEntries := fetchUnreadEntries(minifluxURL, minifluxUser, minifluxPass, category)
	unreadEntries = append(unreadEntries, categoryEntries...)

	if len(unreadEntries) == 0 {
		log.Println("No unread entries found")
		return
	}

	emailBody := formatEmailBody(unreadEntries)
	sendEmail(gmailEmail, gmailPassword, receiverEmail, emailBody)

	for _, entry := range unreadEntries {
		markEntryAsRead(minifluxURL, minifluxUser, minifluxPass, entry.ID)
	}
}

func retrieveCategoryID(minifluxURL, minifluxUser, minifluxPass, category string) string {
	client := &http.Client{}
	url := minifluxURL + "/v1/categories"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error fetching categories: %v", err)
	}
	req.SetBasicAuth(minifluxUser, minifluxPass)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error fetching categories: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error fetching entries. Status: %d", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bodyBytes)

	return string(bodyBytes)
}

func fetchUnreadEntries(minifluxURL, minifluxUser, minifluxPass, category string) []Entry {
	client := &http.Client{}
	url := minifluxURL + "/v1/entries?status=unread"
	if category != "" {
		url += "&category=" + category
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.SetBasicAuth(minifluxUser, minifluxPass)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error fetching entries: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error fetching entries. Status: %d", resp.StatusCode)
	}

	var entriesResponse EntriesResponse
	err = json.NewDecoder(resp.Body).Decode(&entriesResponse)
	if err != nil {
		log.Fatalf("Error decoding entries response: %v", err)
	}

	return entriesResponse.Entries
}

func formatEmailBody(entries []Entry) string {
	var buffer bytes.Buffer

	for _, entry := range entries {
		buffer.WriteString(fmt.Sprintf("<h2><a href=\"%s\">%s</a></h2>", entry.URL, entry.Title))
		buffer.WriteString("<hr>")
	}

	return buffer.String()
}

func sendEmail(gmailEmail, gmailPassword, toEmail, body string) {
	auth := smtp.PlainAuth("", gmailEmail, gmailPassword, "smtp.gmail.com")

	currentDate := time.Now().Format("2006-01-02")
	subject := fmt.Sprintf("📰 News Updates - %s", currentDate)

	fmt.Println("sending email to: ", toEmail)

	to := []string{toEmail}
	msg := []byte("To: <" + toEmail + ">\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=UTF-8" + "\r\n" +
		"\r\n" +
		body)

	err := smtp.SendMail("smtp.gmail.com:587", auth, gmailEmail, to, msg)
	if err != nil {
		log.Fatalf("Error sending email: %v", err)
	} else {
		log.Println("Email sent successfully")
	}
}

func markEntryAsRead(minifluxURL, minifluxUser, minifluxPass string, entryID int64) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/v1/entries/%d", minifluxURL, entryID)

	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.SetBasicAuth(minifluxUser, minifluxPass)
	req.Header.Set("Content-Type", "application/json")
	payload := `{"status": "read"}`
	req.Body = ioutil.NopCloser(strings.NewReader(payload))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error marking entry as read: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		log.Fatalf("Error marking entry as read. Status: %d /n %e", resp.StatusCode, resp.Body)
	}
}
