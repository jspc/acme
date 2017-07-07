package main

import (
	"fmt"

	"github.com/appscode/go-dns/gandi"
)

func CreateTXTRecord(domain, token, apiKey string) error {
	p, err := gandi.NewDNSProviderCredentials(apiKey)
	if err != nil {
		return err
	}

	return p.Present(fmt.Sprintf("_acme-challenge.", domain), "", token)
}
