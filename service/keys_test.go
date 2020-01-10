package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeys(t *testing.T) {
	// SetLogger(NewLogger(DebugLevel))
	env := newTestEnv(t)
	ctx := context.TODO()

	// Alice
	service, closeFn := newTestService(t, env)
	defer closeFn()

	testAuthSetup(t, service, alice)
	testUserSetup(t, env, service, alice, "alice")
	testPush(t, service, alice)

	testImportKey(t, service, bob)
	testUserSetup(t, env, service, bob, "bob")
	testPush(t, service, bob)

	testImportKey(t, service, charlie)
	testUserSetup(t, env, service, charlie, "charlie")
	testPush(t, service, charlie)

	// Default
	resp, err := service.Keys(ctx, &KeysRequest{})
	require.NoError(t, err)
	require.Equal(t, "user", resp.SortField)
	require.Equal(t, SortAsc, resp.SortDirection)
	require.Equal(t, 3, len(resp.Keys))
	require.Equal(t, "kpe132yw8ht5p8cetl2jmvknewjawt9xwzdlrk2pyxlnwjyqrdq0dawqlrnuen", resp.Keys[0].ID)
	require.Equal(t, 1, len(resp.Keys[0].Users))
	require.Equal(t, "alice", resp.Keys[0].Users[0].Name)
	require.Equal(t, PrivateKeyType, resp.Keys[0].Type)
	require.Equal(t, "kpe1syuhwr4g05t4744r23nvxnr7en9cmz53knhr0gja7c84hr7fkw2qrt73l9", resp.Keys[1].ID)
	require.Equal(t, 1, len(resp.Keys[1].Users))
	require.Equal(t, "bob", resp.Keys[1].Users[0].Name)
	require.Equal(t, "kpe1a4yj333g68pvd6hfqvufqkv4vy54jfe6t33ljd3kc9rpfty8xlgs474npw", resp.Keys[2].ID)
	require.Equal(t, 1, len(resp.Keys[2].Users))
	require.Equal(t, "charlie", resp.Keys[2].Users[0].Name)

	// KID (asc)
	resp, err = service.Keys(ctx, &KeysRequest{
		SortField: "kid",
	})
	require.NoError(t, err)
	require.Equal(t, "kid", resp.SortField)
	require.Equal(t, SortAsc, resp.SortDirection)
	require.Equal(t, 3, len(resp.Keys))
	require.Equal(t, "kpe132yw8ht5p8cetl2jmvknewjawt9xwzdlrk2pyxlnwjyqrdq0dawqlrnuen", resp.Keys[0].ID)
	require.Equal(t, "kpe1a4yj333g68pvd6hfqvufqkv4vy54jfe6t33ljd3kc9rpfty8xlgs474npw", resp.Keys[1].ID)
	require.Equal(t, "kpe1syuhwr4g05t4744r23nvxnr7en9cmz53knhr0gja7c84hr7fkw2qrt73l9", resp.Keys[2].ID)

	// KID (desc)
	resp, err = service.Keys(ctx, &KeysRequest{
		SortField:     "kid",
		SortDirection: SortDesc,
	})
	require.NoError(t, err)
	require.Equal(t, "kid", resp.SortField)
	require.Equal(t, SortDesc, resp.SortDirection)
	require.Equal(t, 3, len(resp.Keys))
	require.Equal(t, "kpe1syuhwr4g05t4744r23nvxnr7en9cmz53knhr0gja7c84hr7fkw2qrt73l9", resp.Keys[0].ID)
	require.Equal(t, "kpe1a4yj333g68pvd6hfqvufqkv4vy54jfe6t33ljd3kc9rpfty8xlgs474npw", resp.Keys[1].ID)
	require.Equal(t, "kpe132yw8ht5p8cetl2jmvknewjawt9xwzdlrk2pyxlnwjyqrdq0dawqlrnuen", resp.Keys[2].ID)

	// User (asc)
	resp, err = service.Keys(ctx, &KeysRequest{
		SortField: "user",
	})
	require.NoError(t, err)
	require.Equal(t, "user", resp.SortField)
	require.Equal(t, SortAsc, resp.SortDirection)
	require.Equal(t, 3, len(resp.Keys))
	require.Equal(t, "kpe132yw8ht5p8cetl2jmvknewjawt9xwzdlrk2pyxlnwjyqrdq0dawqlrnuen", resp.Keys[0].ID)
	require.Equal(t, "kpe1syuhwr4g05t4744r23nvxnr7en9cmz53knhr0gja7c84hr7fkw2qrt73l9", resp.Keys[1].ID)
	require.Equal(t, "kpe1a4yj333g68pvd6hfqvufqkv4vy54jfe6t33ljd3kc9rpfty8xlgs474npw", resp.Keys[2].ID)

	// User (desc)
	resp, err = service.Keys(ctx, &KeysRequest{
		SortField:     "user",
		SortDirection: SortDesc,
	})
	require.NoError(t, err)
	require.Equal(t, "user", resp.SortField)
	require.Equal(t, SortDesc, resp.SortDirection)
	require.Equal(t, 3, len(resp.Keys))
	require.Equal(t, "kpe1a4yj333g68pvd6hfqvufqkv4vy54jfe6t33ljd3kc9rpfty8xlgs474npw", resp.Keys[0].ID)
	require.Equal(t, "kpe1syuhwr4g05t4744r23nvxnr7en9cmz53knhr0gja7c84hr7fkw2qrt73l9", resp.Keys[1].ID)
	require.Equal(t, "kpe132yw8ht5p8cetl2jmvknewjawt9xwzdlrk2pyxlnwjyqrdq0dawqlrnuen", resp.Keys[2].ID)

	// Type
	resp, err = service.Keys(ctx, &KeysRequest{
		SortField: "type",
	})
	require.NoError(t, err)
	require.Equal(t, "type", resp.SortField)
	require.Equal(t, SortAsc, resp.SortDirection)
	require.Equal(t, 3, len(resp.Keys))
	require.Equal(t, "kpe132yw8ht5p8cetl2jmvknewjawt9xwzdlrk2pyxlnwjyqrdq0dawqlrnuen", resp.Keys[0].ID)
	require.Equal(t, "kpe1syuhwr4g05t4744r23nvxnr7en9cmz53knhr0gja7c84hr7fkw2qrt73l9", resp.Keys[1].ID)
	require.Equal(t, "kpe1a4yj333g68pvd6hfqvufqkv4vy54jfe6t33ljd3kc9rpfty8xlgs474npw", resp.Keys[2].ID)
}

func TestKeysMissingSigchain(t *testing.T) {
	env := newTestEnv(t)
	service, closeFn := newTestService(t, env)
	defer closeFn()
	ctx := context.TODO()

	testAuthSetup(t, service, alice)
	testUserSetup(t, env, service, alice, "alice")
	testPush(t, service, alice)

	_, err := service.scs.DeleteSigchain(alice.ID())
	require.NoError(t, err)

	resp, err := service.Keys(ctx, &KeysRequest{})
	require.NoError(t, err)
	require.Equal(t, 1, len(resp.Keys))
}
