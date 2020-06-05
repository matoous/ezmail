# Ezmail

Ezmail is easy CLI utility for sending emails from terminal useful for scripts and people
that don't want to setup everything up for the default `mail` command to work.

## Features

* Support for multiple accounts
* Saves passwords into OS keychain

## Usage

1. Add an email account using `ezmail accounts add`, you will be prompted for address, password
and smtp server.

2. Send email using `ezmail`, for example `echo "Hello!" | ezmail recipient@gmail.com -s "Subject"`
