package types_test

import (
	"testing"

	"blit/x/blit/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				TaskList: []types.Task{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				TaskResultList: []types.TaskResult{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				FutureTaskList: []types.FutureTask{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated task",
			genState: &types.GenesisState{
				TaskList: []types.Task{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated taskResult",
			genState: &types.GenesisState{
				TaskResultList: []types.TaskResult{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated futureTask",
			genState: &types.GenesisState{
				FutureTaskList: []types.FutureTask{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
