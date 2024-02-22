package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/wavesplatform/gowaves/pkg/client"
	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func sendAsset(amount uint64, assetId string, recipient string) error {
	var networkByte byte
	var nodeURL string

	networkByte = 55
	nodeURL = AnoteNodeURL

	// Create sender's public key from BASE58 string
	sender, err := crypto.NewPublicKeyFromBase58(conf.PublicKey)
	if err != nil {
		log.Println(err)
		return err
	}

	// Create sender's private key from BASE58 string
	sk, err := crypto.NewSecretKeyFromBase58(conf.PrivateKey)
	if err != nil {
		log.Println(err)
		return err
	}

	// Current time in milliseconds
	ts := time.Now().Unix() * 1000

	asset, err := proto.NewOptionalAssetFromString(assetId)
	if err != nil {
		log.Println(err)
		return err
	}

	assetW, err := proto.NewOptionalAssetFromString("")
	if err != nil {
		log.Println(err)
		return err
	}

	rec, err := proto.NewAddressFromString(recipient)
	if err != nil {
		log.Println(err)
		return err
	}

	tr := proto.NewUnsignedTransferWithSig(sender, *asset, *assetW, uint64(ts), amount, Fee, proto.Recipient{Address: &rec}, nil)

	err = tr.Sign(networkByte, sk)
	if err != nil {
		log.Println(err)
		return err
	}

	// Create new HTTP client to send the transaction to public TestNet nodes
	client, err := client.NewClient(client.Options{BaseUrl: nodeURL, Client: &http.Client{}})
	if err != nil {
		log.Println(err)
		return err
	}

	// Context to cancel the request execution on timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// // Send the transaction to the network
	_, err = client.Transactions.Broadcast(ctx, tr)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func dataTransaction(key string, valueStr *string, valueInt *int64, valueBool *bool) error {
	// Create sender's public key from BASE58 string
	sender, err := crypto.NewPublicKeyFromBase58(conf.PublicKey)
	if err != nil {
		log.Println(err)
		return err
	}

	// Create sender's private key from BASE58 string
	sk, err := crypto.NewSecretKeyFromBase58(conf.PrivateKey)
	if err != nil {
		log.Println(err)
		return err
	}

	a, _ := proto.NewAddressFromPublicKey(55, sender)
	log.Println(a)

	// Current time in milliseconds
	ts := time.Now().Unix() * 1000

	tr := proto.NewUnsignedDataWithProofs(2, sender, Fee, uint64(ts))

	if valueStr == nil && valueInt == nil && valueBool == nil {
		tr.Entries = append(tr.Entries,
			&proto.DeleteDataEntry{
				Key: key,
			},
		)
	}

	if valueStr != nil {
		tr.Entries = append(tr.Entries,
			&proto.StringDataEntry{
				Key:   key,
				Value: *valueStr,
			},
		)
	}

	if valueInt != nil {
		tr.Entries = append(tr.Entries,
			&proto.IntegerDataEntry{
				Key:   key,
				Value: *valueInt,
			},
		)
	}

	if valueBool != nil {
		tr.Entries = append(tr.Entries,
			&proto.BooleanDataEntry{
				Key:   key,
				Value: *valueBool,
			},
		)
	}

	err = tr.Sign(55, sk)
	if err != nil {
		log.Println(err)
		return err
	}

	// Create new HTTP client to send the transaction to public TestNet nodes
	cl, err := client.NewClient(client.Options{BaseUrl: AnoteNodeURL, Client: &http.Client{}})
	if err != nil {
		log.Println(err)
		return err
	}

	// Context to cancel the request execution on timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// // Send the transaction to the network
	_, err = cl.Transactions.Broadcast(ctx, tr)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
