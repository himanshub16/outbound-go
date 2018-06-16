# Outbound

Create and manage outbound links and shorten URLs without hassle.

**Use cases**
* Host your own **URL shortener**
* Log clicks to outbound links

---

Table of Contents

* [What does outbound mean?](#what-does-outbound-mean)
* [Features](#features)
* [Installation](#installation)
* [Configuration](#configuration)

## Features
* Written in Go - just a binary with 0 dependencies.
* Uses MongoDB for write-intensive storage.
* Redirect your preferred way - [**client-side**](https://www.w3.org/TR/WCAG20-TECHS/H76.html) or [**server-side**](https://www.w3.org/TR/WCAG20-TECHS/SVR1.html).
* Shorten URLs to reduce usage.
* Count clicks on each shortened URL.

## Installation
Download the latest binary for Linux [here](https://github.com/himanshub16/outbound-go/releases/latest).
Configure and just run the binary.

## Configuration
`env.json` or `environment variables` is what you are looking for.

It first checks `CONFIG_FILE` environment variable for required file, and if not found fetches each environment variable.
The variables are described as under:

| Field           | Description                                                     | Example                     |
| ------          | -----------                                                     | -------                     |
| DB_URL          | URL of MongoDB instance                                         | `mongodb://localhost:27017` |
| DB_NAME         | Name of database                                                | outbound                    |
| LINKS_COLL   | The collection which stores the links                           | `links`                     |
| COUNTER_COLL    | The collection which stores the counter                         | `counter`                   |
|                 |                                                                 |                             |
| PORT            |                                                                 | 9000                        |
| REDIRECT_METHOD | Default redirect method, one of `client-side` and `server-side` | `client-side`               |


## What does outbound mean?
* Outbound refers to traffic outside your domain/website.
* Websites log clicks to other domains for analytical purpose. Example, Facebook uses lm.facebook.com, Slack uses slack-redir.net, Twitter has t.co, etc.

This is similar to https://git.io


---
Liked this? Star this repo, or [Grab me a coffee.](https://github.com/himanshub16/outbound-go/raw/master/static/paytm.png)
