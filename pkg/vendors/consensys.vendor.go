package vendors

import (
	"context"
	"fmt"
	"github.com/SidVermaS/Ethereum-Consensus/pkg/types/structs"

	eth2client "github.com/attestantio/go-eth2-client"
	"github.com/attestantio/go-eth2-client/http"
	"github.com/rs/zerolog"
)

type Consensys struct {
	Vendor structs.Vendor
}

func (consensys *Consensys) AccessConsensysNode() {
	if consensys != nil {

	}
	fmt.Println("~~~ 1 AccessConsensysNode")
	// Provide a cancellable context to the creation function.
	ctx, cancel := context.WithCancel(context.Background())
	client, err := http.New(ctx,
		// WithAddress supplies the address of the beacon node, as a URL.
		http.WithAddress("http://localhost:5051/"),
		// LogLevel supplies the level of logging to carry out.
		http.WithLogLevel(zerolog.WarnLevel),
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connected to %s\n", client.Name())

	// Client functions have their own interfaces.  Not all functions are
	// supported by all clients, so checks should be made for each function when
	// casting the service to the relevant interface.
	if provider, isProvider := client.(eth2client.GenesisProvider); isProvider {
		genesis, err := provider.Genesis(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Genesis time is %v\n", genesis.GenesisTime)
	}

	// You can also access the struct directly if required.
	httpClient := client.(*http.Service)
	genesis, err := httpClient.Genesis(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Genesis validators root is %#x\n", genesis.GenesisValidatorsRoot)

	// Cancelling the context passed to New() frees up resources held by the
	// client, closes connections, clears handlers, etc.
	cancel()
}
