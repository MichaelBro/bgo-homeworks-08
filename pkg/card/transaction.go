package card

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
)

type Transaction struct {
	XMLName   string `xml:"transaction"`
	Id        uint32 `json:"id" xml:"id"`
	From      string `json:"from" xml:"from"`
	To        string `json:"to" xml:"to"`
	Amount    uint32 `json:"amount" xml:"amount"`
	Timestamp uint32 `json:"timestamp" xml:"timestamp"`
}

type Transactions struct {
	XMLName      string        `xml:"transactions"`
	Transactions []Transaction `xml:"transaction"`
}

var (
	headStringSliceForCsv = []string{"Id", "From", "To", "Amount", "Timestamp"}
	ErrWrongFile          = errors.New("wrong file")
)

func ExportToCsv(t []Transaction) error {
	file, err := os.Create("transactions.csv")

	if err != nil {
		log.Println(err)
		return err
	}

	defer func(c io.Closer) {
		if cerr := c.Close(); cerr != nil {
			log.Println(cerr)
			if err == nil {
				err = cerr
			}
		}
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(headStringSliceForCsv)

	if err != nil {
		log.Println(err)
		return err
	}

	for _, transaction := range t {
		err = writer.Write(transaction.toStringSlice())
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func (t Transaction) toStringSlice() []string {
	return []string{
		strconv.FormatUint(uint64(t.Id), 10),
		t.From,
		t.To,
		strconv.FormatUint(uint64(t.Amount), 10),
		strconv.FormatUint(uint64(t.Timestamp), 10),
	}
}

func ImportOfCsv(filePath string) ([]Transaction, error) {
	file, err := os.Open(filePath)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer func(c io.Closer) {
		if cerr := c.Close(); cerr != nil {
			log.Println(cerr)
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if !reflect.DeepEqual(records[0], headStringSliceForCsv) {
		log.Println(ErrWrongFile)
		return nil, ErrWrongFile
	}

	return MapRowToTransaction(records)
}

func MapRowToTransaction(records [][]string) ([]Transaction, error) {
	var transactions = make([]Transaction, 0)

	for i, record := range records {
		if i == 0 {
			continue
		}

		t, err := StringSliceToTransaction(record)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		transactions = append(transactions, t)
	}
	return transactions, nil
}

func StringSliceToTransaction(slice []string) (Transaction, error) {

	id, err := strconv.ParseUint(slice[0], 10, 64)
	if err != nil {
		log.Println(err)
		return Transaction{}, err
	}

	amount, err := strconv.ParseUint(slice[3], 10, 64)
	if err != nil {
		log.Println(err)
		return Transaction{}, err
	}

	timestamp, err := strconv.ParseUint(slice[4], 10, 64)
	if err != nil {
		log.Println(err)
		return Transaction{}, err
	}

	transaction := Transaction{
		Id:        uint32(id),
		From:      slice[1],
		To:        slice[2],
		Amount:    uint32(amount),
		Timestamp: uint32(timestamp),
	}
	return transaction, nil
}

func ExportToJson(t []Transaction) error {
	jsonString, err := json.Marshal(t)

	if err != nil {
		log.Println(err)
		return err
	}

	err = ioutil.WriteFile("transactions.json", jsonString, os.ModePerm)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func ImportFromJson(filePath string) ([]Transaction, error) {
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var decoded []Transaction

	err = json.Unmarshal(file, &decoded)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return decoded, nil
}

func ExportToXml(t []Transaction) error {
	transactions := Transactions{Transactions: t}
	encoded, err := xml.Marshal(transactions)
	if err != nil {
		log.Println(err)
		return err
	}

	encoded = append([]byte(xml.Header), encoded...)
	err = ioutil.WriteFile("transactions.xml", encoded, os.ModePerm)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func ImportFromXml(filePath string) ([]Transaction, error) {
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	fmt.Println(string(file))
	var decoded Transactions
	err = xml.Unmarshal(file, &decoded)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return decoded.Transactions, nil
}
