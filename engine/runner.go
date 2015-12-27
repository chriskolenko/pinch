package engine

import "github.com/webcanvas/pinch/shared/models"

// Runner is the interface that every fact, service and pincher has to implement
type Runner interface {
	Setup(models.Raw) (models.Result, error)
	Run(models.Raw) (models.Result, error)
}
