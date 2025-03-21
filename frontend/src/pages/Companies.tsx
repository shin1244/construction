import { useEffect, useState } from 'react';

interface Company {
  name: string;
  year: number;
}

const Companies = () => {
  const [companies, setCompanies] = useState<Company[]>([]);

  useEffect(() => {
    fetch('http://localhost:3000/companies')
      .then(resp => resp.json())
      .then(data => {
        console.log(data)
        setCompanies(data);
      })
  }, [])
    return (
    <div>
      <h1>회사 목록</h1>
      {companies.length === 0 ? (
        <p>회사가 없습니다.</p>
      ) : (
        <ul>
          {companies.map((company, index) => (
            <li key={index}>
              <h3>{company.name}</h3>
              <p>설립 연도: {company.year}</p>
            </li>
          ))}
        </ul>
      )}
    </div>
    );
  };
  
  export default Companies;