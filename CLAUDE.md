# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Python-based email notification system for job monitoring using the Resend API. It sends styled HTML email notifications for job success and failure events to devrium.cloud email addresses.

## Commands

### Running the application
```bash
python main.py
```

### Installing dependencies
This project uses `uv` for dependency management:
```bash
uv sync
```

## Architecture

### Core Components

**main.py** - Single-file application containing all functionality:
- `send_email()` - Sends emails via Resend API
- `load_template()` - Loads HTML email templates from the project root
- `notify_error()` - Sends error notifications with traceback
- `notify_success()` - Sends success notifications

### Email Templates

Two HTML email templates in the project root:
- **email_error.html** - Red-themed error notification template
- **email_success.html** - Green-themed success notification template

Both templates use:
- Placeholder variables: `{{JOB_NAME}}`, `{{TIMESTAMP}}`, `{{ERROR_MESSAGE}}`, `{{SUCCESS_MESSAGE}}`
- Dark theme styling with gradient backgrounds
- Table-based layout for email client compatibility
- DevRium Cloud branding

### Configuration

- **Environment variables**: `RESEND_API_KEY` must be set in `.env` file
- **Email sender**: Hardcoded to `jobs@devrium.cloud` in main.py:17
- **Email recipient**: Hardcoded to `seuemail@gmail.com` in main.py:17
- **Timestamps**: Uses UTC timezone formatted as `dd/mm/yyyy às HH:MM:SS`

### Template Variable Replacement

The system uses simple string replacement (not a templating engine):
```python
html = template.replace("{{PLACEHOLDER}}", value)
```

When modifying templates, ensure placeholder variables remain intact for proper substitution.
