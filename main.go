package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Renos-id/go-starter-template/infrastructure"
	"github.com/Renos-id/go-starter-template/lib/response"
	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
)

func init() {
	infrastructure.InitLoadEnv()
	httpin.UseGochiURLParam("path", chi.URLParam)

}

func main() {
	//init DB
	// var dbConn *sqlx.DB
	// if os.Getenv("DB_HOST") != "" {
	// 	dbConn = database.Open()
	// 	dbMaxIdleConnection, _ := strconv.ParseInt(os.Getenv("DB_MAX_IDLE_CONNECTION"), 10, 0)
	// 	dbMaxOpenConns, _ := strconv.ParseInt(os.Getenv("DB_MAX_OPEN_CONNS"), 10, 0)
	// 	dbConnMaxLifetime, _ := strconv.ParseInt(os.Getenv("DB_CONN_MAX_LIFETIME"), 10, 0)
	// 	dbConnMaxLifetimeDuration := time.Duration(dbConnMaxLifetime * int64(time.Minute))

	// 	dbConn.SetMaxIdleConns(int(dbMaxIdleConnection))
	// 	dbConn.SetMaxOpenConns(int(dbMaxOpenConns))
	// 	dbConn.SetConnMaxLifetime(dbConnMaxLifetimeDuration)

	// if err := dbConn.Ping(); err != nil {
	// 	log.Fatal("Can't connect to Database!", err)
	// }
	// }
	//End Init DB
	r := infrastructure.InitChiRouter()
	logger := infrastructure.InitLog()

	response.SetLogging(logger)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:              fmt.Sprintf("0.0.0.0:%s", os.Getenv("APP_PORT")),
		Handler:           r,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	log.Printf("%s running on PORT : %s \n", os.Getenv("APP_NAME"), os.Getenv("APP_PORT"))
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed Running %s on PORT : %s \n", os.Getenv("APP_NAME"), os.Getenv("APP_PORT"))
	}
}

// func graceful_shutdown() {
// 	// The HTTP Server
// 	server := &http.Server{Addr: fmt.Sprintf("0.0.0.0:%s", os.Getenv("APP_PORT")), Handler: service()}

// 	// Server run context
// 	serverCtx, serverStopCtx := context.WithCancel(context.Background())

// 	// Listen for syscall signals for process to interrupt/quit
// 	sig := make(chan os.Signal, 1)
// 	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
// 	go func() {
// 		<-sig

// 		// Shutdown signal with grace period of 30 seconds
// 		shutdownCtx, cancelFunc := context.WithTimeout(serverCtx, 30*time.Second)
// 		defer cancelFunc()
// 		go func() {
// 			<-shutdownCtx.Done()
// 			if shutdownCtx.Err() == context.DeadlineExceeded {
// 				log.Fatal("graceful shutdown timed out.. forcing exit.")
// 			}
// 		}()

// 		// Trigger graceful shutdown
// 		fmt.Println("Trigger graceful shutdown")
// 		err := server.Shutdown(shutdownCtx)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		serverStopCtx()
// 	}()

// 	// Run the server
// 	err := server.ListenAndServe()
// 	if err != nil && err != http.ErrServerClosed {
// 		log.Fatal(err)
// 	}

// 	// Wait for server context to be stopped
// 	<-serverCtx.Done()
// }
