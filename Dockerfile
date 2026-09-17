# ---------------------------------------------------------
# Node
# Keep aligned with upstream Anubis CI
# ---------------------------------------------------------
FROM node:24.15.0-bookworm AS node


# ---------------------------------------------------------
# Build
# ---------------------------------------------------------
FROM golang:1.26.3-bookworm AS builder

# Copy Node/npm from the official Node image.
COPY --from=node /usr/local/ /usr/local/

# Dependencies required by the Anubis asset build.
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        build-essential \
        brotli \
        ca-certificates \
        curl \
        git \
        zstd \
    && rm -rf /var/lib/apt/lists/*

# Install Rust for the WebAssembly challenge code.
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs \
    | sh -s -- -y --profile minimal \
    && /root/.cargo/bin/rustup target add wasm32-unknown-unknown

ENV PATH="/root/.cargo/bin:${PATH}"
ENV HUSKY=0

WORKDIR /src

# Copy our Anubis fork.
COPY . .

# Install JS dependencies.
RUN npm ci

# Generate all Anubis assets:
# - generated Go assets
# - WASM
# - JS fallback
# - web assets
# - XESS
RUN npm run assets

# Build a completely static Anubis binary.
RUN mkdir -p /out \
    && CGO_ENABLED=0 GOOS=linux go build \
        -trimpath \
        -ldflags='-s -w -extldflags "-static"' \
        -o /out/anubis \
        ./cmd/anubis


# ---------------------------------------------------------
# Runtime
# ---------------------------------------------------------
FROM cgr.dev/chainguard/static:latest

COPY --from=builder --chown=1000:1000 /out/anubis /anubis

USER 1000:1000

EXPOSE 8923

ENTRYPOINT ["/anubis"]