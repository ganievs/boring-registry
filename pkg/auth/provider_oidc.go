package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"github.com/MicahParks/keyfunc/v3"
)

type OIDCProvider struct {
	jwksURL string
	claims  map[string]string
}

func (p *OIDCProvider) String() string { return "oidc" }

func (p *OIDCProvider) Verify(ctx context.Context, token string) error {
	// Create the keyfunc.Keyfunc.
	jwks, err := keyfunc.NewDefault([]string{p.jwksURL})
	if err != nil {
		return fmt.Errorf("Failed to create JWK Set from resource at the given URL.\nError: %s", err)
	}

	// Parse the JWT.
	tokenParsed, err := jwt.Parse(token, jwks.Keyfunc)
	if err != nil {
		return fmt.Errorf("Failed to parse the JWT.\nError: %s", err)
	}

	// Check if the token is valid.
	if !tokenParsed.Valid {
		return errors.New("The token is not valid.")
	}

	return nil
}

func NewOIDCProvider(jwksURL string, claims ...string) Provider {
	m := make(map[string]string)

	for _, claim := range claims {
		parts := strings.Split(claim, "=")
		if len(parts) != 2 {
			continue
		}

		key, val := parts[0], parts[1]

		m[key] = val
	}

	return &OIDCProvider{
		jwksURL: jwksURL,
		claims:  m,
	}
}
