import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Companies from "../pages/Companies";

const AppRouter = () => {
  return (
    <Router>
      <Routes>
        <Route path="/companies" element={<Companies />} />
      </Routes>
    </Router>
  );
};

export default AppRouter;
