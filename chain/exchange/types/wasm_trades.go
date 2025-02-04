package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

type PositionTransfer struct {
	MarketID                common.Hash `json:"market_id"`
	SourceSubaccountID      common.Hash `json:"source_subaccount_id"`
	DestinationSubaccountID common.Hash `json:"destination_subaccount_id"`
	Quantity                sdk.Dec     `json:"quantity"`
}

func (p *PositionTransfer) ValidateBasic() error {
	if p.Quantity.IsNil() || !p.Quantity.IsPositive() {
		return ErrInvalidQuantity
	}

	return nil
}

type SyntheticTradeAction struct {
	UserTrades     []*SyntheticTrade `json:"user_trades"`
	ContractTrades []*SyntheticTrade `json:"contract_trades"`
}

func (a *SyntheticTradeAction) ValidateBasic() error {
	for _, t := range a.UserTrades {
		if err := t.Validate(); err != nil {
			return err
		}
	}

	for _, t := range a.ContractTrades {
		if err := t.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type SyntheticTrade struct {
	MarketID     common.Hash `json:"market_id"`
	SubaccountID common.Hash `json:"subaccount_id"`
	IsBuy        bool        `json:"is_buy"`
	Quantity     sdk.Dec     `json:"quantity"`
	Price        sdk.Dec     `json:"price"`
	Margin       sdk.Dec     `json:"margin"`
}

func (t *SyntheticTrade) Validate() error {
	if t.Quantity.IsNil() || !t.Quantity.IsPositive() {
		return ErrInvalidQuantity
	}

	if t.Price.IsNil() || !t.Price.IsPositive() {
		return ErrInvalidPrice
	}

	// Margin can be 0 or even negative!
	if t.Margin.IsNil() {
		return ErrInvalidMargin
	}

	return nil
}

func (t *SyntheticTrade) ToPositionDelta() *PositionDelta {
	return &PositionDelta{
		IsLong:            t.IsBuy,
		ExecutionQuantity: t.Quantity,
		ExecutionMargin:   t.Margin,
		ExecutionPrice:    t.Price,
	}
}
