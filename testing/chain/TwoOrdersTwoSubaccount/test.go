package main

import (
	"fmt"
	"os"
	"time"

	"github.com/InjectiveLabs/sdk-go/client/common"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/**
Two orders will be broadcasted from the same address, but using two different subaccounts. And before that happens we will deposit some USDT
on the second subaccount.
**/

func main() {
	network := common.LoadNetwork("devnet-1", "ignore")
	tmRPC, err := rpchttp.New(network.TmEndpoint, "/websocket")

	if err != nil {
		fmt.Println(err)
	}

	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
		os.Getenv("HOME")+"/.injectived",
		"",
		"",
		"",
		"",
		"369dda7bfb4f3e0d10d64e1425b5e8008e1cd9f0208e7e5826d360358148256c", // private key will be used directly, we don't need other variables
		false,
	)

	if err != nil {
		panic(err)
	}

	clientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		senderAddress.String(),
		cosmosKeyring,
	)

	if err != nil {
		fmt.Println(err)
	}

	clientCtx.WithNodeURI(network.TmEndpoint)
	clientCtx = clientCtx.WithClient(tmRPC)

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("500000000inj"),
	)

	if err != nil {
		fmt.Println(err)
	}

	defaultSubaccountID := chainClient.DefaultSubaccount(senderAddress)

	defaultSubaccountAsString := defaultSubaccountID.String()
	asStringShort := string(defaultSubaccountAsString[:len(defaultSubaccountAsString)-1])
	firstSubaccountIDasString := asStringShort + "1"
	firstSubaccountID := ethcommon.HexToHash(firstSubaccountIDasString)

	msgDeposit := &exchangetypes.MsgDeposit{
		Sender:       senderAddress.String(),
		SubaccountId: firstSubaccountIDasString,
		Amount: sdk.Coin{
			Denom: "peggy0xdAC17F958D2ee523a2206206994597C13D831ec7", Amount: sdk.NewInt(10000000), // 10 USDT
		},
	}

	dResp, err := chainClient.SyncBroadcastMsg(msgDeposit)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("Deposit trx hash: %v", dResp.TxResponse.TxHash))

	time.Sleep(time.Second * 5)

	//let's prepare first order

	marketId := "0x26413a70c9b78a495023e5ab8003c9cf963ef963f6755f8b57255feb5744bf31" // LINK/USDT spot

	sellOrder := chainClient.SpotOrder(defaultSubaccountID, network, &chainclient.SpotOrderData{
		OrderType:    exchangetypes.OrderType_SELL,
		Quantity:     decimal.NewFromFloat(1),
		Price:        decimal.NewFromFloat(13),
		FeeRecipient: senderAddress.String(),
		MarketId:     marketId,
	})

	sellMsg := new(exchangetypes.MsgCreateSpotLimitOrder)
	sellMsg.Sender = senderAddress.String()
	sellMsg.Order = exchangetypes.SpotOrder(*sellOrder)

	// let's create second order

	buyOrder := chainClient.SpotOrder(firstSubaccountID, network, &chainclient.SpotOrderData{
		OrderType:    exchangetypes.OrderType_BUY,
		Quantity:     decimal.NewFromFloat(1),
		Price:        decimal.NewFromFloat(4),
		FeeRecipient: senderAddress.String(),
		MarketId:     marketId,
	})

	buyMsg := new(exchangetypes.MsgCreateSpotLimitOrder)
	buyMsg.Sender = senderAddress.String()
	buyMsg.Order = exchangetypes.SpotOrder(*buyOrder)

	resp, err := chainClient.SyncBroadcastMsg(sellMsg, buyMsg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("Trx hash: %v", resp.TxResponse.TxHash))
}
