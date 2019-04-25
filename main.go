package main

import (
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

	q := client.Query(`
		SELECT contracts.address, COUNT(1) AS tx_count
		FROM ` + "`bigquery-public-data.ethereum_blockchain.contracts`" + ` AS contracts
		JOIN ` + "`bigquery-public-data.ethereum_blockchain.transactions`" + ` AS transactions ON (transactions.to_address = contracts.address)
		WHERE contracts.is_erc721 = TRUE
		GROUP BY contracts.address
		ORDER BY tx_count DESC
		LIMIT 10
	`)

	it, err := q.Read(ctx)
	if err != nil {
		// TODO: Handle error.
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
