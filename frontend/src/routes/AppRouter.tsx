import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Companies from "../pages/Companies";
import AddCompany from "../pages/AddCompany"

const AppRouter = () => {
  return (
    <Router>
      <Routes>
        <Route path="/companies" element={<Companies />} />
        <Route path="/companies/new" element={<AddCompany />} />
      </Routes>
    </Router>
  );
};

export default AppRouter;
