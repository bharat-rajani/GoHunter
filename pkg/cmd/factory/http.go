package factory

import (
	"github.com/bharat-rajani/GoHunter/internal/client"
	"github.com/bharat-rajani/GoHunter/internal/config"
	"github.com/bharat-rajani/GoHunter/pkg/iostreams"
	"net/http"
)

func NewHttpClient(io *iostreams.IOStreams, configuration *config.Configuration) *http.Client {

	var opts []client.ClientOption
	//if verbose := os.Getenv("DEBUG"); verbose != "" {
	//}
	//logTraffic := true
	//opts = append(opts, client.VerboseLog(io.ErrOut, true, io.IsStderrTTY()))

	//create opts based on configuration options
	opts = append(opts,
		client.AddHeader("User-Agent",
				"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:55.0) Gecko/20100101 Firefox/55.0"),
	)

	return client.NewHTTPClient(opts...)
}
