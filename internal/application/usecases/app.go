package app

import (
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/start"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/status"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/stop"
)

type App struct {
	StartFlowSessionUseCase  start.UseCase
	StopFlowSessionUseCase   stop.UseCase
	FlowSessionStatusUseCase status.UseCase
}

func NewApp(startFlowSessionUseCase start.UseCase, stopFlowSessionUseCase stop.UseCase, flowSessionStatusUseCase status.UseCase) *App {
	return &App{
		StartFlowSessionUseCase:  startFlowSessionUseCase,
		StopFlowSessionUseCase:   stopFlowSessionUseCase,
		FlowSessionStatusUseCase: flowSessionStatusUseCase,
	}
}
