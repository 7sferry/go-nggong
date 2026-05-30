package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"go-nggong/internal/dog/client"
	createctl "go-nggong/internal/dog/controller/create"
	createcontactctl "go-nggong/internal/dog/controller/createcontact"
	getallctl "go-nggong/internal/dog/controller/getall"
	getbyidctl "go-nggong/internal/dog/controller/getbyid"
	getcontactctl "go-nggong/internal/dog/controller/getcontact"
	updatectl "go-nggong/internal/dog/controller/update"
	updatecontactctl "go-nggong/internal/dog/controller/updatecontact"
	createrepo "go-nggong/internal/dog/repository/create"
	createcontactrepo "go-nggong/internal/dog/repository/createcontact"
	getallrepo "go-nggong/internal/dog/repository/getall"
	getbyidrepo "go-nggong/internal/dog/repository/getbyid"
	getcontactrepo "go-nggong/internal/dog/repository/getcontact"
	updaterepo "go-nggong/internal/dog/repository/update"
	updatecontactrepo "go-nggong/internal/dog/repository/updatecontact"
	createuc "go-nggong/internal/dog/usecase/create"
	createcontactuc "go-nggong/internal/dog/usecase/createcontact"
	getalluc "go-nggong/internal/dog/usecase/getall"
	getbyiduc "go-nggong/internal/dog/usecase/getbyid"
	getcontactuc "go-nggong/internal/dog/usecase/getcontact"
	updateuc "go-nggong/internal/dog/usecase/update"
	updatecontactuc "go-nggong/internal/dog/usecase/updatecontact"
	"go-nggong/internal/platform/config"
	"go-nggong/internal/platform/db"
	"go-nggong/internal/platform/logging"
)

func main() {
	cfgPath := flag.String("config", "configs/dog-service.json", "path to config")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "load config:", err)
		os.Exit(1)
	}

	logger, err := logging.Setup(logging.Options{
		ConsoleLevel: parseLevel(cfg.Logging.ConsoleLevel, slog.LevelInfo),
		FileLevel:    parseLevel(cfg.Logging.FileLevel, slog.LevelWarn),
		FilePath:     cfg.Logging.File,
		AppName:      cfg.AppName,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "setup logger:", err)
		os.Exit(1)
	}

	conn, err := db.Open(cfg.Database.URL)
	if err != nil {
		logger.Error("db open failed", "err", err)
		os.Exit(1)
	}
	defer conn.Close()

	priceClient := client.New(cfg.Services["dogprice"])

	getAllUC := getalluc.NewUseCase(getallrepo.New(conn), priceClient)
	getByIDUC := getbyiduc.NewUseCase(getbyidrepo.New(conn), priceClient)
	createUC := createuc.NewUseCase(createrepo.New(conn))
	updateUC := updateuc.NewUseCase(updaterepo.New(conn))
	getContactUC := getcontactuc.NewUseCase(getcontactrepo.New(conn))
	createContactUC := createcontactuc.NewUseCase(createcontactrepo.New(conn))
	updateContactUC := updatecontactuc.NewUseCase(updatecontactrepo.New(conn))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /dogs", getallctl.New(getAllUC, logger).Handle)
	mux.HandleFunc("GET /dogs/{id}", getbyidctl.New(getByIDUC, logger).Handle)
	mux.HandleFunc("POST /dogs", createctl.New(createUC, logger).Handle)
	mux.HandleFunc("PUT /dogs/{id}", updatectl.New(updateUC, logger).Handle)
	mux.HandleFunc("GET /contacts/{dogId}", getcontactctl.New(getContactUC, logger).Handle)
	mux.HandleFunc("POST /contacts", createcontactctl.New(createContactUC, logger).Handle)
	mux.HandleFunc("PUT /contacts/{dogId}", updatecontactctl.New(updateContactUC, logger).Handle)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(`{"status":"UP"}`))
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: mux,
	}

	go func() {
		logger.Info("dog-service listening", "port", cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("http server failed", "err", err)
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	logger.Info("shutting down")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_ = server.Shutdown(shutdownCtx)
}

func parseLevel(s string, def slog.Level) slog.Level {
	switch strings.ToUpper(s) {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return def
	}
}
