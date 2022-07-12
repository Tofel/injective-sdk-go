package main

import (
	"fmt"
	"os"
	"time"

	"github.com/InjectiveLabs/sdk-go/client/common"
	"github.com/shopspring/decimal"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
)

/**
Two orders will be broadcasted together and processed in the same block. Two different accounts will be used and second account
will grant the first one the permissions to send messages on it behalf. That way it will be possible to broadcast them together.
**/

func main() {
	network := common.LoadNetwork("devnet-1", "ignore")
	tmRPC, err := rpchttp.New(network.TmEndpoint, "/websocket")

	if err != nil {
		fmt.Println(err)
	}

	masterAddress, masterCosmosKeyring, err := chainclient.InitCosmosKeyring(
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

	masterClientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		masterAddress.String(),
		masterCosmosKeyring,
	)

	if err != nil {
		fmt.Println(err)
	}

	masterClientCtx.WithNodeURI(network.TmEndpoint)
	masterClientCtx = masterClientCtx.WithClient(tmRPC)

	masterChainClient, err := chainclient.NewChainClient(
		masterClientCtx,
		network.ChainGrpcEndpoint,
		common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("500000000inj"),
	)

	if err != nil {
		fmt.Println(err)
	}

	slaveAddress, slaveCosmosKeyring2, err := chainclient.InitCosmosKeyring(
		os.Getenv("HOME")+"/.injectived",
		"",
		"",
		"",
		"",
		"689a395a99564c7e84bd3b6e9a717433c1167ff70bb077834672bf545b8d548a",
		false,
	)

	if err != nil {
		panic(err)
	}

	saveClientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		slaveAddress.String(),
		slaveCosmosKeyring2,
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("grantee (receives permissions): %s", masterAddress.String()))
	fmt.Println(fmt.Sprintf("granter (gives permssions): %s", slaveAddress.String()))

	saveClientCtx.WithNodeURI(network.TmEndpoint)
	saveClientCtx = saveClientCtx.WithClient(tmRPC)

	slaveChainClient, err := chainclient.NewChainClient(
		saveClientCtx,
		network.ChainGrpcEndpoint,
		common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("500000000inj"),
	)

	// Slave grants master permission to send any and all messages on his behalf

	grantee := masterAddress.String()
	granter := slaveAddress.String()
	expireIn := time.Now().AddDate(1, 0, 0) // years months days

	msgtype := "/injective.exchange.v1beta1.MsgCreateSpotLimitOrder"
	msgGrant := slaveChainClient.BuildGenericAuthz(granter, grantee, msgtype, expireIn)

	rGrant, err := slaveChainClient.SyncBroadcastMsg(msgGrant)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("Grant response code: %v", rGrant.TxResponse.Code))
	fmt.Println(fmt.Sprintf("Grant trx hash: %v", rGrant.TxResponse.TxHash))

	time.Sleep(time.Second * 5)

	// let's prepare master's transaction

	masterSubaccountID := masterChainClient.DefaultSubaccount(masterAddress)

	marketId := "0x26413a70c9b78a495023e5ab8003c9cf963ef963f6755f8b57255feb5744bf31" // LINK/USDT spot

	masterOrder := masterChainClient.SpotOrder(masterSubaccountID, network, &chainclient.SpotOrderData{
		OrderType:    exchangetypes.OrderType_SELL, //BUY SELL
		Quantity:     decimal.NewFromFloat(1),
		Price:        decimal.NewFromFloat(13),
		FeeRecipient: masterAddress.String(),
		MarketId:     marketId,
	})

	msg := new(exchangetypes.MsgCreateSpotLimitOrder)
	msg.Sender = masterAddress.String()
	msg.Order = exchangetypes.SpotOrder(*masterOrder)

	// let's prepare slave's transaction and wrap it in MsgExec

	slaveSubaccountID := masterChainClient.DefaultSubaccount(slaveAddress)

	order2 := masterChainClient.SpotOrder(slaveSubaccountID, network, &chainclient.SpotOrderData{
		OrderType:    exchangetypes.OrderType_BUY, //BUY SELL
		Quantity:     decimal.NewFromFloat(1),
		Price:        decimal.NewFromFloat(4),
		FeeRecipient: slaveAddress.String(),
		MarketId:     marketId,
	})

	// time to wrap

	msg0 := exchangetypes.MsgCreateSpotLimitOrder{
		Sender: granter,
		Order:  *order2,
	}

	msg0Bytes, _ := msg0.Marshal()
	msg0Any := &codectypes.Any{}
	msg0Any.TypeUrl = sdk.MsgTypeURL(&msg0)
	msg0Any.Value = msg0Bytes
	msgExec := &authztypes.MsgExec{
		Grantee: grantee,
		Msgs:    []*codectypes.Any{msg0Any},
	}

	// broadcast both messages together so that they are processed in the same block

	resp, err := masterChainClient.SyncBroadcastMsg(msgExec, msg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("Trx hash: %v", resp.TxResponse.TxHash))
}
