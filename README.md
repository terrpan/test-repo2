# test-repo2

A small Go HTTP service for exercising [Scrapyard](https://github.com/terrpan/scrapyard) pipelines. It has no workflows of its own: Scrapyard runs the pipelines mapped to it in `terrpan/scrapyard-pipelines`.

## Endpoints

- `GET /healthz` returns `ok`.
- `GET /hello?name=<name>` returns `{"message":"Hello, <name>!"}`. Names over 64 bytes get a 400.

Listens on `:8080`, or `$ADDR` if set.

## Local checks

These match what the `go-ci.yml` and `security-scan.yml` pipelines run:

```sh
make        # gofmt, go vet, go test -race, go build
make scan   # trivy fs, fails on HIGH or CRITICAL
```

Container image: `docker build -t test-repo2 .`
