# Miniflux Email Client

A Miniflux Client that sends feed updates via email.

This is a perfect "no-cost" solution based on GitHub and SMTP for people who often forget to check RSS readers for updates but never miss emails.

Just Fork it! ©

## Background

While there are numerous "RSS-over-email" services available, such as Mailbrew, Briefcake, or Tacodigest, this project aims to recreate essential features of those services using Git, CI, and SMTP. In most cases, these tools can be used free of charge and are available with open-source GitHub/GitLab repositories.

## Prerequisites

- GitHub repository with CI support
- A Miniflux instance with an account
- SMTP account (Gmail is recommended, but ensure it's not your main Gmail account if you choose to use it)

## Setup

1. Fork this repository
2. Retrieve your Miniflux API token
3. Configure Gmail SMTP by following this [guide](https://community.cloudflare.com/t/solved-how-to-use-gmail-smtp-to-send-from-an-email-address-which-uses-cloudflare-email-routing/382769/2)
4. Modify the `.github/workflows/runner.yml` file
5. Add GitHub Action secrets (SMTP_SERVER, SMTP_USERNAME, and SMTP_PASSWORD are required)

## Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `MINIFLUX_URL` | No | `https://reader.miniflux.app/` | Miniflux instance URL |
| `MINIFLUX_TOKEN` | Yes | - | Miniflux API token |
| `CATEGORY` | No | - | Filter entries by category name |
| `LIMIT` | No | - | Maximum number of entries to fetch |
| `SMTP_SERVER` | No | `smtp.gmail.com` | SMTP server hostname |
| `SMTP_PORT` | No | `587` | SMTP server port |
| `SMTP_USERNAME` | Yes | - | SMTP authentication username |
| `SMTP_PASSWORD` | Yes | - | SMTP authentication password |
| `SEND_FROM` | No | `SMTP_USERNAME` | Email sender address |
| `SEND_TO` | Yes | - | Email recipient address |

## GitHub Action

Every 3 months, GitHub will ask if you want to continue running daily actions. Click "Yes" if you wish to keep receiving updates.

## Email

The email template is powered by the [Acorn framework](http://docs.thememountain.com/acorn/).
