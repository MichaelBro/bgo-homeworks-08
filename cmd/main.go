package main

import (
	"bgo-homeworks-08/pkg/card"
	"fmt"
	"os"
)

func main() {

	transactions := []card.Transaction{
		{
			Id:        0,
			From:      "5555 1232 2222 5555",
			To:        "2345 7874 7437 2232",
			Amount:    600_00,
			Timestamp: 1613983040,
		},
		{
			Id:        1,
			From:      "1234 4321 1234 5666",
			To:        "2345 7874 2222 7874",
			Amount:    100_00,
			Timestamp: 1613989232,
		},
		{
			Id:        2,
			From:      "2345 7874 2222 7874",
			To:        "1234 4321 1234 2222",
			Amount:    1_000_00,
			Timestamp: 1613989841,
		},
	}

	fmt.Println(transactions)

	err := card.ExportToCsv(transactions)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	transactionsByCsv, err := card.ImportOfCsv("transactions.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(transactionsByCsv)

	err = card.ExportToJson(transactions)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	transactionsByJson, err := card.ImportFromJson("transactions.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(transactionsByJson)

	err = card.ExportToXml(transactions)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	transactionsByXml, err := card.ImportFromXml("transactions.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(transactionsByXml)

}
