# Miniflux Email Client
Sends RSS updates from miniflux to your email. 

This client is a minimal self-hosted version of mailbrew, briefcake and tacodigest. It doesn't require any server or email provider, but relies on free github action and a gmail account.

## Prerequisites
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
It initially started as a proof-of-concept project. Since concept works, i'm trying to get this miniflux client to some "done" state.

Any contributions or ideas are welcome.

Current plan includes:
- [x] Generalize email configuration, so not only GMAIL could be used.
- [ ] Improve email design - add table of contents, improve template
- [x] Introduce PLAIN TEXT emails
- [ ] simplify configuration for others, now things are a bit all over the place
- [ ] make project easier to maintain (simplify, write tests and etc)
