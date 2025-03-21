import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const AddCompany = () => {
  const [name, setName] = useState('');
  const [year, setYear] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const company = { name, year: parseInt(year) };

    // POST 요청으로 회사 정보 전송
    const response = await fetch('http://localhost:3000/companies', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(company),
    });

    if (response.ok) {
      navigate('/companies');  // 회사 목록 페이지로 이동
    } else {
      alert('회사 추가에 실패했습니다.');
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
