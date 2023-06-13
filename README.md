# Miniflux Email Client
Sends RSS updates from miniflux to your email. 

This client is a minimal self-hosted version of mailbrew, briefcake and tacodigest. It doesn't require any server or email provider, but relies on free github action and gmail smtp.

**warning** to enable smtp support in GMAIL, you'll have to reduce security of that account. Please be sure to not use you're main account for that.

## Prerequisites

- A Miniflux instance with an account
- A Gmail account with "Less secure apps" access enabled or an App Password generated

## Setup

1. Fork this repo
2. Retrieve miniflux API token
3. Configure gmail smtp based on a [following guide](https://community.cloudflare.com/t/solved-how-to-use-gmail-smtp-to-send-from-an-email-address-which-uses-cloudflare-email-routing/382769/2)
4. Modify `.github/workflows/daily_run.yml` file
5. Add github/action secrets
