# ScanWP - WordPress Vulnerability Scanning Tool

## Introduction
ScanWP is an open-source scanning tool designed to identify potential security vulnerabilities in WordPress sites.

## Usage
You can use the tool to scan your WordPress site and get reports on vulnerabilities. You can also perform a detailed scan on a specific URL address and examine the results in depth.

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/tayfunurkut/scanwp.git
   cd scanwp
   go build .

## Usage
1. `-h, --help`              help for scanwp
2. `-s, --scanwp`            Scanning parameter (default true)
3. `-t, --target string`     Target Scan
4. `-w, --wordlist string`   Wordlist (default "wordlist/test.txt")

## Configuration

Before running GoScaner, make sure to set up your API key:

1. Create a `.env` file in the project directory.
2. Add your API key to the `.env` file:

## Examples

1. `./scanpw -t https://example.com -w /usr/share/wordlist/wordlist.txt -s` 
