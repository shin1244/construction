package models

type Company struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Year int    `json:"year"`
}

// 회사 목록을 가져오는 함수 예시
func GetAllCompanies() []Company {
	return []Company{
		{ID: 1, Name: "Company A", Year: 2000},
		{ID: 2, Name: "Company B", Year: 2010},
	}
}
