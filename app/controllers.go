// Code generated by goagen v1.1.0, command line:
// $ goagen
// --design=github.com/upamune/goa-playground/design
// --out=$(GOPATH)/src/github.com/upamune/goa-playground
// --version=v1.1.0-dirty
//
// API "adder": Application Controllers
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// OperandsController is the controller interface for the Operands actions.
type OperandsController interface {
	goa.Muxer
	Add(*AddOperandsContext) error
	Sub(*SubOperandsContext) error
}

// MountOperandsController "mounts" a Operands resource controller on the given service.
func MountOperandsController(service *goa.Service, ctrl OperandsController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAddOperandsContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Add(rctx)
	}
	service.Mux.Handle("GET", "/add/:left/:right", ctrl.MuxHandler("Add", h, nil))
	service.LogInfo("mount", "ctrl", "Operands", "action", "Add", "route", "GET /add/:left/:right")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSubOperandsContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Sub(rctx)
	}
	service.Mux.Handle("GET", "/sub/:left/:right", ctrl.MuxHandler("Sub", h, nil))
	service.LogInfo("mount", "ctrl", "Operands", "action", "Sub", "route", "GET /sub/:left/:right")
}
