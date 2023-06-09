package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateProfile{}, "socialapp/CreateProfile", nil)
	cdc.RegisterConcrete(&MsgUpdateProfile{}, "socialapp/UpdateProfile", nil)
	cdc.RegisterConcrete(&MsgDeleteProfile{}, "socialapp/DeleteProfile", nil)
	cdc.RegisterConcrete(&MsgCreatePost{}, "socialapp/CreatePost", nil)
	cdc.RegisterConcrete(&MsgDeletePost{}, "socialapp/DeletePost", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateProfile{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateProfile{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteProfile{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePost{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeletePost{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
