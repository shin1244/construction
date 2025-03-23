package models

type memoryHandler struct {
	companyMap map[int]*Company
}

func (m *memoryHandler) getCompanies() []*Company {
	list := []*Company{}
	for _, v := range m.companyMap {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler) addCompany(company *Company) *Company {
	id := len(m.companyMap) + 1
	company.ID = id
	m.companyMap[id] = company
	return company
}

func (m *memoryHandler) removeCompany(id int) bool {
	if _, ok := m.companyMap[id]; ok {
		delete(m.companyMap, id)
		return true
	} else {
		return false
	}
}
