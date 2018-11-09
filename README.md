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
* [Using docker](#using-docker)
* [Configuration](#configuration)
* [Usage](#usage)

## Features
* Written in Go - just a binary with 0 dependencies.
* Supports both MongoDB and PostgreSQL.
* Redirect your preferred way - [**client-side**](https://www.w3.org/TR/WCAG20-TECHS/H76.html) or [**server-side**](https://www.w3.org/TR/WCAG20-TECHS/SVR1.html).
* Shorten URLs to reduce usage.
* Count clicks on each shortened URL.

## Installation
Download the latest binary for Linux [here](https://github.com/himanshub16/outbound-go/releases/latest).
Configure and just run the binary.

### Using docker
1. Clone the repo 
2. Docker compose 
  ```
  docker-compose up --build
  ```
3. To change configurations update `docker.env`

## Configuration
`.env.postgresql.json`, `.env.mongodb.json` or `environment variables` is what you are looking for.

It first checks `CONFIG_FILE` environment variable for required file, and if not found fetches each environment variable.
The variables are described as under:

| Field           | Description                                                     | Example                     |
| ------          | -----------                                                     | -------                     |
| DBTYPE          | string : one of `postgresql` or `mongodb`                       | `mongodb`                   |
| DB_URL          | URL of MongoDB instance                                         | `mongodb://localhost:27017` |
| DB_NAME         | Name of database                                                | outbound                    |
| LINKS_COLL      | The collection which stores the links                           | `links`                     |
| COUNTER_COLL    | The collection which stores the counter                         | `counter`                   |
|                 |                                                                 |                             |
| PORT            |                                                                 | 9000                        |
| REDIRECT_METHOD | Default redirect method, one of `client-side` and `server-side` | `client-side`               |
| REQUIRE_SSL     | boolean : Use SSL (for Postgres, to support Heroku)             | true                        |


## Usage
* Create a new entry

Make a  post request to `/new` with `access_token` (_if required_) and `url`.
The result will be a new **Link** object with the `short_id` of the shortened URL.

* Redirect a shortened URL (`short_id`)
  - **Client side** : `/c/:short_id`
  - **Server side** : `/s/:short_id`
  - **Default** : `/r/:short_id`


## What does outbound mean?
* Outbound refers to traffic outside your domain/website.
* Websites log clicks to other domains for analytical purpose. Example, Facebook uses lm.facebook.com, Slack uses slack-redir.net, Twitter has t.co, etc.

This is similar to https://git.io


---
Liked this? Star this repo, or [Grab me a coffee.](https://github.com/himanshub16/outbound-go/raw/master/static/paytm.png)
