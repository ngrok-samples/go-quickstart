# ngrok Go SDK Quickstart

A minimal Go app demonstrating the ngrok Go SDK.

## What you'll need

- An [ngrok account](https://dashboard.ngrok.com/signup).
- Your [ngrok auth token](https://dashboard.ngrok.com/get-started/your-authtoken).
- [Go installed](https://go.dev/doc/install) on your machine.

## Setup

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Create a `.env` file from the example:
   ```bash
   cp .env.example .env
   ```

3. Add your ngrok auth token to the `.env` file:
   ```txt
   NGROK_AUTHTOKEN=your_actual_authtoken_here
   ```

4. (Optional) Reserve a domain in the [ngrok dashboard](https://dashboard.ngrok.com/domains) and update the `ngrok.WithURL` call in `example.go`.

## Running the app

1. Start the Go service:
   ```bash
   go run service.go
   ```

2. In another terminal, start the ngrok agent endpoint:
   ```bash
   go run example.go
   ```

The ngrok agent endpoint will output a URL that forwards traffic to your local app. If you configured OAuth, visitors will need to log in with Google to access it.

## Files

- `service.go` - Basic Go HTTP server
- `example.go` - ngrok agent endpoint configuration with OAuth
- `.env.example` - Environment variable template
