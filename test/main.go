package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

// tokenFile is an implementation of the oauth2.TokenSource interface
// that stores and retrieves tokens from a file.
type tokenFile struct {
	filename string
}

func main() {
	// Set the path for the client secrets file and the token file.
	clientSecretsFile := "/path/to/client_secrets.json"
	tokenFile := "/path/to/token.json"

	// Set the scopes for the Gmail API.
	scopes := []string{gmail.GmailReadonlyScope}

	// Initialize the Gmail API client.
	gmailClient, err := gmail.NewService(context.Background())
	if err != nil {
		log.Fatalf("Unable to initialize Gmail client: %v", err)
	}

	// Initialize the OAuth2 config.
	config, err := google.ConfigFromJSON(readFile(clientSecretsFile), scopes...)
	if err != nil {
		log.Fatalf("Unable to read client secrets file: %v", err)
	}

	// Initialize the OAuth2 token store.
	tokenStore := newTokenFile(tokenFile)

	// Retrieve the access token from the token store or initiate the OAuth2 flow to obtain a new token.
	token, err := tokenStore.Token()
	if err != nil {
		token = getTokenFromWeb(config)
		err = tokenStore.Save(token)
		if err != nil {
			log.Fatalf("Unable to save token to file: %v", err)
		}
	}

	// Set the access token on the Gmail API client.
	gmailClient.UserAgent = "Gmail API Quickstart"
	gmailClient.BasePath = "https://www.googleapis.com/gmail/v1/"
	gmailClient.HttpClient = config.Client(context.Background(), token)

	// Use the Gmail API client to make requests.
	// For example, you could retrieve the user's messages like this:
	messages, err := gmailClient.Users.Messages.List("me").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve messages: %v", err)
	}
	for _, message := range messages.Messages {
		fmt.Printf("Message ID: %v\n", message.Id)
	}
}

// getTokenFromWeb initiates the OAuth2 flow to obtain a new access token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	token, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve access token: %v", err)
	}
	return token
}

// newTokenFile creates a new OAuth2 token file.
func newTokenFile(tokenFile string) *tokenFile {
	return &tokenFile{
		filename: tokenFile,
	}
}

// readFile reads a file and returns its contents as a byte slice.
func readFile(filename string) []byte {
	absPath, err := filepath.Abs(filename)
	if err != nil {
		log.Fatalf("Unable to determine absolute path for file %v: %v", filename, err)
	}
	data, err := os.ReadFile(absPath)
	if err != nil {
		log.Fatalf("Unable to read file %v: %v", filename, err)
	}
	return data
}

// save saves the token to the file.
func (f *tokenFile) Save(token *oauth2.Token) error {
	data, err := json.Marshal(token)
	if err != nil {
		return fmt.Errorf("Unable to marshal token: %v", err)
	}
	err = os.WriteFile(f.filename, data, 0600)
	if err != nil {
		return fmt.Errorf("Unable to write token file: %v", err)
	}
	return nil
}

// load loads the token from the file.
func (f *tokenFile) Token() (*oauth2.Token, error) {
	data, err := os.ReadFile(f.filename)
	if err != nil {
		return nil, fmt.Errorf("Unable to read token file: %v", err)
	}
	token := &oauth2.Token{}
	err = json.Unmarshal(data, token)
	if err != nil {
		return nil, fmt.Errorf("Unable to unmarshal token: %v", err)
	}
	return token, nil
}
