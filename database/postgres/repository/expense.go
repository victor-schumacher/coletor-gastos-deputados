package repository

import (
	"coletor-gastos-deputados/database"

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
	Save(expense Expense) error
}

type ExpenseRepo struct {
	db database.DBConnection
}

func NewExpense(db database.DBConnection) ExpenseRepo {
	return ExpenseRepo{db: db}
}

func (er ExpenseRepo) Save(expense Expense) error {
	db := er.db.ConnectHandle()
	defer db.Close()
	stmt := "INSERT INTO deputados.gastos VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	if _, err := db.Exec(
		stmt,
		id,
		expense.Date,
		expense.Legislatura,
		expense.Partido,
		expense.NomeParlamentar,
		expense.CPFCNPJ,
		expense.Description,
		expense.Provider,
		expense.Value,
	); err != nil {
		return err
	}

	return nil
}
