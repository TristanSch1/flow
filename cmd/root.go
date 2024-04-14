package cmd

import (
	"fmt"
	"os"

	app "github.com/TristanSch1/flow/internal/application/usecases"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/start"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/stop"
	"github.com/TristanSch1/flow/internal/infra"
	"github.com/TristanSch1/flow/internal/infra/fs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "flow",
	Short: "Flow is a tool to manage your time tracking",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, world!")
	},
}

func initializeApp() *app.App {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return nil
	}

	sessionRepository := &fs.FileSystemSessionRepository{
		FlowPath: homePath,
	}
	dateProvider := &infra.RealDateProvider{}

	startFlowSessionUseCase := start.NewStartFlowSessionUseCase(sessionRepository, dateProvider)
	stopFlowSessionUseCase := stop.NewStopSessionUseCase(sessionRepository, dateProvider)

	return app.NewApp(startFlowSessionUseCase, stopFlowSessionUseCase)
}

func Execute() {
	app := initializeApp()

	rootCmd.AddCommand(startCmd(app))
	rootCmd.AddCommand(stopCmd(app))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
