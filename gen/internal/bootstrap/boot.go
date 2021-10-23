package bootstrap

import (
	"github.com/wishperera/GVAT/gen/internal/application"
	"github.com/wishperera/GVAT/gen/internal/container"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Boot() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("starting application...\n")

	con := container.NewContainment()
	bindAndInit(con)
	con.Start(
		application.ModuleRouter,
	)

	<-quit
	log.Printf("shutting down application...\n")
	con.ShutDown(
		application.ModuleRouter,
	)

	log.Printf("application closed...\n")
}
