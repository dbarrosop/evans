package gateway

import (
	"path/filepath"
	"testing"

	"github.com/jhump/protoreflect/dynamic"
	"github.com/k0kubun/pp"
	"github.com/ktr0731/evans/entity"
	"github.com/ktr0731/evans/tests/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockPrompt struct{}

func (p *mockPrompt) Input() string {
	return "foo"
}

func TestPromptInputter_Input(t *testing.T) {
	setup := func(t *testing.T, fpath, pkgName, svcName string) *entity.Env {
		set := helper.ReadProto(t, []string{fpath})

		env := helper.NewEnv(t, set, helper.TestConfig().Env)

		err := env.UsePackage(pkgName)
		require.NoError(t, err)

		err = env.UseService(svcName)
		require.NoError(t, err)

		return env
	}

	t.Run("normal/simple", func(t *testing.T) {
		env := setup(t, filepath.Join("helloworld", "helloworld.proto"), "helloworld", "Greeter")

		inputter := &PromptInputter{newPromptInputter(&mockPrompt{}, env)}

		rpc, err := env.RPC("SayHello")
		require.NoError(t, err)

		dmsg, err := inputter.Input(rpc.RequestType)
		require.NoError(t, err)

		msg, ok := dmsg.(*dynamic.Message)
		require.True(t, ok)

		assert.Equal(t, `name:"foo" message:"foo"`, msg.String())
	})

	// t.Run("normal/nested_message", func(t *testing.T) {
	// 	env := setup(t, filepath.Join("nested_message", "library.proto"), "library", "Library")
	//
	// 	inputter := &PromptInputter{newPromptInputter(&mockPrompt{}, env)}
	//
	// 	rpc, err := env.RPC("BorrowBook")
	// 	require.NoError(t, err)
	//
	// 	dmsg, err := inputter.Input(rpc.RequestType)
	// 	require.NoError(t, err)
	//
	// 	msg, ok := dmsg.(*dynamic.Message)
	// 	require.True(t, ok)
	//
	// 	assert.Equal(t, `name:"foo" message:"foo"`, msg.String())
	// })

	t.Run("foo", func(t *testing.T) {
		env := setup(t, filepath.Join("nested_message", "library.proto"), "library", "Library")

		rpc, err := env.RPC("BorrowBook")
		require.NoError(t, err)

		for _, f := range rpc.RequestType.GetFields() {
			pp.Println(f.GetName())
			pp.Println(f.GetMessageType().GetFullyQualifiedName())
		}
	})
}
