# Gator

Gator is a lightweight blog aggregator written in Go with a PostgreSQL backend. It allows users to register, log in, manage blog feeds, aggregate posts on a schedule, and browse saved posts.

---

## Features

- User registration and login
- Add and follow RSS/Atom blog feeds
- Aggregate posts from followed feeds at a configurable interval
- Browse saved blog posts
- Unfollow feeds

---

## Requirements

- Go 1.16 or newer
- PostgreSQL 12 or newer

---

## Installing PostgreSQL

### macOS

```sh
# Install Homebrew if needed
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# Install PostgreSQL
brew install postgresql

# Start PostgreSQL
brew services start postgresql
```

### Ubuntu / Debian

```sh
# Update package list
sudo apt update

# Install PostgreSQL
sudo apt install postgresql postgresql-contrib

# Start the PostgreSQL service
sudo service postgresql start
```

### Windows

1. Download the installer from: https://www.postgresql.org/download/windows/
2. Run the installer and follow the prompts.
3. Start PostgreSQL via the Start Menu or Services console.
4. Optionally use pgAdmin for managing the database.

---

## Installing Gator

Install the Gator CLI tool with `go install`:

```sh
go install github.com/warrco/gator@latest
```

Make sure `$GOPATH/bin` is in your system `PATH` so the `gator` command is available globally.

---

## Database Setup

1. Open a terminal and access the PostgreSQL shell:

```sh
sudo -u postgres psql
```

2. Create the Gator database:

```sql
CREATE DATABASE gatordb;
```

3. Exit the shell:

```sql
\q
```

Ensure that your application connects to this database. 

4. Create a .gatorconfig.json file in your home directory with the following contents:

{"db_url":"<postgresql connection string>"}

Connection string format may vary based on your operating system.

Some examples are as follows:

```
    macOS (no password, your username): postgres://wagslane:@localhost:5432/gator
    Linux (password from last lesson, postgres user): postgres://postgres:postgres@localhost:5432/gator
```

---

## Usage

Here are the available commands:

### Register a new user

```sh
gator register <username>
```

Registers a new user in the database.

---

### Log in as an existing user

```sh
gator login <username>
```

Logs in using an existing user account.

---

### Add a new feed and follow it

```sh
gator addfeed <feed name> <url>
```

Adds a feed to the database and follows it for the logged-in user.

---

### Aggregate followed feeds

```sh
gator agg <timeframe>
```

Aggregates all followed feeds and scrapes them periodically.

- Timeframes can be specified like `1s`, `30s`, `5m`, `1h`, etc.

---

### Show followed feeds

```sh
gator following
```

Displays the feeds the currently logged-in user is following.

---

### Unfollow a feed

```sh
gator unfollow <url>
```

Removes a feed from the logged-in user's follow list.

---

### Browse saved posts

```sh
gator browse <number of posts>
```

Displays the specified number of recent posts.

---