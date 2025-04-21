# Chirpy

**My goals for this project were:**

- Understand what web servers are and how they power real-world web applications
- Build a production-style HTTP server in Go, without using any frameworks.
- Learn how to use JSON headers, and status codes to communicate with clients with RESTful API
- Learn what makes Go a great language for building fast web servers
- Use type safe SQL to store and retrieve data from a Postgres database
- Implement a secure application with authentication/authorization system with well-tested cryptography libraries
- Build and understand webhooks and API keys
- Document the RESTful API with markdown

## Installation

Inside a Go module:

```bash
    go get github.com/devambarbhaya/chirpy
```

## API URLs

### GET requests

- **GET /api/healthz** - This request will check the status of the API. If it is working or if it is down.

- **GET /api/chirps** - This request will return all the chirps in the database. It can take in two queries `author_id` and `sort`. You need to provide the exact uuid for the `author_id`. `sort` defaults to "asc" ascending time of creation of the chirp but also takes in "desc" for descending order of time of creation of the chirp

- **GET /api/chirps/{chirpID}** - This requests will return the data of the exact chirp which has the `chirpID` passed into the request URL.

### GET requests (admin privilege only)

- **GET /admin/metrics** - This request can only be sent by the admin and returns the number of times the website has been hit (called for).

### POST requests

- **POST /api/users** - This requests creates your user and stores it in the database.

- **POST /api/login** - This requests logs in your user.

- **POST /api/chirps** - This requests creates a chirp and stores it in the database. You need to be logged in for this request to go through.

- **POST /api/refresh** - This requests refreshes the token for your user.

- **POST /api/revoke** - This requests revokes the existing access token of your user and refreshes it with another.

- **POST /api/polka/webhooks** - This request will be continuously listening for responses from the polka webhooks. If the transaction will be successful the user will become a `chirpy_red` user.

### POST requests (admin privilege only)

- **POST /admin/reset** - This request resets the data in the database (deletes everything). Be careful when using this command.

### PUT requests

- **PUT /api/users** - This request will be used to update the user's credentials.

### DELETE requests

- **DELETE /api/chirps/{chirpID}** - This request will delete the chirp from the database.
