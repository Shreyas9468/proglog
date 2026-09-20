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
	if model == "" || policy == "" {
		return &Authorizer{}
	}
	enforcer, err := casbin.NewEnforcer(model, policy)
	if err != nil {
		return &Authorizer{}
	}
	return &Authorizer{enforcer: enforcer}
}

func (a *Authorizer) Authorize(sub, obj, act string) error {
	if a == nil || a.enforcer == nil {
		return nil
	}
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
