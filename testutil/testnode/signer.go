package testnode

import (
	"github.com/sunrise-zone/sunrise-app/app/encoding"
	"github.com/sunrise-zone/sunrise-app/pkg/user"
	"github.com/sunrise-zone/sunrise-app/testutil"
	"github.com/sunrise-zone/sunrise-app/testutil/testfactory"
)

func NewOfflineSigner() (*user.Signer, error) {
	encCfg := encoding.MakeConfig(testutil.ModuleBasics)
	kr, addr := NewKeyring(testfactory.TestAccName)
	return user.NewSigner(kr, nil, addr[0], encCfg.TxConfig, testfactory.ChainID, 1, 0)
}

func NewSingleSignerFromContext(ctx Context) (*user.Signer, error) {
	encCfg := encoding.MakeConfig(testutil.ModuleBasics)
	return user.SetupSingleSigner(ctx.GoContext(), ctx.Keyring, ctx.GRPCClient, encCfg)
}

func NewSignerFromContext(ctx Context, acc string) (*user.Signer, error) {
	encCfg := encoding.MakeConfig(testutil.ModuleBasics)
	addr := testfactory.GetAddress(ctx.Keyring, acc)
	return user.SetupSigner(ctx.GoContext(), ctx.Keyring, ctx.GRPCClient, addr, encCfg)
}
