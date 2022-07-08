import "./App.css";
import Navbar from "./components/Navbar";
import { BrowserRouter as Router, Routes , Route } from "react-router-dom";
import CheckoutPage from "./pages/CheckoutPage";
import LoginPage from "./pages/LoginPage"
import HomePage from "./pages/HomePage";
import StickyFooter from "./components/StickyFooter";

function App() {
  return (
    <Router>
      <div className='app'>
        <Navbar/>
        <Routes>
          <Route path="/" element={<HomePage/>}></Route>
          <Route path='/checkout-page' element={<CheckoutPage/>}></Route>
          <Route path="/login" element={<LoginPage/>}></Route>
        </Routes>
      </div>
      <StickyFooter/>
    </Router>
  );
}

export default App;