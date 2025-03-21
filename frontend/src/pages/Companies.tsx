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
      <h1>시공사 목록</h1>
      <article className='list'>
        {companies.map((company, index) => (
          <article key={index}>
            <h3>{company.name}</h3>
            <p>설립 연도: {company.year}</p>
          </article>
        ))}
      </article>
    </div>
    );
  };
  
  export default Companies;