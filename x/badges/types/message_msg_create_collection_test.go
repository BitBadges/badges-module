package types

import (
	"testing"

	"github.com/bitbadges/badges-module/x/badges/testutil/sample"

	"github.com/stretchr/testify/require"
)

func TestMsgCreateCollection_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateCollection
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateCollection{
				Creator: "invalid_address",
			},
			err: ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateCollection{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
