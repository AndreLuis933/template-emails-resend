# Email Notification - Go

Go implementation of the email notification system using Resend API.

## Installation

```bash
go mod download
```

## Usage

```go
package main

import (
    "fmt"
    "log"
)

func main() {
    // Success notification
    if err := NotifySuccess("My Job Name"); err != nil {
        log.Printf("Failed to send success notification: %v", err)
    }

    // Error notification
    err := fmt.Errorf("something went wrong")
    if notifyErr := NotifyError("My Job Name", err); notifyErr != nil {
        log.Printf("Failed to send error notification: %v", notifyErr)
    }
}
```

## Running

```bash
go run notification.go
```

## Environment Variables

Create a `.env` file in the same directory:

```env
RESEND_API_KEY=your_resend_api_key_here
SENDER_EMAIL=jobs@yourdomain.com
DESTINATION_EMAIL=your-email@example.com
GITHUB_TEMPLATE_URL=https://raw.githubusercontent.com/AndreLuis933/template-emails-resend/master/template
```
