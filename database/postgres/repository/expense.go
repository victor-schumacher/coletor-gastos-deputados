package repository

import (
	"coletor-gastos-deputados/database"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Expense struct {
	Date            string  `csv:"datemissao"`
	Legislatura     string  `csv:"nulegislatura"`
	Partido         string  `csv:"sgpartido"`
	NomeParlamentar string  `csv:"txnomeparlamentar"`
	CPFCNPJ         string  `csv:"txtcnpjcpf"`
	Description     string  `csv:"txtdescricao"`
	Provider        string  `csv:"txtfornecedor"`
	Value           float32 `csv:"vlrliquido"`
}

type Manager interface {
	Save(expenses []Expense) error
}

type ExpenseRepo struct {
	db database.DBConnection
}

func NewExpense(db database.DBConnection) ExpenseRepo {
	return ExpenseRepo{db: db}
}

func (er ExpenseRepo) Save(expenses []Expense) error {
	db := er.db.ConnectHandle()
	defer db.Close()

	var valueStrings []string
	var valueArgs []interface{}
	for _, expense := range expenses {
		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?, ?, ?, ?)")
		valueArgs = append(valueArgs, id)
		valueArgs = append(valueArgs, expense.Date)
		valueArgs = append(valueArgs, expense.Legislatura)
		valueArgs = append(valueArgs, expense.Partido)
		valueArgs = append(valueArgs, expense.NomeParlamentar)
		valueArgs = append(valueArgs, expense.CPFCNPJ)
		valueArgs = append(valueArgs, expense.Description)
		valueArgs = append(valueArgs, expense.Provider)
		valueArgs = append(valueArgs, expense.Value)
	}
	values := ReplaceSQL(strings.Join(valueStrings, ","), "?")
	stmt := fmt.Sprintf("INSERT INTO deputados.gastos VALUES %s", values)
	fmt.Println("sttm " + stmt)
	if _, err := db.Exec(
		stmt, valueArgs...
	); err != nil {
		return err
	}

	return nil
}

func ReplaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}