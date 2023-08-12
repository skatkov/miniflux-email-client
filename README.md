# Miniflux Email Client
Sends RSS updates from miniflux to personal email. 

There are services that allow to send RSS updates through email, like mailbrew, briefcake or tacodigest. This is a miniflux-client that does something similar, but doesn't need to be hosted. 

This project uses github actions and any smtp account available. In most cases, these things could be used free of charge.

Just fork away and get one for you're personal enjoyment.

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
It initially started as a proof-of-concept project. But it kinda works, I'm trying to get this miniflux client to some "done" state.

Any contributions or ideas are welcome.

Current plan includes:
- [x] Generalize email configuration, so not only GMAIL could be used.
- [ ] Improve email design - add table of contents, improve template
- [x] Introduce PLAIN TEXT emails
- [ ] simplify configuration for others, now things are a bit all over the place
- [ ] make project easier to maintain (simplify, write tests and etc)
