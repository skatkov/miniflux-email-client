# Miniflux Email Updates

This Go application fetches the latest unread entries from your Miniflux instance and sends them as an email using Gmail's SMTP server.

## Prerequisites

- A Miniflux instance with an account
- A Gmail account with "Less secure apps" access enabled or an App Password generated
- [GoReleaser](https://goreleaser.com/) (for releasing)

## Setup

1. Clone the repository:

```bash
git clone https://github.com/your-username/your-repo.git
cd your-repo

    Set the required environment variables:

bash

export MINIFLUX_URL=https://your-miniflux-instance.com
export MINIFLUX_USER=your-miniflux-username
export MINIFLUX_PASS=your-miniflux-password
export GMAIL_EMAIL=your-gmail-email@example.com
export GMAIL_PASSWORD=your-gmail-password
export CATEGORIES=news,technology

Note: You can also set the environment variables directly in the Go application or use a .env file and load them using a package like godotenv.

    Build the application:

bash

go build

    Run the application:

bash

./miniflux-email-updates

Running the Application Daily
Locally
Linux and macOS

Add a new line to the crontab with the following format:

javascript

MM HH * * * /path/to/launch.sh

Replace MM with the minute (00-59) and HH with the hour (00-23) you want the application to run every day.
Windows

Create a new task in the Task Scheduler to run the launch.sh script at the desired time every day.
Using GitHub Actions

    Create a new GitHub Actions workflow in your repository by following the instructions in the GitHub Actions section.

    Store your sensitive information (credentials) as GitHub secrets.

    The GitHub Action will run the Go application every day at the specified time.

GitHub Actions

To run the Go application in a GitHub Action every day at a specific time, create a new workflow file in your repository:

    In your repository, create a directory named .github (if it doesn't already exist).
    Inside the .github directory, create another directory named workflows.
    Inside the workflows directory, create a new file named daily_run.yml.

Add the following content to the daily_run.yml file:

yaml

name: Daily Run

on:
  schedule:
    - cron: '0 9 * * *'

jobs:
  run_app:
    runs-on: ubuntu-latest
    steps:
    - name: Check out repository
      uses: actions/checkout@v2

    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: 1.17

    - name: Build and run the application
      env:
        MINIFLUX_URL: ${{ secrets.MINIFLUX_URL }}
        MINIFLUX_USER: ${{ secrets.MINIFLUX_USER }}
        MINIFLUX_PASS: ${{ secrets.MINFLUX_PASS }}
        GMAIL_EMAIL: ${{ secrets.GMAIL_EMAIL }}
        GMAIL_PASSWORD: ${{ secrets.GMAIL_PASSWORD }}
        CATEGORIES: ${{ secrets.CATEGORIES }}
      run: |
        go build
        ./miniflux-email-updates

This GitHub Action will run the Go application every day at the specified time (UTC). Make sure to replace the Go version in the setup-go action with the version you are using in your project, if different from 1.17.

Additionally, store your sensitive information (credentials) as GitHub secrets. To do this, follow these steps:

    In your GitHub repository, click the "Settings" tab.
    In the left sidebar, click "Secrets".
    Click the "New repository secret" button.
    Add each environment variable (e.g., MINIFLUX_URL, MINIFLUX_USER, MINIFLUX_PASS, GMAIL_EMAIL, GMAIL_PASSWORD, and CATEGORIES) as a separate secret with the corresponding value.

Once you have set up the secrets, the GitHub Action will use them when running the workflow.
Releasing with GoReleaser

To release the Go application using GoReleaser, follow these steps:

    Install GoReleaser and UPX.
    Update the goreleaser.yml file with the appropriate information for your repository.
    Create a new Git tag for the release:

bash

git tag -a vX.Y.Z -m "Release vX.Y.Z"
git push origin vX.Y.Z

    Run GoReleaser:

bash

goreleaser release --rm-dist

This will create a new release, compress the binary using UPX, and publish it to your personal Homebrew tap.
