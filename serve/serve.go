package serve

import (
	"fmt"
	"net/http"
	"path/filepath"

	quic "github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/h2quic"
	log "github.com/sirupsen/logrus"
)

// Serve starts the server
func Serve(cfg Config) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	versions := []quic.VersionNumber{
		quic.VersionGQUIC43,
		quic.VersionGQUIC42,
		quic.VersionGQUIC39,
	}

	server := h2quic.Server{
		Server: &http.Server{
			Addr:    addr,
			Handler: ServerHandler(),
		},
		QuicConfig: &quic.Config{Versions: versions},
	}
	log.Info("Starting QUIC server at: ", addr)
	log.Fatal(server.ListenAndServeTLS(filepath.Join(cfg.CertsPath, cfg.CertName), filepath.Join(cfg.CertsPath, cfg.KeyName)))
}

// ServerHandler maps all the paths to handlers via mux
func ServerHandler() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/tile", handlerTile)
	mux.HandleFunc("/tiles", handlerTiles)
	mux.HandleFunc("/hello-world", handlerHelloWorld)
	mux.HandleFunc("/", handlerRoot)

	return mux
}
