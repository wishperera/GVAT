package bootstrap

import (
	"github.com/wishperera/GVAT/internal/application"
	"github.com/wishperera/GVAT/internal/container"
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
		application.ModuleEUVIESAdaptor,
		application.ModuleRouter,
	)

	<-quit
	log.Printf("shutting down application...\n")
	con.ShutDown(
		application.ModuleRouter,
		application.ModuleEUVIESAdaptor,
	)

	log.Printf("application closed...\n")
}
