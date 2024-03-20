package cosmos

import (
	"fmt"

	wasmclient "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	solomachine "github.com/cosmos/ibc-go/v8/modules/light-clients/06-solomachine"
	tmclient "github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint"
	localhost "github.com/cosmos/ibc-go/v8/modules/light-clients/09-localhost"
)

// GetClientLatestHeight returns the latest height of the ibc ClientState
func GetClientLatestHeight(clientState ibcexported.ClientState) clienttypes.Height {
	switch cs := clientState.(type) {
	case *localhost.ClientState:
		return cs.LatestHeight
	case *solomachine.ClientState:
		// NOTE: RevisionNumber is always 0 for solomachine client heights.
		return clienttypes.NewHeight(0, cs.Sequence)
	case *tmclient.ClientState:
		return cs.LatestHeight
	case *wasmclient.ClientState:
		return cs.LatestHeight
	default:
		panic(fmt.Errorf("client type %T is unsupported", cs))
	}
}
