# =============================================================================
#  Test Container for Vaious Go Versions
# =============================================================================
# Default version (minimum supported version. See go.mod)
ARG VARIANT="1.24-alpine"

# -----------------------------------------------------------------------------
#  Main Stage
# -----------------------------------------------------------------------------
FROM golang:${VARIANT}

RUN apk add --no-cache \
    git \
    alpine-sdk \
    build-base

WORKDIR /workspaces

ENTRYPOINT go mod download && go test -race ./...
