package models

type Company struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Year int    `json:"year"`
}

type dbHandler interface {
	getCompanies() []*Company
	addCompany(company *Company) *Company
	removeCompany(id int) bool
}

var handler dbHandler

func newMemoryHandler() dbHandler {
	m := &memoryHandler{}
	m.companyMap = make(map[int]*Company)
	return m
}

func init() {
	handler = newMemoryHandler()
}

func GetCompanies() []*Company {
	return handler.getCompanies()
}

func AddCompany(company *Company) *Company {
	return handler.addCompany(company)
}

func RemoveCompany(id int) bool {
	return handler.removeCompany(id)
}
