FROM golang:1.27.3-trixie AS build

ENV CGO_ENABLED=0
WORKDIR /src
COPY go.* ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build \
    go build -trimpath -ldflags=-s -o /out/app ./cmd/app

FROM gcr.io/distroless/static-debian13:nonroot

COPY --from=build /out/app /app
ENTRYPOINT ["/app"]
