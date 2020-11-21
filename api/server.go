package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type Server struct {
	*http.Server
}

func NewServer() (*Server,error)  {
/*	cfg, err := config.New()

	if err != nil {
		return nil, err
	}

	api, err := New(true, cfg)
	if err != nil {
		return nil, err
	}
*/
	srv := http.Server{
		Addr:    ":8080",
		//Handler: api,
	}

	return &Server{&srv}, nil
}

//Start runs ListenAndServe on the http.Server with graceful shutdown.
func (srv *Server) Start() {

	log.Println("Starting Nyumbapoa API Server...")

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	log.Printf("Listening on %s\n", srv.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit

	log.Println("Shutting down fupisha API server... Reason:", sig)

	//teardown logic here

	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Println("Fupisha API server gracefully stopped.")
}


