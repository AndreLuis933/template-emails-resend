package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var (
	ResendAPIKey      string
	DestinationEmail  string
	SenderEmail       string
	GithubTemplateURL string
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	ResendAPIKey = os.Getenv("RESEND_API_KEY")
	DestinationEmail = os.Getenv("DESTINATION_EMAIL")
	SenderEmail = os.Getenv("SENDER_EMAIL")
	GithubTemplateURL = os.Getenv("GITHUB_TEMPLATE_URL")
}

type EmailRequest struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	HTML    string   `json:"html"`
}

type EmailResponse struct {
	ID string `json:"id"`
}

// SendEmail sends an email using the Resend API
func SendEmail(subject, htmlContent string) (*EmailResponse, error) {
	reqBody := EmailRequest{
		From:    SenderEmail,
		To:      []string{DestinationEmail},
		Subject: subject,
		HTML:    htmlContent,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.resend.com/emails", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+ResendAPIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var emailResp EmailResponse
	if err := json.NewDecoder(resp.Body).Decode(&emailResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &emailResp, nil
}

// loadTemplate fetches an email template from GitHub
func loadTemplate(templateName string) (string, error) {
	templateURL := fmt.Sprintf("%s/%s", GithubTemplateURL, templateName)

	resp, err := http.Get(templateURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch template: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch template: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read template: %w", err)
	}

	return string(body), nil
}

// formatTimestamp formats the current time in the Brazilian format
func formatTimestamp() string {
	now := time.Now()
	return now.Format("02/01/2006 às 15:04:05")
}

// NotifyError sends an error notification email
func NotifyError(jobName string, err error) error {
	template, loadErr := loadTemplate("email_error.html")
	if loadErr != nil {
		return fmt.Errorf("failed to load error template: %w", loadErr)
	}

	timestamp := formatTimestamp()

	// Get stack trace
	errorMsg := fmt.Sprintf("%v\n\n%s", err, string(debug.Stack()))

	// Replace placeholders
	html := strings.ReplaceAll(template, "{{JOB_NAME}}", jobName)
	html = strings.ReplaceAll(html, "{{TIMESTAMP}}", timestamp)
	html = strings.ReplaceAll(html, "{{ERROR_MESSAGE}}", errorMsg)

	subject := fmt.Sprintf("🚨 %s - Falhou", jobName)

	_, sendErr := SendEmail(subject, html)
	return sendErr
}

// NotifySuccess sends a success notification email
func NotifySuccess(jobName string) error {
	template, err := loadTemplate("email_success.html")
	if err != nil {
		return fmt.Errorf("failed to load success template: %w", err)
	}

	timestamp := formatTimestamp()

	// Replace placeholders
	html := strings.ReplaceAll(template, "{{JOB_NAME}}", jobName)
	html = strings.ReplaceAll(html, "{{TIMESTAMP}}", timestamp)

	subject := fmt.Sprintf("✅ %s - Sucesso", jobName)

	_, sendErr := SendEmail(subject, html)
	return sendErr
}

func main() {
	// Example usage with error
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("panic: %v", r)
			if notifyErr := NotifyError("Job Mensal Das", err); notifyErr != nil {
				log.Printf("Failed to send error notification: %v", notifyErr)
			}
			panic(r)
		}
	}()

	// Simulate an error
	if true {
		err := fmt.Errorf("erro proposital para testar notificação")
		if notifyErr := NotifyError("Job Mensal Das", err); notifyErr != nil {
			log.Printf("Failed to send error notification: %v", notifyErr)
		}
		panic(err)
	}

	// Success case
	if err := NotifySuccess("Job Mensal Das"); err != nil {
		log.Printf("Failed to send success notification: %v", err)
	}
}
