package engine

import "github.com/webcanvas/pinch/shared/models"

// Runner is the interface that every fact, service and pincher has to implement
type Runner interface {
	Run() (models.Result, error)
}
