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

	createctl "go-nggong/internal/dogprice/controller/create"
	getallctl "go-nggong/internal/dogprice/controller/getall"
	getpricectl "go-nggong/internal/dogprice/controller/getprice"
	updatepricectl "go-nggong/internal/dogprice/controller/updateprice"
	createrepo "go-nggong/internal/dogprice/repository/create"
	getallrepo "go-nggong/internal/dogprice/repository/getall"
	getpricerepo "go-nggong/internal/dogprice/repository/getprice"
	updatepricerepo "go-nggong/internal/dogprice/repository/updateprice"
	createuc "go-nggong/internal/dogprice/usecase/create"
	getalluc "go-nggong/internal/dogprice/usecase/getall"
	getpriceuc "go-nggong/internal/dogprice/usecase/getprice"
	updatepriceuc "go-nggong/internal/dogprice/usecase/updateprice"
	"go-nggong/internal/platform/config"
	"go-nggong/internal/platform/db"
	"go-nggong/internal/platform/logging"
)

func main() {
	cfgPath := flag.String("config", "configs/dogprice-service.json", "path to config")
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

	getAllUC := getalluc.NewUseCase(getallrepo.New(conn))
	getPriceUC := getpriceuc.NewUseCase(getpricerepo.New(conn))
	createUC := createuc.NewUseCase(createrepo.New(conn))
	updatePriceUC := updatepriceuc.NewUseCase(updatepricerepo.New(conn))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /prices", getallctl.New(getAllUC, logger).Handle)
	mux.HandleFunc("GET /prices/{dogId}", getpricectl.New(getPriceUC, logger).Handle)
	mux.HandleFunc("POST /prices", createctl.New(createUC, logger).Handle)
	mux.HandleFunc("PUT /prices/{dogId}", updatepricectl.New(updatePriceUC, logger).Handle)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(`{"status":"UP"}`))
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: mux,
	}

	go func() {
		logger.Info("dogprice-service listening", "port", cfg.Server.Port)
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
