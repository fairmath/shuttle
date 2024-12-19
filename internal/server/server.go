package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/rpc"
	"go.uber.org/zap"

	"github.com/fairmath/shuttle/cmd/config"
	"github.com/fairmath/shuttle/internal/server/handlers"
)

type Handler interface {
	Name() string
}

type ProxyTo handlers.EthProxy

type Server struct {
	srv         *rpc.Server
	ws          *WebsocketPool
	addr        string
	proxyTo     ProxyTo
	log         *zap.Logger
	handlersCfg *handlers.Config
}

func NewServer(cfg config.Config, proxyTo ProxyTo, log *zap.Logger) (*Server, error) {
	server := &Server{
		addr:        cfg.ListenAddr,
		proxyTo:     proxyTo,
		log:         log,
		handlersCfg: cfg.HandlersConfig,
	}

	srv := rpc.NewServer()

	for _, h := range server.registeredHandlers(proxyTo) {
		if err := srv.RegisterName(h.Name(), h); err != nil {
			return nil, fmt.Errorf("register '%s' handler: %w", h.Name(), err)
		}
	}

	ws, err := NewWebsocketPool(cfg.WSTendermintURL, log)
	if err != nil {
		return nil, fmt.Errorf("websocket: %w", err)
	}

	server.srv = srv
	server.ws = ws

	return server, nil
}

func (s *Server) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.srv.ServeHTTP)
	mux.HandleFunc("/websocket", s.ws.ServeHTTP)

	//nolint:gosec // use http.Server with settings later
	if err := http.ListenAndServe(s.addr, mux); err != nil {
		s.log.Error("start server", zap.Error(err))
	}

	return nil
}

func (s *Server) Close(ctx context.Context) error {
	done := make(chan struct{})
	go func() {
		defer close(done)

		s.srv.Stop()
	}()

	select {
	case <-done:
	case <-ctx.Done():
		return fmt.Errorf("stopping server: %w", ctx.Err())
	}

	return nil
}

func (s *Server) registeredHandlers(proxyTo ProxyTo) []Handler {
	return []Handler{
		handlers.NewEthServer(proxyTo, *s.handlersCfg),
		handlers.NewTxPool(),
		handlers.NewDebug(proxyTo),
	}
}
