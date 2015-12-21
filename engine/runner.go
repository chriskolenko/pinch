package engine

import "github.com/webcanvas/pinch/shared/models"

type Runner interface {
	Run(map[string]string) (models.Result, error)
}
