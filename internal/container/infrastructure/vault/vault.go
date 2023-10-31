package vaultclient

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

var (
	ctx = context.Background()
)

func Setup() *vault.Client {

	client, err := vault.New(
		// with environment variables:
		// - export VAULT_ADDR=http://localhost:8200
		// - export VAULT_TOKEN=my-token
		vault.WithEnvironment(),
	)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return client
}

func PopulateConfigToFile(client *vault.Client) {
	// read secret from vault
	secret, err := client.Secrets.KvV2Read(ctx,
		"hello-service",
		vault.WithToken(os.Getenv("VAULT_TOKEN")),
		vault.WithNamespace("dev"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// write secret to file
	_, err = os.Create("./internal/config/hello-service.json")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile("./internal/config/hello-service.json", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}

	// write all secret data to file

	_, err = f.WriteString(secret.Data.Data["data"].(map[string]interface{})["otel_exporter_otlp_endpoint"].(string))

	if err != nil {
		log.Fatal(err)
	}

}
