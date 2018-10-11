package usecase

import (
	"testing"

	"github.com/ktr0731/evans/adapter/presenter"
	"github.com/ktr0731/evans/entity"
	"github.com/ktr0731/evans/tests/mock/entity/mockenv"
	"github.com/ktr0731/evans/usecase/port"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeader(t *testing.T) {
	params := &port.HeaderParams{
		Headers: []*entity.Header{
			{Key: "tsukasa", Val: "ayatsuji"},
			{Key: "miya", Val: "tachibana", NeedToRemove: true},
		},
	}
	presenter := presenter.NewStubPresenter()
	env := &mockenv.EnvironmentMock{
		AddHeaderFunc:    func(*entity.Header) {},
		RemoveHeaderFunc: func(string) {},
	}

	r, err := Header(params, presenter, env)
	require.NoError(t, err)
	assert.Equal(t, r, nil)

	assert.Len(t, env.RemoveHeaderCalls(), 1)
}
