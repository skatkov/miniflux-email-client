# Miniflux Email Client
Miniflux Cliet to send feed updates over email.

Fork it and use it!

## Background
There are plenty of "RSS-over-email" services out there - mailbrew, briefcake or tacodigest as example. This project aims to rebuild essential features of those services, but using git + CI + SMTP. In most cases, these things could be used free of charge and no need to selfhost if you don't want to.

## Prerequisites
- Github repo + CI support
- A Miniflux instance with an account
- SMTP account (I use gmail, but make sure it's not you're main gmail account if you do)

## Setup
1. Fork this repo
2. Retrieve miniflux API token
3. Configure gmail smtp based on a [following guide](https://community.cloudflare.com/t/solved-how-to-use-gmail-smtp-to-send-from-an-email-address-which-uses-cloudflare-email-routing/382769/2)
4. Modify `.github/workflows/daily_run.yml` file
5. Add github/action secrets (SMTP_SERVER, SMTP_USERNAME and SMTP_PASSWORD are required)

## Github Action
Every 3 months github will ask if you want to keep running daily actions. Click "yes", if you want to continue receiving updates.

## Maintenance
I use this project myself daily, it's not feature complete, but it works stably. I'd encourage leaving your feedback and contributions.

Current plan includes:
- [x] Generalize email configuration, so not only GMAIL could be used.
- [ ] Improve email design - add table of contents, improve template
- [x] Introduce PLAIN TEXT emails
- [ ] Ability to provide multiple categories
- [ ] Ability to not provide any categories at all
- [ ] Ability to limit number of RSS updates sent out daily
- [x] make project easier to maintain (simplify, write tests and etc)
