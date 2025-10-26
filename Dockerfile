# syntax=docker/dockerfile:1

################################################################################
# Build stage
ARG GO_VERSION=1.25.3
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
WORKDIR /src

# Download dependencies
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

ARG TARGETARCH

# Build the Go binary
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/server .

################################################################################
# Runtime stage (RHEL-compatible, minimal)
FROM rockylinux:9-minimal AS final

# Install runtime dependencies and clean up
RUN microdnf install -y \
        ca-certificates \
        tzdata \
    && update-ca-trust \
    && microdnf clean all \
    && rm -rf /var/cache/dnf /var/cache/yum

# Create a non-privileged user
ARG UID=10001
RUN useradd -u ${UID} -r -s /sbin/nologin appuser
USER appuser

# Copy the Go binary from the build stage
COPY --from=build /bin/server /bin/server

# Expose the TCP port
EXPOSE 38759

# Run the server
ENTRYPOINT ["/bin/server"]
