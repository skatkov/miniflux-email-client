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

## GitHub Action

Every 3 months, GitHub will ask if you want to continue running daily actions. Click "Yes" if you wish to keep receiving updates.

## Email

The email template is powered by the [Acorn framework](http://docs.thememountain.com/acorn/).
