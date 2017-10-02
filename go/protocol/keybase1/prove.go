// Auto-generated by avdl-compiler v1.3.20 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/prove.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type CheckProofStatus struct {
	Found     bool        `codec:"found" json:"found"`
	Status    ProofStatus `codec:"status" json:"status"`
	ProofText string      `codec:"proofText" json:"proofText"`
	State     ProofState  `codec:"state" json:"state"`
}

func (o CheckProofStatus) DeepCopy() CheckProofStatus {
	return CheckProofStatus{
		Found:     o.Found,
		Status:    o.Status.DeepCopy(),
		ProofText: o.ProofText,
		State:     o.State.DeepCopy(),
	}
}

type StartProofResult struct {
	SigID SigID `codec:"sigID" json:"sigID"`
}

func (o StartProofResult) DeepCopy() StartProofResult {
	return StartProofResult{
		SigID: o.SigID.DeepCopy(),
	}
}

type StartProofArg struct {
	SessionID    int    `codec:"sessionID" json:"sessionID"`
	Service      string `codec:"service" json:"service"`
	Username     string `codec:"username" json:"username"`
	Force        bool   `codec:"force" json:"force"`
	PromptPosted bool   `codec:"promptPosted" json:"promptPosted"`
	Auto         bool   `codec:"auto" json:"auto"`
}

func (o StartProofArg) DeepCopy() StartProofArg {
	return StartProofArg{
		SessionID:    o.SessionID,
		Service:      o.Service,
		Username:     o.Username,
		Force:        o.Force,
		PromptPosted: o.PromptPosted,
		Auto:         o.Auto,
	}
}

type CheckProofArg struct {
	SessionID int   `codec:"sessionID" json:"sessionID"`
	SigID     SigID `codec:"sigID" json:"sigID"`
}

func (o CheckProofArg) DeepCopy() CheckProofArg {
	return CheckProofArg{
		SessionID: o.SessionID,
		SigID:     o.SigID.DeepCopy(),
	}
}

type ProveInterface interface {
	StartProof(context.Context, StartProofArg) (StartProofResult, error)
	CheckProof(context.Context, CheckProofArg) (CheckProofStatus, error)
}

func ProveProtocol(i ProveInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.prove",
		Methods: map[string]rpc.ServeHandlerDescription{
			"startProof": {
				MakeArg: func() interface{} {
					ret := make([]StartProofArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]StartProofArg)
					if !ok {
						err = rpc.NewTypeError((*[]StartProofArg)(nil), args)
						return
					}
					ret, err = i.StartProof(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"checkProof": {
				MakeArg: func() interface{} {
					ret := make([]CheckProofArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]CheckProofArg)
					if !ok {
						err = rpc.NewTypeError((*[]CheckProofArg)(nil), args)
						return
					}
					ret, err = i.CheckProof(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type ProveClient struct {
	Cli rpc.GenericClient
}

func (c ProveClient) StartProof(ctx context.Context, __arg StartProofArg) (res StartProofResult, err error) {
	err = c.Cli.Call(ctx, "keybase.1.prove.startProof", []interface{}{__arg}, &res)
	return
}

func (c ProveClient) CheckProof(ctx context.Context, __arg CheckProofArg) (res CheckProofStatus, err error) {
	err = c.Cli.Call(ctx, "keybase.1.prove.checkProof", []interface{}{__arg}, &res)
	return
}
