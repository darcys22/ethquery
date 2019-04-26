package main

import (
	"bytes"
	"fmt"

	"cloud.google.com/go/bigquery"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

func main() {

	ctx := context.Background()

	projectID := "festive-planet-220111"

	client, err := bigquery.NewClient(ctx, projectID)

	if err != nil {
		fmt.Println(err)
	}

	//address := "0x9675Bc053B9Eb84EA23Dbced85F2B66Bd9daCa18"
	var query bytes.Buffer

	//query.WriteString(fmt.Sprintf("SELECT hash, value, receipt_gas_used "))
	query.WriteString(fmt.Sprintf("SELECT hash "))
	query.WriteString(fmt.Sprintf("FROM `bigquery-public-data.ethereum_blockchain.transactions` "))
	//query.WriteString(fmt.Sprintf("WHERE from_address='%s' or to_address='%s' ", address, address))
	//query.WriteString(fmt.Sprintf("ORDER BY block_number "))
	query.WriteString(fmt.Sprintf("LIMIT 10 "))

	fmt.Print(query.String())
	q := client.Query(query.String())

	it, err := q.Read(ctx)
	if err != nil {
		// TODO: Handle error.
	} else {
		fmt.Println("succeeded")
	}

	for {
		var values []bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		fmt.Println(values)
	}

	if err != nil {
		// TODO: Handle error.
	}

}
