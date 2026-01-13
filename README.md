# Miniflux Email Client

A Miniflux Client that sends feed updates via email.

This is a perfect "no-cost" solution based on GitHub and SMTP for people who often forget to check RSS readers for updates but never miss emails.

Just Fork it! ©

## Background

While there are numerous "RSS-over-email" services available, such as Mailbrew, Briefcake, or Tacodigest, this project aims to recreate essential features of those services using Git, CI, and SMTP. In most cases, these tools can be used free of charge and are available with open-source GitHub/GitLab repositories.

## GitHub Actions
### Prerequisites

- GitHub repository with CI support
- A Miniflux instance with an account
- SMTP account (Gmail is recommended, but ensure it's not your main Gmail account if you choose to use it)

### Setup

1. Fork this repository
2. Retrieve your Miniflux API token
3. Configure Gmail SMTP by following this [guide](https://community.cloudflare.com/t/solved-how-to-use-gmail-smtp-to-send-from-an-email-address-which-uses-cloudflare-email-routing/382769/2)
4. Modify the `.github/workflows/runner.yml` file
5. Add GitHub Action secrets (SMTP_SERVER, SMTP_USERNAME, and SMTP_PASSWORD are required)

Every 3 months, GitHub will ask if you want to continue running daily actions. Click "Yes" if you wish to keep receiving updates.

## Self-Hosted

For self-hosted deployments, a container image is available at `ghcr.io/skatkov/miniflux-email-client`.

### Example Kubernetes CronJob

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: miniflux-email-client
type: Opaque
stringData:
  MINIFLUX_TOKEN: "your-miniflux-api-token"
  SMTP_USERNAME: "your-smtp-username"
  SMTP_PASSWORD: "your-smtp-password"
---
apiVersion: batch/v1
kind: CronJob
metadata:
  name: miniflux-email-client
spec:
  schedule: "0 8 * * *"  # Daily at 8 AM
  jobTemplate:
    spec:
      template:
        spec:
          restartPolicy: Never
          containers:
            - name: miniflux-email-client
              image: ghcr.io/skatkov/miniflux-email-client:latest
              envFrom:
                - secretRef:
                    name: miniflux-email-client
              env:
                - name: SEND_TO
                  value: "you@example.com"
                - name: MINIFLUX_URL
                  value: "https://your-miniflux-instance.com/"
                - name: CATEGORY
                  value: "Daily"
```
## Email

The email template is powered by the [Acorn framework](http://docs.thememountain.com/acorn/).
