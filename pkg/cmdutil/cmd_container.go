package cmdutil

import (
	"github.com/bharat-rajani/GoHunter/internal/config"
	"github.com/bharat-rajani/GoHunter/pkg/iostreams"
	"net/http"
)

type CMDContainerFactory struct {
	IOStreams  *iostreams.IOStreams
	Config     func() (*config.Configuration, error)
	HttpClient func() (*http.Client, error)
}
