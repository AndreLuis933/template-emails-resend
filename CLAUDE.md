# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a reusable email notification template system for job monitoring using the Resend API. It provides three language implementations (Python, TypeScript, Go) that send styled HTML email notifications for job success and failure events. Templates are fetched from GitHub for easy centralized management and reusability across projects.

## Project Structure

```
.
├── email-py/              # Python implementation
│   └── notification.py    # Main notification module
├── email-ts/              # TypeScript implementation (Bun)
│   ├── index.ts          # Main notification module
│   └── package.json      # Bun dependencies
├── email-go/              # Go implementation
│   ├── notification.go   # Main notification module
│   ├── go.mod           # Go module file
│   └── go.sum           # Go dependencies
├── template/             # Email HTML templates
│   ├── email_error.html  # Error notification template
│   └── email_success.html # Success notification template
├── .env                  # Environment variables (root)
└── README.md            # Project documentation
```

## Commands

### Python Implementation
```bash
cd email-py
python notification.py
```

### TypeScript Implementation (Bun)
```bash
cd email-ts
bun run index.ts
```

### Go Implementation
```bash
cd email-go
go run notification.go
```

### Installing Dependencies

**Python:**
```bash
cd email-py
pip install requests python-dotenv
```

**TypeScript (Bun):**
```bash
cd email-ts
bun install
```

**Go:**
```bash
cd email-go
go mod download
```

## Architecture

### Core Components

All three implementations provide the same core functionality:

**Python (`email-py/notification.py`):**
- `send_email()` - Sends emails via Resend API
- `load_template()` - Fetches HTML templates from GitHub
- `notify_error(job_name)` - Sends error notifications with traceback
- `notify_success(job_name)` - Sends success notifications

**TypeScript (`email-ts/index.ts`):**
- `sendEmail()` - Sends emails via Resend API
- `loadTemplate()` - Fetches HTML templates from GitHub
- `notifyError(jobName, error)` - Sends error notifications with stack trace
- `notifySuccess(jobName)` - Sends success notifications

**Go (`email-go/notification.go`):**
- `SendEmail()` - Sends emails via Resend API
- `loadTemplate()` - Fetches HTML templates from GitHub
- `NotifyError(jobName, err)` - Sends error notifications with stack trace
- `NotifySuccess(jobName)` - Sends success notifications

### Email Templates

Templates are hosted on GitHub and fetched at runtime:
- **email_error.html** - Red-themed error notification template with stack trace
- **email_success.html** - Green-themed success notification template

Both templates use:
- Placeholder variables: `{{JOB_NAME}}`, `{{TIMESTAMP}}`, `{{ERROR_MESSAGE}}`
- Dark theme styling with gradient backgrounds
- Table-based layout for email client compatibility
- DevRium Cloud branding

Templates are fetched from: `https://raw.githubusercontent.com/AndreLuis933/template-emails-resend/master/template/`

### Configuration

All implementations use environment variables from `.env` file:
- `RESEND_API_KEY` - Resend API key for authentication
- `SENDER_EMAIL` - Email address to send from (e.g., `jobs@devrium.cloud`)
- `DESTINATION_EMAIL` - Email address to send to
- `GITHUB_TEMPLATE_URL` - Base URL for GitHub-hosted templates

**Example `.env` file:**
```env
RESEND_API_KEY=re_YourApiKeyHere
SENDER_EMAIL=jobs@devrium.cloud
DESTINATION_EMAIL=your-email@example.com
GITHUB_TEMPLATE_URL=https://raw.githubusercontent.com/AndreLuis933/template-emails-resend/master/template
```

### Template Variable Replacement

All implementations use simple string replacement (not a templating engine):

**Python:**
```python
html = template.replace("{{PLACEHOLDER}}", value)
```

**TypeScript:**
```typescript
html = template.replace("{{PLACEHOLDER}}", value)
```

**Go:**
```go
html = strings.ReplaceAll(template, "{{PLACEHOLDER}}", value)
```

When modifying templates, ensure placeholder variables remain intact for proper substitution.

### Timestamp Format

All implementations format timestamps consistently:
- Format: `dd/mm/yyyy às HH:MM:SS` (Brazilian Portuguese format)
- Example: `04/02/2026 às 14:30:45`

### Error Handling

**Python:**
- Uses `traceback.format_exc()` to capture full error stack trace
- Errors are sent as formatted text in the email

**TypeScript:**
- Uses `error.stack` or `error.message` to capture error details
- Error objects must be passed to `notifyError()`

**Go:**
- Uses `runtime/debug.Stack()` to capture full stack trace
- Errors include both error message and stack trace

## Design Principles

1. **Single File per Implementation** - Each implementation is self-contained in a single file for easy copying to other projects

2. **No Wildcard Imports** - TypeScript uses named imports only (`import { config }` not `import *`)

3. **GitHub-Based Templates** - Templates are always fetched from GitHub to ensure:
   - Consistency across projects
   - Easy updates without changing code
   - No local file dependencies

4. **Environment-Based Configuration** - All configuration via `.env` file for security and flexibility

5. **Consistent API** - All three implementations provide the same functions with similar signatures

## Usage as a Template

This project is designed to be easily copied to other projects:

1. Copy the implementation file for your language
2. Copy the `.env` file and update values
3. Install dependencies
4. Use `notify_error()` and `notify_success()` in your code

No need to copy HTML templates - they're fetched from GitHub automatically.
