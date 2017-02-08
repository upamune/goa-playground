package main

import (
	"strconv"

	"github.com/goadesign/goa"
	"github.com/upamune/goa-playground/app"
)

// OperandsController implements the operands resource.
type OperandsController struct {
	*goa.Controller
}

// NewOperandsController creates a operands controller.
func NewOperandsController(service *goa.Service) *OperandsController {
	return &OperandsController{Controller: service.NewController("OperandsController")}
}

// Add runs the add action.
func (c *OperandsController) Add(ctx *app.AddOperandsContext) error {
	// OperandsController_Add: start_implement

	// Put your logic here
	sum := ctx.Left + ctx.Right

	// operandscontroller_add: end_implement
	return ctx.OK([]byte(strconv.Itoa(sum)))
}

// Sub runs the sub action.
func (c *OperandsController) Sub(ctx *app.SubOperandsContext) error {
	// OperandsController_Sub: start_implement

	// Put your logic here
	sub := ctx.Left - ctx.Right

	// operandscontroller_add: end_implement
	return ctx.OK([]byte(strconv.Itoa(sub)))
}
