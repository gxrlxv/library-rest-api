package app

import (
	"github.com/gxrlxv/library-rest-api/internal/adapters/api/user"
	"github.com/gxrlxv/library-rest-api/internal/service"
	"github.com/gxrlxv/library-rest-api/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func Run() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	userService := service.NewUserService(nil)

	handler := user.NewUserHandler(userService)
	handler.Register(router)

	listener, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatalln(server.Serve(listener))
}
