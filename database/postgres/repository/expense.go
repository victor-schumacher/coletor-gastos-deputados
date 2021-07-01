package repository

import (
	"coletor-gastos-deputados/database"
	"fmt"
	"github.com/rs/zerolog/log"
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
	Clean() error
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
			log.Err(err).Msg("generate uuid error, check repository/expense.go")
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
	values := replaceSQL(strings.Join(valueStrings, ","), "?")
	stmt := fmt.Sprintf("INSERT INTO gastos VALUES %s", values)
	if _, err := db.Exec(
		stmt, valueArgs...,
	); err != nil {
		log.Err(err).Msg("query error, check repository/expense.go Save func")
		return err
	}
	log.Info().Msgf("successfuly saved expenses")
	return nil
}

func (er ExpenseRepo) Clean() error {
	db := er.db.ConnectHandle()
	defer db.Close()
	if _, err := db.Exec("TRUNCATE table gastos"); err != nil {
		log.Err(err).Msg("query error, check repository/expense.go Save func")
		return err
	}
	log.Info().Msgf("successfuly cleaned table gastos")
	return nil
}

func replaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}
