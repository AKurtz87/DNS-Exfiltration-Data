# Go DNS Exfiltration Scripts

This repository contains two Go scripts designed to demonstrate DNS-based data exfiltration techniques. These scripts simulate data exfiltration over DNS queries and responses using a local DNS server.

## Table of Contents
- [Overview](#overview)
- [How It Works](#how-it-works)
  - [DNS Server Script (`server.go`)](#dns-server-script-servergo)
  - [DNS Client Script (`client.go`)](#dns-client-script-clientgo)
- [Setup and Requirements](#setup-and-requirements)
- [File Structure](#file-structure)
- [Testing the Scripts](#testing-the-scripts)
- [Security Considerations](#security-considerations)
- [License](#license)

## Overview

This project demonstrates a basic proof-of-concept (PoC) for exfiltrating data using DNS requests. The DNS server script listens for DNS queries and extracts portions of the query name, writing them to a log file (`exfil.txt`). The client script reads sensitive data from a file (`secrets.txt`), splits it into smaller chunks, and sends each chunk as part of a DNS query to the server.

## How It Works

### DNS Server Script (`server.go`)

1. The script reads domain and IP address pairs from a file called `host.txt`.
2. It runs a DNS server on `127.0.0.1:53` that listens for DNS queries.
3. When a DNS request is received, it:
   - Extracts and processes the domain name from the request.
   - Logs specific parts of the domain name to `exfil.txt`.
   - Responds to queries with the IP address associated with the requested domain, if available in `host.txt`.

### DNS Client Script (`client.go`)

1. The script reads data from a file called `secrets.txt`.
2. It splits the content of the file into chunks of 10 bytes each.
3. For each chunk:
   - The script constructs a DNS query, appending the chunk to a domain name.
   - It sends the DNS query to the server at `127.0.0.1:53`.
   - The server processes the request and logs the chunks to `exfil.txt`.
