package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.ngrok.com/ngrok/v2"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file", err)
	}
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

const trafficPolicy = `
on_http_request:
  - actions:
      - type: oauth
        config:
          provider: google
`
const address = "http://localhost:8080"

func run(ctx context.Context) error {
	agent, err := ngrok.NewAgent(ngrok.WithAuthtoken(os.Getenv("NGROK_AUTHTOKEN")))
	if err != nil {
	    return err
	}

	ln, err := agent.Forward(ctx,
		ngrok.WithUpstream(address),
		ngrok.WithURL("your-reserved-domain"),
		ngrok.WithTrafficPolicy(trafficPolicy),
	)

	if err != nil {
		fmt.Println("Error", err)
		return err
	}

	fmt.Println("Endpoint online: forwarding from", ln.URL(), "to", address)

	// Explicitly stop forwarding; otherwise it runs indefinitely
	<-ln.Done()
	return nil
}
