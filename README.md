# What's this
This is an integration script for Nextcloud and Keycloak. This app integrates itself into a Nextcloud Docker container and configures OpenID connect authentication for Nextcloud via Keycloak accordint to `app` container environment variables.

# How this works

This script installs Nextcloud in a Docker container, installs Social Login in it and lauches Keycloak integration script.
# Environment variables:

See `./run.sh` (except for NEXTCLOUD_ADMIN_USER, NEXTCLOUD_ADMIN_PASSWORD vars that are set in `.env`)

This image is based on

| Variable                   | Meaning                                                                                                                                                                    |
| -------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `NEXTCLOUD_ADMIN_USER`     | Your Nextcloud admin  username (set in `.env`)                                                                                                                             |
| `NEXTCLOUD_ADMIN_PASSWORD` | Your Nextcloud admin password  (set in `.env`)                                                                                                                             |
| `NEXTCLOUD_URL`            | Your Nextcloud URL. As the script runs inside a container, leave it  http://localhost:8080 if you use HTTP. Change the scheme to HTTPS if Nextcloud is configured that way |
| `KEYCLOAK_URL`             | Keycloak URL to connect to.                                                                                                                                                |
| `KEYCLOAK_REALM`           | Keycloak realm to create a new client on                                                                                                                                   |
| `KEYCLOAK_LOGIN_REALM`     | Keycloak realm to login administrator on                                                                                                                                   |
| `KEYCLOAK_USER`            | Keycloak administrator username                                                                                                                                            |
| `KEYCLOAK_PASSWORD`        | Keycloak administrator password                                                                                                                                            |
| `CLIENT_ID`                | Keycloak NEW client ID for Nextcloud                                                                                                                                       |
| `CLIENT_NAME`              | Keycloak NEW client name for Nextcloud                                                                                                                                     |
| `CLIENT_SECRET`            | Keycloak NEW client secret for Nextcloud                                                                                                                                   |
| `NEXTCLOUD_LOGIN_NAME`     | How a Nextcloud social login method should be named                                                                                                                        |


To build & run Nextcloud in docker, run:
```
make docker
```

# Alternative uses

You can use this utility w/o docker for integrating external Nextcloud with Keycloak. To do so, run
`go build`

and then set environment variables from the table above in `./run.sh` and run `./nextcloud_kcintegrate`.
