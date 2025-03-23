import axios from 'axios';
import { useEffect, useState } from 'react';

interface Company {
  name: string;
  year: number;
}

const Companies = () => {
  const [companies, setCompanies] = useState<Company[]>([]);

  useEffect(() => {
    axios.get("http://localhost:3000/companies")
      .then(res => {setCompanies(res.data)})
      .catch(error => {
        console.log(error)
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