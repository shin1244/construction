package models

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type sqlHandler struct {
	db *sql.DB
}

func (s *sqlHandler) getCompanies() []*Company {
	companies := []*Company{}
	rows, err := s.db.Query("SELECT * FROM company")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		company := &Company{}
		err := rows.Scan(&company.ID, &company.Name, &company.Year)
		if err != nil {
			panic(err)
		}
		companies = append(companies, company)
	}

	return companies
}

func (s *sqlHandler) addCompany(company *Company) *Company {
	return nil
}

func (s *sqlHandler) removeCompany(id int) bool {
	return false
}

func (s *sqlHandler) close() {
	s.db.Close()
}

func newMySqlHandler() dbHandler {
	dsn := os.Getenv("DB_DSN")
	database, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return &sqlHandler{db: database}
}
