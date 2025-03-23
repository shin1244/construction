import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

const AddCompany = () => {
  const [name, setName] = useState('');
  const [year, setYear] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const company = { name, year: parseInt(year) };
    console.log(company)

    try {
      await axios.post("http://localhost:3000/companies", company);
      navigate("/companies"); 
    } catch (error) {
      console.error("회사 추가 실패:", error);
      alert("회사 추가에 실패했습니다.");
    }
  };

  return (
    <div>
      <h1>회사를 추가하세요</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="name">회사명</label>
          <input
            id="name"
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
        </div>
        <div>
          <label htmlFor="year">설립 연도</label>
          <input
            id="year"
            type="number"
            value={year}
            onChange={(e) => setYear(e.target.value)}
          />
        </div>
        <button type="submit">회사를 추가</button>
      </form>
    </div>
  );
};

export default AddCompany;
