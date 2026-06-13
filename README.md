### Steps taken

1. Updated go by running
`sudo rm -rf /usr/local/go`
Download the tar file, then from ~/Downloads/
`tar -C /usr/local -xzf go1.26.4.linux-amd64.tar.gz` (had 1.22.5 before)
add `export PATH=$PATH:/usr/local/go/bin` (it was already there)
Verify with `go version`

2. `go mod init <module_path_url>` example `go mod init github.com/pro542/my-go-module` if you want others to use it
otherwise `go mod init example/hello`

3. Check (pkg.go.dev)[https://pkg.go.dev] to search for packages

4. After adding in code, use `go mod tidy` to add as requirement and auth the module


