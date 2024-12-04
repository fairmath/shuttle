package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/rpc"
	"go.uber.org/zap"

	"github.com/fairmath/shuttle/internal/server/handlers"
)

type Handler interface {
	Name() string
}

type ProxyTo handlers.EthProxy

type Server struct {
	srv     *rpc.Server
	ws      *WebsocketPool
	addr    string
	proxyTo ProxyTo
	log     *zap.Logger
}

func NewServer(listenAddr string, proxyTo ProxyTo, log *zap.Logger) (*Server, error) {
	srv := rpc.NewServer()

	for _, h := range registeredHandlers(proxyTo) {
		if err := srv.RegisterName(h.Name(), h); err != nil {
			return nil, fmt.Errorf("register '%s' handler: %w", h.Name(), err)
		}
	}

	return &Server{
		srv:     srv,
		ws:      NewWebsocketPool(log),
		addr:    listenAddr,
		proxyTo: proxyTo,
		log:     log,
	}, nil
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

func registeredHandlers(proxyTo ProxyTo) []Handler {
	return []Handler{
		handlers.NewEthServer(proxyTo),
		handlers.NewTxPool(),
	}
}
