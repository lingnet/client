// Auto-generated by avdl-compiler v1.3.20 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/identify.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type IdentifyProofBreak struct {
	RemoteProof RemoteProof     `codec:"remoteProof" json:"remoteProof"`
	Lcr         LinkCheckResult `codec:"lcr" json:"lcr"`
}

func (o IdentifyProofBreak) DeepCopy() IdentifyProofBreak {
	return IdentifyProofBreak{
		RemoteProof: o.RemoteProof.DeepCopy(),
		Lcr:         o.Lcr.DeepCopy(),
	}
}

type IdentifyTrackBreaks struct {
	Keys   []IdentifyKey        `codec:"keys" json:"keys"`
	Proofs []IdentifyProofBreak `codec:"proofs" json:"proofs"`
}

func (o IdentifyTrackBreaks) DeepCopy() IdentifyTrackBreaks {
	return IdentifyTrackBreaks{
		Keys: (func(x []IdentifyKey) []IdentifyKey {
			if x == nil {
				return nil
			}
			var ret []IdentifyKey
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Keys),
		Proofs: (func(x []IdentifyProofBreak) []IdentifyProofBreak {
			if x == nil {
				return nil
			}
			var ret []IdentifyProofBreak
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Proofs),
	}
}

type Identify2Res struct {
	Upk          UserPlusKeys         `codec:"upk" json:"upk"`
	IdentifiedAt Time                 `codec:"identifiedAt" json:"identifiedAt"`
	TrackBreaks  *IdentifyTrackBreaks `codec:"trackBreaks,omitempty" json:"trackBreaks,omitempty"`
}

func (o Identify2Res) DeepCopy() Identify2Res {
	return Identify2Res{
		Upk:          o.Upk.DeepCopy(),
		IdentifiedAt: o.IdentifiedAt.DeepCopy(),
		TrackBreaks: (func(x *IdentifyTrackBreaks) *IdentifyTrackBreaks {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.TrackBreaks),
	}
}

type IdentifyLiteRes struct {
	Ul          UserOrTeamLite       `codec:"ul" json:"ul"`
	TrackBreaks *IdentifyTrackBreaks `codec:"trackBreaks,omitempty" json:"trackBreaks,omitempty"`
}

func (o IdentifyLiteRes) DeepCopy() IdentifyLiteRes {
	return IdentifyLiteRes{
		Ul: o.Ul.DeepCopy(),
		TrackBreaks: (func(x *IdentifyTrackBreaks) *IdentifyTrackBreaks {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.TrackBreaks),
	}
}

type ResolveIdentifyImplicitTeamRes struct {
	DisplayName string                              `codec:"displayName" json:"displayName"`
	TeamID      TeamID                              `codec:"teamID" json:"teamID"`
	Writers     []UserVersion                       `codec:"writers" json:"writers"`
	TrackBreaks map[UserVersion]IdentifyTrackBreaks `codec:"trackBreaks" json:"trackBreaks"`
}

func (o ResolveIdentifyImplicitTeamRes) DeepCopy() ResolveIdentifyImplicitTeamRes {
	return ResolveIdentifyImplicitTeamRes{
		DisplayName: o.DisplayName,
		TeamID:      o.TeamID.DeepCopy(),
		Writers: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			var ret []UserVersion
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Writers),
		TrackBreaks: (func(x map[UserVersion]IdentifyTrackBreaks) map[UserVersion]IdentifyTrackBreaks {
			if x == nil {
				return nil
			}
			ret := make(map[UserVersion]IdentifyTrackBreaks)
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.TrackBreaks),
	}
}

type Resolve3Arg struct {
	Assertion string `codec:"assertion" json:"assertion"`
}

func (o Resolve3Arg) DeepCopy() Resolve3Arg {
	return Resolve3Arg{
		Assertion: o.Assertion,
	}
}

type Identify2Arg struct {
	SessionID             int                 `codec:"sessionID" json:"sessionID"`
	Uid                   UID                 `codec:"uid" json:"uid"`
	UserAssertion         string              `codec:"userAssertion" json:"userAssertion"`
	Reason                IdentifyReason      `codec:"reason" json:"reason"`
	UseDelegateUI         bool                `codec:"useDelegateUI" json:"useDelegateUI"`
	AlwaysBlock           bool                `codec:"alwaysBlock" json:"alwaysBlock"`
	NoErrorOnTrackFailure bool                `codec:"noErrorOnTrackFailure" json:"noErrorOnTrackFailure"`
	ForceRemoteCheck      bool                `codec:"forceRemoteCheck" json:"forceRemoteCheck"`
	NeedProofSet          bool                `codec:"needProofSet" json:"needProofSet"`
	AllowEmptySelfID      bool                `codec:"allowEmptySelfID" json:"allowEmptySelfID"`
	NoSkipSelf            bool                `codec:"noSkipSelf" json:"noSkipSelf"`
	CanSuppressUI         bool                `codec:"canSuppressUI" json:"canSuppressUI"`
	IdentifyBehavior      TLFIdentifyBehavior `codec:"identifyBehavior" json:"identifyBehavior"`
	ForceDisplay          bool                `codec:"forceDisplay" json:"forceDisplay"`
}

func (o Identify2Arg) DeepCopy() Identify2Arg {
	return Identify2Arg{
		SessionID:             o.SessionID,
		Uid:                   o.Uid.DeepCopy(),
		UserAssertion:         o.UserAssertion,
		Reason:                o.Reason.DeepCopy(),
		UseDelegateUI:         o.UseDelegateUI,
		AlwaysBlock:           o.AlwaysBlock,
		NoErrorOnTrackFailure: o.NoErrorOnTrackFailure,
		ForceRemoteCheck:      o.ForceRemoteCheck,
		NeedProofSet:          o.NeedProofSet,
		AllowEmptySelfID:      o.AllowEmptySelfID,
		NoSkipSelf:            o.NoSkipSelf,
		CanSuppressUI:         o.CanSuppressUI,
		IdentifyBehavior:      o.IdentifyBehavior.DeepCopy(),
		ForceDisplay:          o.ForceDisplay,
	}
}

type IdentifyLiteArg struct {
	SessionID             int                 `codec:"sessionID" json:"sessionID"`
	Id                    UserOrTeamID        `codec:"id" json:"id"`
	Assertion             string              `codec:"assertion" json:"assertion"`
	Reason                IdentifyReason      `codec:"reason" json:"reason"`
	UseDelegateUI         bool                `codec:"useDelegateUI" json:"useDelegateUI"`
	AlwaysBlock           bool                `codec:"alwaysBlock" json:"alwaysBlock"`
	NoErrorOnTrackFailure bool                `codec:"noErrorOnTrackFailure" json:"noErrorOnTrackFailure"`
	ForceRemoteCheck      bool                `codec:"forceRemoteCheck" json:"forceRemoteCheck"`
	NeedProofSet          bool                `codec:"needProofSet" json:"needProofSet"`
	AllowEmptySelfID      bool                `codec:"allowEmptySelfID" json:"allowEmptySelfID"`
	NoSkipSelf            bool                `codec:"noSkipSelf" json:"noSkipSelf"`
	CanSuppressUI         bool                `codec:"canSuppressUI" json:"canSuppressUI"`
	IdentifyBehavior      TLFIdentifyBehavior `codec:"identifyBehavior" json:"identifyBehavior"`
	ForceDisplay          bool                `codec:"forceDisplay" json:"forceDisplay"`
}

func (o IdentifyLiteArg) DeepCopy() IdentifyLiteArg {
	return IdentifyLiteArg{
		SessionID:             o.SessionID,
		Id:                    o.Id.DeepCopy(),
		Assertion:             o.Assertion,
		Reason:                o.Reason.DeepCopy(),
		UseDelegateUI:         o.UseDelegateUI,
		AlwaysBlock:           o.AlwaysBlock,
		NoErrorOnTrackFailure: o.NoErrorOnTrackFailure,
		ForceRemoteCheck:      o.ForceRemoteCheck,
		NeedProofSet:          o.NeedProofSet,
		AllowEmptySelfID:      o.AllowEmptySelfID,
		NoSkipSelf:            o.NoSkipSelf,
		CanSuppressUI:         o.CanSuppressUI,
		IdentifyBehavior:      o.IdentifyBehavior.DeepCopy(),
		ForceDisplay:          o.ForceDisplay,
	}
}

type ResolveIdentifyImplicitTeamArg struct {
	SessionID        int                 `codec:"sessionID" json:"sessionID"`
	Assertions       string              `codec:"assertions" json:"assertions"`
	Suffix           string              `codec:"suffix" json:"suffix"`
	IsPublic         bool                `codec:"isPublic" json:"isPublic"`
	DoIdentifies     bool                `codec:"doIdentifies" json:"doIdentifies"`
	Create           bool                `codec:"create" json:"create"`
	Reason           IdentifyReason      `codec:"reason" json:"reason"`
	IdentifyBehavior TLFIdentifyBehavior `codec:"identifyBehavior" json:"identifyBehavior"`
}

func (o ResolveIdentifyImplicitTeamArg) DeepCopy() ResolveIdentifyImplicitTeamArg {
	return ResolveIdentifyImplicitTeamArg{
		SessionID:        o.SessionID,
		Assertions:       o.Assertions,
		Suffix:           o.Suffix,
		IsPublic:         o.IsPublic,
		DoIdentifies:     o.DoIdentifies,
		Create:           o.Create,
		Reason:           o.Reason.DeepCopy(),
		IdentifyBehavior: o.IdentifyBehavior.DeepCopy(),
	}
}

type IdentifyInterface interface {
	// Resolve an assertion to a (UID,username) or (TeamID,teamname). On failure, returns an error.
	Resolve3(context.Context, string) (UserOrTeamLite, error)
	Identify2(context.Context, Identify2Arg) (Identify2Res, error)
	IdentifyLite(context.Context, IdentifyLiteArg) (IdentifyLiteRes, error)
	ResolveIdentifyImplicitTeam(context.Context, ResolveIdentifyImplicitTeamArg) (ResolveIdentifyImplicitTeamRes, error)
}

func IdentifyProtocol(i IdentifyInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.identify",
		Methods: map[string]rpc.ServeHandlerDescription{
			"Resolve3": {
				MakeArg: func() interface{} {
					ret := make([]Resolve3Arg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]Resolve3Arg)
					if !ok {
						err = rpc.NewTypeError((*[]Resolve3Arg)(nil), args)
						return
					}
					ret, err = i.Resolve3(ctx, (*typedArgs)[0].Assertion)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"identify2": {
				MakeArg: func() interface{} {
					ret := make([]Identify2Arg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]Identify2Arg)
					if !ok {
						err = rpc.NewTypeError((*[]Identify2Arg)(nil), args)
						return
					}
					ret, err = i.Identify2(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"identifyLite": {
				MakeArg: func() interface{} {
					ret := make([]IdentifyLiteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]IdentifyLiteArg)
					if !ok {
						err = rpc.NewTypeError((*[]IdentifyLiteArg)(nil), args)
						return
					}
					ret, err = i.IdentifyLite(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"resolveIdentifyImplicitTeam": {
				MakeArg: func() interface{} {
					ret := make([]ResolveIdentifyImplicitTeamArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ResolveIdentifyImplicitTeamArg)
					if !ok {
						err = rpc.NewTypeError((*[]ResolveIdentifyImplicitTeamArg)(nil), args)
						return
					}
					ret, err = i.ResolveIdentifyImplicitTeam(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type IdentifyClient struct {
	Cli rpc.GenericClient
}

// Resolve an assertion to a (UID,username) or (TeamID,teamname). On failure, returns an error.
func (c IdentifyClient) Resolve3(ctx context.Context, assertion string) (res UserOrTeamLite, err error) {
	__arg := Resolve3Arg{Assertion: assertion}
	err = c.Cli.Call(ctx, "keybase.1.identify.Resolve3", []interface{}{__arg}, &res)
	return
}

func (c IdentifyClient) Identify2(ctx context.Context, __arg Identify2Arg) (res Identify2Res, err error) {
	err = c.Cli.Call(ctx, "keybase.1.identify.identify2", []interface{}{__arg}, &res)
	return
}

func (c IdentifyClient) IdentifyLite(ctx context.Context, __arg IdentifyLiteArg) (res IdentifyLiteRes, err error) {
	err = c.Cli.Call(ctx, "keybase.1.identify.identifyLite", []interface{}{__arg}, &res)
	return
}

func (c IdentifyClient) ResolveIdentifyImplicitTeam(ctx context.Context, __arg ResolveIdentifyImplicitTeamArg) (res ResolveIdentifyImplicitTeamRes, err error) {
	err = c.Cli.Call(ctx, "keybase.1.identify.resolveIdentifyImplicitTeam", []interface{}{__arg}, &res)
	return
}
