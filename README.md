# Go Starter Template

This is a starter template for building web applications using Go, Pocketbase, and HTMX. It provides a simple and efficient setup to get you started quickly.

## Features

- **Pocketbase**: A powerful backend service for your application.
- **HTMX**: A modern library for building dynamic web applications with minimal JavaScript.
- **Tailwind CSS**: A utility-first CSS framework for rapid UI development.

## Prerequisites

Before you begin, ensure you have the following tools installed on your system:

1. [Go](https://golang.org/): The Go programming language, required to build and run the application.

## Installation

Follow these steps to set up the project:

1. **Install required tools:**

   ```sh
   go install github.com/air-verse/air@latest
   go install github.com/a-h/templ/cmd/templ@latest
   go install github.com/hookenz/gotailwind/v4@latest
   ```

2. **Download project dependencies:**
   ```sh
   go mod tidy
   ```

## Running the Project

To start the development server, run the following command:

```sh
air
```

Once the server is running, open your browser and navigate to:

[http://127.0.0.1:8090/hello/yourname](http://127.0.0.1:8090/hello/yourname)

## TODO

### Authentication

1. **HTMX Configuration:** Use `htmx:config` to set the authentication header on the client side. For more details, check out this [Reddit discussion](https://www.reddit.com/r/htmx/comments/187n06e/comment/kd5ug8c/).

2. **Echo Middleware:** Use Echo middleware to handle claims and set cookies. Learn more in this [example project](https://git.sunshine.industries/efim/go-ssr-pocketbase-oauth-attempt#headline-12).
