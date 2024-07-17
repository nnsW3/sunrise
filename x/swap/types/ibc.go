package types

import (
	"encoding/json"
	"fmt"
	"strings"

	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
)

const DefaultRetryCount uint8 = 3

func (m *ForwardMetadata) Validate() error {
	if m.Receiver == "" {
		return fmt.Errorf("failed to validate metadata. receiver cannot be empty")
	}
	if err := host.PortIdentifierValidator(m.Port); err != nil {
		return fmt.Errorf("failed to validate metadata: %w", err)
	}
	if err := host.ChannelIdentifierValidator(m.Channel); err != nil {
		return fmt.Errorf("failed to validate metadata: %w", err)
	}

	return nil
}

func (m *SwapMetadata) Validate() error {
	if err := m.Route.Validate(); err != nil {
		return err
	}

	if m.ExactAmountIn != nil && m.ExactAmountOut != nil {
		return fmt.Errorf("cannot have both exact_amount_in and exact_amount_out")
	}

	if m.ExactAmountIn == nil && m.ExactAmountOut == nil {
		return fmt.Errorf("must have either exact_amount_in or exact_amount_out")
	}

	if m.ExactAmountIn != nil {
		if !m.ExactAmountIn.MinAmountOut.IsPositive() {
			return fmt.Errorf("min_amount_out must be positive")
		}
	}

	if m.ExactAmountOut != nil {
		if m.ExactAmountOut.Change != nil {
			if err := m.ExactAmountOut.Change.Validate(); err != nil {
				return err
			}
		}
	}

	if m.Forward != nil {
		if err := m.Forward.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type SwapAcknowledgement struct {
	Result      RouteResult `json:"result"`
	IncomingAck []byte      `json:"ibc_ack"`
	ChangeAck   []byte      `json:"change_ack,omitempty"`
	ForwardAck  []byte      `json:"forward_ack,omitempty"`
}

func (a SwapAcknowledgement) Acknowledgement() ([]byte, error) {
	bz, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

func GetDenomForThisChain(port, channel, counterpartyPort, counterpartyChannel, denom string) string {
	counterpartyPrefix := transfertypes.GetDenomPrefix(counterpartyPort, counterpartyChannel)
	if strings.HasPrefix(denom, counterpartyPrefix) {
		// unwind denom
		unwoundDenom := denom[len(counterpartyPrefix):]
		denomTrace := transfertypes.ParseDenomTrace(unwoundDenom)
		if denomTrace.Path == "" {
			// denom is now unwound back to native denom
			return unwoundDenom
		}
		// denom is still IBC denom
		return denomTrace.IBCDenom()
	}
	// append port and channel from this chain to denom
	prefixedDenom := transfertypes.GetDenomPrefix(port, channel) + denom
	return transfertypes.ParseDenomTrace(prefixedDenom).IBCDenom()
}
