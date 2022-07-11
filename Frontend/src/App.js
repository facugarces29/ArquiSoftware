import "./App.css";
import { BrowserRouter as Router, Routes , Route } from "react-router-dom";
import Navbar from "./components/Navbar";
import CheckoutPage from "./pages/CheckoutPage";
import LoginPage from "./pages/LoginPage"
import HomePage from "./pages/HomePage";
import StickyFooter from "./components/StickyFooter";
import OrdersPage from "./pages/OrdersPage"
import SearchPage from "./pages/SearchPage"
import ReloadPage from "./pages/ReloadPage"

function App() {
  return (
    <Router forceRefresh={true}>
      <div className='app'>
        <Navbar/>
        <Routes>
          <Route path="/" element={<HomePage/>}></Route>
          <Route path='/checkout-page' element={<CheckoutPage/>}></Route>
          <Route path="/login" element={<LoginPage/>}></Route>
          <Route path="/orders" element={<OrdersPage/>}></Route>
          <Route path="/search/:id" element={<SearchPage/>}></Route>
          <Route path="/reload/:id" element={<ReloadPage/>}></Route>
        </Routes>
      </div>
      <StickyFooter></StickyFooter>
    </Router>
  );
}

export default App;