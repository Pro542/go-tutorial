### Running a module

1. Run with `cd hello` and `go run`

2. Check (pkg.go.dev)[https://pkg.go.dev] to search for packages

3. After adding in code, use `go mod tidy` to add as requirement and auth the module

### Creating a module

1. When creating a new module, `mkdir`, `cd` to it, then do `go mod init`

(`go mod init <module_path_url>` example `go mod init github.com/pro542/my-go-module` if you want others to use it
otherwise `go mod init example/hello`)

### Installing go

1. Updated go by running
`sudo rm -rf /usr/local/go`
Download the tar file, then from ~/Downloads/
`tar -C /usr/local -xzf go1.26.4.linux-amd64.tar.gz` (had 1.22.5 before)
add `export PATH=$PATH:/usr/local/go/bin` (it was already there)
Verify with `go version`

=======
### Suggestions

A progression of projects to learn Go and backend concepts, ordered so each teaches a distinct BE concept. The highest signal comes from doing 2 → 3 → 4 thoroughly.

1. **CLI tool** (e.g. a git-log analyzer or Markdown link-checker)
   - Learn syntax, structs, interfaces, `if err != nil` error handling, and the standard library without networking noise.
   - *BE concepts:* file I/O, goroutines for parallel work, idiomatic error handling.

2. **JSON REST API with no framework** (e.g. a URL shortener or todo API)
   - Use only `net/http` + `encoding/json`. Go 1.22+ has good built-in routing via `http.ServeMux`.
   - *BE concepts:* HTTP servers, routing, request/response lifecycle, middleware, status codes, REST design.

3. **Same API + a real database** (Postgres via `database/sql`, `pgx`, or `sqlc`)
   - The single most CV-relevant jump. Use `goose` or `golang-migrate` for migrations.
   - *BE concepts:* persistence, schema design, migrations, transactions, connection pooling, SQL injection, N+1 queries.

4. **A concurrency-heavy project** (web scraper, rate limiter, or job queue / worker pool)
   - Goroutines + channels are Go's flagship feature and the thing FE engineers rarely touch.
   - *BE concepts:* concurrency, channels, `context` for cancellation/timeouts, mutexes, race detection (`go test -race`).

5. **A small real-time service** (chat server over WebSockets, or an SSE live-feed)
   - Plays to your FE intuition for the client side while teaching the server side of real-time.
   - *BE concepts:* WebSockets, pub/sub, in-memory state, handling many connections.

6. **Capstone — containerize and deploy one of the above**
   - Multi-stage `Dockerfile` (~10MB image), env-based config, a `/healthz` endpoint, structured logging (`slog`), graceful shutdown.
   - *BE concepts:* the deployment/ops story interviewers ask about — Docker, config, observability.

**Tips for a FE engineer:**
- Resist reaching for a framework (Gin, Echo, Fiber) at first — the stdlib teaches the concepts. Add one later if you want.
- Write tests from project 2 onward; table-driven tests are a Go cultural expectation.
- Run `go vet` and `golangci-lint` to show idiomatic Go.
- Resources: the [Go tour](https://go.dev/tour) + *"Let's Go"* by Alex Edwards (builds project 2/3 properly).

