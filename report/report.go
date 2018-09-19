package report

import (
	"fmt"
	"github.com/pcarleton/cc-grpc/lib"
	"github.com/pcarleton/sheets"
	gsheets "google.golang.org/api/sheets/v4"
	"strings"
	"time"
)

// Steps
// 1) Fetch transactions for target range (plus a couple days on the end to catch the payment)
// 2) Turn in to TSV in proper directory
// 3) Remove payment for previous month
// 4) Write out this month's "Due"
// 5) Reconcile "Due" with the payment

// Once we're happy with that:
//
// 1) Find spreadsheet ID
// 2) Import transactions into new sheet
// 3) (Manual): Label transactions as who did them (taylor, paul, both)
// 4) Copy formatting
// 5) Copy less certain ones into separate column (or maybe just flag them at the top?)

// Statement
// Account label
// Start date
// End date

// [Label] -> [Transaction] -> [Statement] -> [Sheet]

// Trigger:
// * Payment made
// * Manual specification
// * Date passed

// Focus on manual trigger.  That looks like:
// 1) Query for transaction date + range. (Spec needs date range, plaid key, account)

// Report needs several statements, it figures out how to query to get transactions it needs.

type StatementDesc struct {
	PlaidLabel string // Nickname for plaid API key, corresponds to a login with a finandical institution
	Account    string // Nickname for the account e.g. sapphire
	StartDay   int
}

// Pink : 255, 164, 164
// Blue:  76, 161, 161
// Green: 144, 199, 35

// TODO: there's gotta be a better way
func getPink() *gsheets.Color {
	return &gsheets.Color{
		Alpha: 1.0,
		Red:   1.0,
		Green: 0.639,
		Blue:  0.639,
	}
}

func getBlue() *gsheets.Color {
	return &gsheets.Color{
		Alpha: 1.0,
		Red:   0.298,
		Green: 0.631,
		Blue:  0.631,
	}
}

func getGreen() *gsheets.Color {
	return &gsheets.Color{
		Alpha: 1.0,
		Red:   0.565,
		Green: 0.78,
		Blue:  0.137,
	}
}

func (sd *StatementDesc) GetInterval(targetMonth time.Month) (start, end time.Time) {
	targetYear := 2018
	// Subtract a second so when we ask if a date is "after" this date, it is true
	start = time.Date(targetYear, time.Month(targetMonth-1), sd.StartDay, 0, 0, 0, 0, time.Local).Add(-1 * time.Second)
	// We end on the statement date + a month, and then later exclude anything that happens after this time
	// so it ends up being 1 day before.  For instance, all transactions on the 18th through and including the 17th
	// of the next month, but not anything on the 18th.
	end = time.Date(targetYear, time.Month(targetMonth), sd.StartDay, 0, 0, 0, 0, time.Local)

	return start, end
}

func (s *Statement) MatchAccount(trans lib.TransactionRow) bool {
	return trans.Account == s.Desc.Account
}

func (s *Statement) MatchRange(trans lib.TransactionRow) bool {
	return trans.Date.After(s.Start) && trans.Date.Before(s.End)
}

type Statement struct {
	Desc         StatementDesc
	Transactions lib.TransactionTable
	Start        time.Time
	End          time.Time
	Payment      lib.TransactionRow
	Due          float64
}

func isPayment(trans lib.TransactionRow) bool {
	if strings.Contains(trans.Description, "Payment") {
		return true
	}
	return false
}

func GetStatement(c *lib.Config, targetMonth int, desc StatementDesc) (*Statement, error) {
	start, end := desc.GetInterval(time.Month(targetMonth))

	statement := Statement{Desc: desc, Start: start, End: end}

	// Add some buffer at the end to try to pick up the payment transaction
	// TODO: Make this optional
	bufferEnd := end.AddDate(0, 0, 7)

	ttable, err := c.TransactionsForRange(desc.PlaidLabel, start, bufferEnd)
	if err != nil {
		return nil, err
	}

	// Next up, 1) exclude payment transactions, 2) Figure out amount due, 3) reconcile with payment
	//  4) Print out result, run for real.

	for _, trans := range ttable {

		if !statement.MatchAccount(trans) {
			continue
		}

		payment := isPayment(trans)

		if !statement.MatchRange(trans) {
			if payment {
				statement.Payment = trans
			}
			continue
		}

		if !payment {
			statement.Transactions = append(statement.Transactions, trans)
			statement.Due += trans.Amount
		}
	}

	return &statement, nil
}

