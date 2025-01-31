# OIDC Authentication in Boring Registry

Boring Registry supports authentication using a generic OpenID Connect (OIDC) provider. Below are example configurations for common authentication providers.

---

## Common Configuration

Boring Registry supports configuring OIDC authentication using environment variables or CLI flags. Below is a description of common options:

### Configuration options

| Variable | Flag | Description |
|----------|------|-------------|
| `BORING_REGISTRY_AUTH_OIDC_CLAIMS` | `--auth-oidc-claims` | OIDC claims to request (e.g., `openid,profile,email`). |
| `BORING_REGISTRY_AUTH_OIDC_JWKS_URL` |  `--auth-oidc-jwks-url` | URL for JSON Web Key Set (JWKS) to verify tokens. |
| `BORING_REGISTRY_LOGIN_CLIENT` | `--login-client` | OIDC client ID used for authentication. |
| `BORING_REGISTRY_LOGIN_PORTS` | `--login-ports` | Comma-separated list of allowed login redirect ports. |
| `BORING_REGISTRY_LOGIN_GRANT_TYPES` | `--login-grant-types` | OAuth2 grant types, e.g., `authz_code`. |
| `BORING_REGISTRY_LOGIN_AUTHZ` |  `--login-authz`| Authorization endpoint of the OIDC provider. |
| `BORING_REGISTRY_LOGIN_TOKEN` | `--login-token` | Token endpoint of the OIDC provider. |

???+ note
    Redirect ports (`BORING_REGISTRY_LOGIN_PORTS` or `--login-ports`) is inclusive range of TCP ports
    that Terraform may use to start a temporary HTTP server
    for handling the authorization code grant redirection.    
    (default [10000,10010])

---

## Keycloak Configuration

### Create a Keycloak Client

1. Log in to Keycloak and navigate to **Clients** → Click **Create Client**.
2. Set **Client ID** to `boring-registry` (or any preferred name).
3. Click **Next**.
4. Ensure the following settings:
   - **Authentication Flow** → Standard Flow (**Authorization Code Flow**) **Enabled**
   - **Direct Access Grants** → **Disabled**
5. Click **Next**.
6. Under **Valid Redirect URIs**, add redirect endpoints:
   ```
   http://localhost:10001/login, ..., http://localhost:10010/login
   ```
7. Click **Save**.
8. In the created client, go to **Advanced Settings** → **Proof Key for Code Exchange (PKCE)** → **Required**.
9. Click **Save**.

### Boring Registry Configuration

Set up authentication using environment variables:

```sh
export BORING_REGISTRY_AUTH_OIDC_CLAIMS="openid,profile,email"
export BORING_REGISTRY_AUTH_OIDC_JWKS_URL="https://keycloak.example.com/realms/your-realm/protocol/openid-connect/certs"
export BORING_REGISTRY_LOGIN_CLIENT="boring-registry"
export BORING_REGISTRY_LOGIN_PORTS="10000,10001"
export BORING_REGISTRY_LOGIN_GRANT_TYPES="authz_code"
export BORING_REGISTRY_LOGIN_AUTHZ="https://keycloak.example.com/realms/your-realm/protocol/openid-connect/auth"
export BORING_REGISTRY_LOGIN_TOKEN="https://keycloak.example.com/realms/your-realm/protocol/openid-connect/token"
```

---

## Auth0 Configuration

### Create an Auth0 Application

1. Log in to the Auth0 dashboard.
2. Navigate to **Applications**.
3. Click **Create Application**:
   - **Name**: `Boring Registry` (or any preferred name).
   - **Application Type**: Single Page Web Application.
4. Click on the **Settings** tab.
5. Under **Application URIs > Allowed Callback URLs**, add a comma-separated list of redirect endpoints:
   ```
   http://localhost:10001/login, ..., http://localhost:10010/login
   ```
6. Click **Save Changes**.


### Auth0 Environment Variables Configuration

Set up authentication using environment variables:

```sh
export BORING_REGISTRY_AUTH_OIDC_CLAIMS="openid,profile,email"
export BORING_REGISTRY_AUTH_OIDC_JWKS_URL="https://YOUR_DOMAIN.auth0.com/.well-known/jwks.json"
export BORING_REGISTRY_LOGIN_CLIENT="AUTH_AUTH0_CLIENT_ID"
export BORING_REGISTRY_LOGIN_GRANT_TYPES="authz_code"
export BORING_REGISTRY_LOGIN_AUTHZ="https://YOUR_DOMAIN.auth0.com/authorize"
export BORING_REGISTRY_LOGIN_TOKEN="https://YOUR_DOMAIN.auth0.com/oauth/token"
```

## Authentication

Use [`terraform login`](https://developer.hashicorp.com/terraform/cli/commands/login) to obtain token.

```sh
$ terraform login registry.example.com
```
