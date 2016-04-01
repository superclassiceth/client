// Auto-generated by avdl-compiler v1.3.1 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/auth.avdl

package gregor1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type AuthToken string
type SessionID string
type AuthenticateArg struct {
	T AuthToken `codec:"t" json:"t"`
}

type AuthInterface interface {
	Authenticate(context.Context, AuthToken) error
}

func AuthProtocol(i AuthInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "gregor.1.auth",
		Methods: map[string]rpc.ServeHandlerDescription{
			"authenticate": {
				MakeArg: func() interface{} {
					ret := make([]AuthenticateArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]AuthenticateArg)
					if !ok {
						err = rpc.NewTypeError((*[]AuthenticateArg)(nil), args)
						return
					}
					err = i.Authenticate(ctx, (*typedArgs)[0].T)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type AuthClient struct {
	Cli rpc.GenericClient
}

func (c AuthClient) Authenticate(ctx context.Context, t AuthToken) (err error) {
	__arg := AuthenticateArg{T: t}
	err = c.Cli.Call(ctx, "gregor.1.auth.authenticate", []interface{}{__arg}, nil)
	return
}
