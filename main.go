package main

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func weakHash(secret []byte) error {
	h := md5.New()
	if _, err := h.Write(secret); err != nil {
		return fmt.Errorf("FAILED TO HASH SECRET WITH MD5: %w", err)
	}

	result := hex.EncodeToString(h.Sum(nil))

	fmt.Printf("LOGGING A SECRET HASH: %s\n", result)

	return nil
}

func handlingSecretInfo() error {
	password := "THIS_IS_A_SECRET_PASSWORD"
	fmt.Printf("LOGGING A SECRET PASSWORD: %q\n", password)

	apiKey := "THIS_IS_A_SECRET_API_KEY"
	fmt.Printf("LOGGING A SECRET API KEY: %q\n", apiKey)

	return nil
}

func weakTLSClientGet(url string) error {
	transport := http.DefaultTransport

	transportType, ok := transport.(*http.Transport)
	if !ok {
		return fmt.Errorf("TRANSPORT WAS NOT A TRANSPORT?")
	}

	transportType.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
		MinVersion:         tls.VersionSSL30,
		MaxVersion:         tls.VersionTLS10,
		CipherSuites: []uint16{
			tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
			tls.TLS_RSA_WITH_RC4_128_SHA,
			tls.TLS_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA,
		},
	}

	client := &http.Client{
		Transport: transportType,
	}

	if _, err := client.Get(url); err != nil {
		return fmt.Errorf("REQUEST ERROR: %w", err)
	}

	return nil
}

func main() {
	// Should trigger warning about golang.org/x/net/html?
	fmt.Printf("%s\n", html.UnescapeString("&nbsp;"))

	if err := weakHash([]byte("super_secret_api_key")); err != nil {
		panic(err)
	}

	if err := handlingSecretInfo(); err != nil {
		panic(err)
	}

	if err := weakTLSClientGet("https://defectdojo.com"); err != nil {
		panic(err)
	}

	print("oh no")
}
