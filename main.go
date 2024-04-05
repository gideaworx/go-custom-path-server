package main

import (
	"net/http"

	"github.com/alecthomas/kong"
	"rsc.io/go-import-redirector/godoc"
)

type programArgs struct {
	ListenAddress string `short:"a" env:"GO_CPS_LISTEN_ADDRESS" default:":http" help:"The address for the server to listen on"`
	VCS           string `short:"v" env:"GO_CPS_VCS_TYPE" default:"git" help:"The VCS type (git, hg, etc)"`
	ImportPath    string `short:"i" env:"GO_CPS_IMPORT_PATH" required:"true" help:"The custom import path (trailing /* wildcards allowed)"`
	Repository    string `short:"r" env:"GO_CPS_REPOSITORY" required:"true" help:"The VCS repository associated with the import path (trailing /* wildcards allowed)"`
}

func main() {
	var args programArgs
	kong.Parse(&args)

	http.ListenAndServe(args.ListenAddress, godoc.Redirect(args.VCS, args.ImportPath, args.Repository))
}
