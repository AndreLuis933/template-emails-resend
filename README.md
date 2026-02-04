# Email Notification Template - Resend

This is a reusable email notification template for job monitoring using the [Resend API](https://resend.com). Templates are fetched directly from GitHub, making it easy to copy and use in any project.

## Features

- ✅ Send styled HTML email notifications
- 🚨 Error notifications with full error stack trace
- 🎨 Beautiful dark-themed email templates
- 🔄 Templates fetched from GitHub (always up-to-date)
- 🐍 Python implementation
- 📘 TypeScript implementation (Bun runtime)
- 🔵 Go implementation

## Quick Start

### 1. Copy the Code

Choose your preferred language and copy the corresponding file to your project:

**Python:**
```bash
# Copy the notification module
cp email-py/notification.py your-project/
```

**TypeScript (Bun):**
```bash
# Copy the index file
cp email-ts/index.ts your-project/
```

**Go:**
```bash
# Copy the notification module
cp email-go/notification.go your-project/
cp email-go/go.mod your-project/
cp email-go/go.sum your-project/
```

### 2. Set Up Environment Variables

Create a `.env` file in your project root:

```env
RESEND_API_KEY=your_resend_api_key_here
SENDER_EMAIL=jobs@yourdomain.com
DESTINATION_EMAIL=your-email@example.com
GITHUB_TEMPLATE_URL=https://raw.githubusercontent.com/AndreLuis933/template-emails-resend/master/template
```

### 3. Install Dependencies

**Python:**
```bash
pip install requests python-dotenv
```

**TypeScript (Bun):**
```bash
bun add dotenv
```

**Go:**
```bash
go mod download
```

### 4. Usage

**Python:**
```python
from notification import notify_success, notify_error

try:
    # Your job code here
    notify_success("My Job Name")
except Exception:
    notify_error("My Job Name")
    raise
```

**TypeScript (Bun):**
```typescript
import { notifySuccess, notifyError } from "./index";

try {
    // Your job code here
    await notifySuccess("My Job Name");
} catch (error) {
    await notifyError("My Job Name", error as Error);
    throw error;
}
```

**Go:**
```go
package main

import (
    "fmt"
    "log"
)

func main() {
    // Your job code here
    err := someFunction()
    if err != nil {
        if notifyErr := NotifyError("My Job Name", err); notifyErr != nil {
            log.Printf("Failed to send error notification: %v", notifyErr)
        }
        panic(err)
    }

    if err := NotifySuccess("My Job Name"); err != nil {
        log.Printf("Failed to send success notification: %v", err)
    }
}
```

## Email Templates

The project uses two HTML email templates hosted on GitHub:

- **email_error.html** - Red-themed error notification with stack trace
- **email_success.html** - Green-themed success notification

Templates support the following placeholders:
- `{{JOB_NAME}}` - Name of the job
- `{{TIMESTAMP}}` - Execution timestamp (dd/mm/yyyy às HH:MM:SS)
- `{{ERROR_MESSAGE}}` - Error message and stack trace (error template only)

## Customization

### Use Your Own Templates

To use your own template repository:

1. Fork this repository or create your own with the email templates
2. Update the `GITHUB_TEMPLATE_URL` in your `.env` file:
   ```env
   GITHUB_TEMPLATE_URL=https://raw.githubusercontent.com/YOUR_USERNAME/YOUR_REPO/BRANCH/template
   ```

### Modify Email Content

The templates use simple string replacement. You can add custom placeholders by:

1. Adding `{{YOUR_PLACEHOLDER}}` in the HTML template
2. Replacing it in the code:
   ```python
   html = html.replace("{{YOUR_PLACEHOLDER}}", "your_value")
   ```

## Project Structure

```
.
├── email-py/              # Python implementation
│   └── notification.py    # Main notification module
├── email-ts/              # TypeScript implementation
│   └── index.ts          # Main notification module (Bun)
├── email-go/              # Go implementation
│   ├── notification.go   # Main notification module
│   ├── go.mod           # Go module file
│   └── go.sum           # Go dependencies
├── template/             # Email HTML templates
│   ├── email_error.html  # Error notification template
│   └── email_success.html # Success notification template
├── .env                  # Environment variables
└── README.md            # This file
```

## Why Templates from GitHub?

Fetching templates from GitHub provides several benefits:

1. **Always Up-to-Date**: Get the latest template improvements automatically
2. **Easy to Copy**: Just copy one file to your project
3. **Centralized Management**: Update templates in one place, all projects benefit
4. **No Local Files**: Reduces project clutter

## License

Feel free to copy and use this in your projects!

## Support

For issues or questions, please open an issue on the [GitHub repository](https://github.com/AndreLuis933/template-emails-resend).
