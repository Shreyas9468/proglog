package auth

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/casbin/casbin/v3"
)

type Authorizer struct {
	enforcer *casbin.Enforcer
}

func New(model, policy string) *Authorizer {
	enforcer, err := casbin.NewEnforcer(model, policy)
	if err != nil {
		panic(err)
	}
	return &Authorizer{enforcer: enforcer}
}

func (a *Authorizer) Authorize(sub, obj, act string) error {
	ok, err := a.enforcer.Enforce(sub, obj, act)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	if !ok {
		msg := fmt.Sprintf(
			"%s not permitted to %s to %s",
			sub,
			act,
			obj,
		)
		st := status.New(codes.PermissionDenied, msg)
		return st.Err()
	}
	return nil
}
