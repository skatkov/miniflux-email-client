# Miniflux Email Client
Sends RSS updates from miniflux to your email. 

This client is a minimal self-hosted version of mailbrew, briefcake and tacodigest. It doesn't require any server or email provider, but relies on free github action and a gmail account.

## Prerequisites
- A Miniflux instance with an account
- A Gmail account, preferably not you're main account.

## Setup
1. Fork this repo
2. Retrieve miniflux API token
3. Configure gmail smtp based on a [following guide](https://community.cloudflare.com/t/solved-how-to-use-gmail-smtp-to-send-from-an-email-address-which-uses-cloudflare-email-routing/382769/2)
4. Modify `.github/workflows/daily_run.yml` file
5. Add github/action secrets (SMTP_SERVER, SMTP_USERNAME and SMTP_PASSWORD are required)

## Github Action
Every 3 months github will ask if you want to keep running daily actions. Click "yes", if you want to continue receiving updates.

## Maintenance
This is just a proof-of-concept, works for me kind of thing. I might improve this further, depending on my own usage.

Everyone is free to do whatever with this, if you plan to return back PR's -- I'll be happy to review.

Current plan includes:
- [ ] Generalize email configuration, so not only GMAIL could be used.
- [ ] Improve email design - add table of contents, improve template
- [ ] simplify configuration, now things are a bit all over the place
