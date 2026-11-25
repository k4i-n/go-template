# go-template

A starting point for a Go 1.27 project: golangci-lint and govulncheck pinned in a tools module and run with `go tool`, stdlib tests with a 100% coverage gate and the race detector, a static binary in a distroless non-root image, and GitHub Actions with SHA-pinned actions.

## Start a project

```sh
go run golang.org/x/tools/cmd/gonew@latest github.com/k4i-n/go-template example.com/myapp
cd myapp
make test
make run       # prints "Hello, world!"
```

Or use this template on GitHub, then change the module path with `go mod edit -module example.com/myapp` and in the imports.

## Commands

| | |
|---|---|
| `make format` | gofmt and goimports |
| `make lint` | `go.mod` of both modules tidy, golangci-lint, govulncheck |
| `make test` | tests with the race detector; fails if any function except `main` is below 100% |
| `make build` | Docker image `app` |
| `make run` | run `cmd/app` |

Add a dependency with `go get <module>`, a development tool with `cd tools && go get -tool <package>`.

## What is set up

- **Go 1.27.3** from `go.mod`; an older local Go downloads it automatically.
- **Tools** live in their own module under `tools/`, so their dependencies never enter the application's `go.mod`, and local runs and CI use the same versions.
- **`main`** only calls `run(args, stdout)`, which is tested like any other function; it is the one function the coverage gate skips.
- **Supply chain**: Dependabot updates both modules, GitHub Actions and the Docker base images weekly, waiting 7 days after each release.
- **Docker**: `CGO_ENABLED=0 -trimpath` build, copied into `gcr.io/distroless/static-debian13:nonroot`; the image is under 10 MB and runs as an unprivileged user.
- **CI** on every pull request and push to `main`: lint, vulnerability scan, tests, Docker build and run; zizmor audits the workflows. Workflows have no permissions by default and actions are pinned by commit SHA.

## License

[MIT](LICENSE)
