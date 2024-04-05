# go-custom-path-server

This is a more k8s-friendly wrapper around [rsc.io/go-import-redirector](https://pkg.go.dev/rsc.io/go-import-redirector), 
that can take arguments as command line args or environment variables.

## Usage
```
Usage: go-custom-path-server --import-path=STRING --repository=STRING [flags]

Flags:
  -h, --help                      Show context-sensitive help.
  -a, --listen-address=":http"    The address for the server to listen on ($GO_CPS_LISTEN_ADDRESS)
  -v, --vcs="git"                 The VCS type (git, hg, etc) ($GO_CPS_VCS_TYPE)
  -i, --import-path=STRING        The custom import path (trailing /* wildcards allowed) ($GO_CPS_IMPORT_PATH)
  -r, --repository=STRING         The VCS repository associated with the import path (trailing /* wildcards allowed) ($GO_CPS_REPOSITORY)
```