func formattingRule(word string, color *gsheets.Color, column int64, sheetId int64) *gsheets.ConditionalFormatRule {
	return &gsheets.ConditionalFormatRule{
		BooleanRule: &gsheets.BooleanRule{
			Condition: &gsheets.BooleanCondition{
				Type:   "TEXT_CONTAINS",
				Values: []*gsheets.ConditionValue{{UserEnteredValue: word}},
			},
			Format: &gsheets.CellFormat{
				BackgroundColor: color,
			},
		},
		Ranges: []*gsheets.GridRange{{StartColumnIndex: column, EndColumnIndex: column + 2, SheetId: sheetId}},
	}
}

func UploadStatement(client *sheets.Client, statement *Statement, spreadsheetId string) (*sheets.Sheet, error) {
	ss, err := client.GetSpreadsheet(spreadsheetId)
	if err != nil {
		return nil, err
	}

	tdata := statement.Transactions.ToDataArr()

	sheetTitle := fmt.Sprintf("%s-%s", statement.End.Month(), statement.Desc.Account)

	sheet := ss.GetSheet(sheetTitle)

	if sheet == nil {
		sheet, err = ss.AddSheet(sheetTitle)
		if err != nil {
			return nil, err
		}
	}

	startingPos := sheets.CellPos{Row: 0, Col: 3}
	sheet.UpdateFromPosition(tdata, startingPos)

	// TODO: this is very specific... pass in as a template
	summaryInfo := [][]string{
		{"total", "=SUM(H2:H)"},
		{"", ""},
		{"for", ""},
		{"both", "=SUMIF(I$2:I,\"=\"&A4,H$2:H)"},
		{"paul", "=SUMIF(I$2:I,\"=\"&A5,H$2:H)"},
		{"taylor", "=SUMIF(I$2:I,\"=\"&A6,H$2:H)"},
		{"splitwise", "=SUMIF(I$2:I,\"=\"&A7,H$2:H)"},
	}

	if statement.Desc.PlaidLabel == "paul" {
		summaryInfo = append(summaryInfo, []string{"taylor owes paul", "=B4/2+B6"})
	} else {
		summaryInfo = append(summaryInfo, []string{"paul owes taylor", "=B4/2+B5"})
	}

	summaryInfo = append(summaryInfo, []string{""}, []string{"sanity check", "=SUM(B4:B7)"})

	sheet.UpdateFromPosition(summaryInfo, sheet.TopLeft())

	// TODO: Freeze first row
	// TODO: Conditional formatting

	// Pink : 255, 164, 164
	// Blue:  76, 161, 161
	// Green: 144, 199, 35

	// TODO: Get sheet properties
	sheetId := sheet.Properties.SheetId
	rules := []*gsheets.ConditionalFormatRule{
		formattingRule("both", getGreen(), 8, sheetId),
		formattingRule("paul", getBlue(), 8, sheetId),
		formattingRule("taylor", getPink(), 8, sheetId),
	}

	var reqs []*gsheets.Request

	for i, rule := range rules {
		addReq := gsheets.AddConditionalFormatRuleRequest{Index: int64(i), Rule: rule}
		reqs = append(reqs, &gsheets.Request{AddConditionalFormatRule: &addReq})
	}

	// Freeze first row
	freezeRowReq := &gsheets.UpdateSheetPropertiesRequest{
		Fields: "gridProperties.frozenRowCount",
		Properties: &gsheets.SheetProperties{
			SheetId: sheetId,
			GridProperties: &gsheets.GridProperties{
				FrozenRowCount: 1,
			},
		},
	}

	reqs = append(reqs, &gsheets.Request{UpdateSheetProperties: freezeRowReq})

	_, err = sheet.Spreadsheet.DoBatch(reqs...)

	if err != nil {
		return nil, err
	}

	return sheet, nil
}

// TODO: Upload sheet to google sheets
