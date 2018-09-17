package lib

import (
	"bufio"
	"encoding/json"
	"fmt"
  "io"
	"os"
	"strconv"
	"strings"
	"time"
  "gopkg.in/yaml.v2"

	"github.com/pcarleton/cc-grpc/plaid"
)

const (
	DateFmt = plaid.DateFmt
)

type Config struct {
  Email string  `yaml:"email"`
  ClientId string `yaml:"client_id`
  ClientSecret string `yaml:"client_secret`
  // The list of accounts should be stored not in this struct
  Accounts []Account `yaml:"accounts`
}

func NewConfig(reader io.Reader) (*Config, error) {
  config := Config{}
  decoder := yaml.NewDecoder(reader)
  err := decoder.Decode(&config)
  if err != nil {
    return nil, err
  }
  return &config, nil
}

type Account struct {
  Name      string  `"yaml:"name"`
  Token string  `"yaml:"token"`
	Nicknames map[string]string  `"yaml:"nicknames"`
}

type Nicknames map[string]string

func (a *Account) NickMap(accts []plaid.Account) Nicknames {
	nickMap := make(map[string]string)

	for _, acct := range accts {
		nickMap[acct.ID] = a.Nicknames[acct.Mask]
	}

	return nickMap
}

func Masks(accts []plaid.Account) []string {
	masks := make([]string, len(accts))

	for i, acct := range accts {
		masks[i] = acct.Mask
	}

	return masks
}

func (c *Config) GetAccount(name string) *Account {
	for _, acct := range c.Accounts {
		if acct.Name == name {
			return &acct
		}
	}
	return nil
}

func (c *Config) GetClient() plaid.Client {
	// TODO: Memoize?
	return plaid.NewClient(
    c.ClientId,
    c.ClientSecret,
		plaid.DevURL)
}

type Interval struct {
	Start time.Time
	End   time.Time
}

func NDaysAgo(n int) time.Time {
	today := time.Now()
	return today.Add(time.Duration(n*-24) * time.Hour)
}

func LastNDays(n int) Interval {
	return Interval{
		Start: NDaysAgo(n),
		End:   time.Now(),
	}
}

func OutputJson(val interface{}) error {
	valj, err := json.Marshal(val)
	if err != nil {
		return err
	}
	fmt.Println(string(valj))
	return nil
}

type TransactionTable []TransactionRow

func (tt TransactionTable) ToDataArr() [][]string {
	headers := []string{
		"account",
		"date",
		"description",
		"category",
		"amount",
	}

	data := [][]string{headers}

	for _, trans := range tt {
		pieces := []string{
			trans.Account,
			trans.Date.Format(DateFmt),
			trans.Description,
			trans.Category,
			fmt.Sprintf("%.2f", trans.Amount),
		}

		data = append(data, pieces)
	}
	return data
}


func (tt TransactionTable) Write(delimiter string) {
	// TODO: Write to arbitrary output
	strData := tt.ToDataArr()

	for _, row := range strData {
		fmt.Println(strings.Join(row, delimiter))
	}
}

type TransactionRow struct {
	Account     string // Nick name, human readable
	Date        time.Time
	Description string
	Category    string
	Label       string
	Amount      float64
}

func (t *TransactionRow) Copy() TransactionRow {
  return TransactionRow{
    t.Account, t.Date, t.Description, t.Category, t.Label, t.Amount,
  }
}

func (acct *Account) TableFromPlaidTrans(resp plaid.TransactionResponse) (TransactionTable, error) {
	nickMap := acct.NickMap(resp.Accounts)

	table := make([]TransactionRow, 0)

	for _, trans := range resp.Transactions {
		date, err := time.ParseInLocation(DateFmt, trans.Date, time.Local)
		if err != nil {
			return nil, fmt.Errorf("Invalid date: %s on transaction %s", trans.Date, trans.ID)
		}

		table = append(table, TransactionRow{
			Account:     nickMap[trans.AccountID],
			Date:        date,
			Description: trans.Name,
			Category:    strings.Join(trans.Category, ":"),
			Amount:      trans.Amount,
		})
	}

	return table, nil
}

func (c *Config) TransactionsForRange(acctLabel string, start, end time.Time) (TransactionTable, error) {
	acct := c.GetAccount(acctLabel)

	if acct == nil {
    return nil, fmt.Errorf("No account: %s", acctLabel)
	}

	client := c.GetClient()

	resp, err := client.Transactions(acct.Token, start, end)

	if err != nil {
	    return nil, err
	}

	return acct.TableFromPlaidTrans(resp)
}

func ReadTransactionTable(fileName string) (TransactionTable, error) {
	reader, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	headers := strings.Split(scanner.Text(), "\t")

	table := make([]TransactionRow, 0)

	for scanner.Scan() {
		line := scanner.Text()
		pieces := strings.Split(line, "\t")
		trans := make(map[string]string)

		for idx, header := range headers {
			trans[header] = pieces[idx]
		}

		date, err := time.ParseInLocation("2006-01-02", trans["date"], time.Local)
		if err != nil {
			return nil, fmt.Errorf("Invalid date %s in line {%s}", trans["date"], line)
		}

		amount, err := strconv.ParseFloat(trans["amount"], 64)
		if err != nil {
			return nil, fmt.Errorf("Invalid amount %s in line {%s}", trans["amount"], line)
		}
		table = append(table, TransactionRow{
			Account:     trans["account"],
			Date:        date,
			Description: trans["description"],
			Category:    trans["category"],
			Amount:      amount,
		})
	}

	return table, nil
}
