package grpc

import (
	"github.com/micro/go-micro/v2/errors"
	"google.golang.org/grpc/status"
)

func microError(err error) error {
	// no error
	switch err {
	case nil:
		return nil
	}

	if verr, ok := err.(*errors.Error); ok {
		return verr
	}

	// grpc error
	if s, ok := status.FromError(err); ok {
		if e := errors.Parse(s.Message()); e.Code > 0 {
			return e // actually a micro error
		}
		return err
		// return errors.InternalServerError("go.micro.client", s.Message())
	}

	// do nothing
	return err
}
