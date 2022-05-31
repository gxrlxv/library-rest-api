package app

import (
	"context"
	"fmt"
	"github.com/gxrlxv/library-rest-api/internal/adapters/api/author"
	"github.com/gxrlxv/library-rest-api/internal/adapters/api/user"
	"github.com/gxrlxv/library-rest-api/internal/config"
	"github.com/gxrlxv/library-rest-api/internal/repository"
	"github.com/gxrlxv/library-rest-api/internal/service"
	"github.com/gxrlxv/library-rest-api/pkg/client/mongodb"
	"github.com/gxrlxv/library-rest-api/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func Run() {
	var listener net.Listener
	var listenErr error

	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()
	cfg := config.GetConfig()

	// MONGO DB
	cfgMongo := cfg.MongoDB
	mongoDBClient, err := mongodb.NewClient(context.Background(), cfgMongo.Host, cfgMongo.Port, cfgMongo.Username,
		cfgMongo.Password, cfgMongo.Database, cfgMongo.AuthDB)
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(mongoDBClient, logger)
	userService := service.NewUserService(userRepo, logger)

	authorRepo := repository.NewAuthorRepository(mongoDBClient, logger)
	authorService := service.NewAuthorService(authorRepo, logger)

	userHandler := user.NewUserHandler(userService, logger)
	userHandler.Register(router)

	authorHandler := author.NewAuthorHandler(authorService, logger)
	authorHandler.Register(router)

	if cfg.Listen.Type == "sock" {
		logger.Info("detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")

		logger.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("server is listening unix socket %s", socketPath)

	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if listenErr != nil {
		log.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
