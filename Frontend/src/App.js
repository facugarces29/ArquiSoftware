//import { useEffect } from "react";
import "./App.css";
import Navbar from "./components/Navbar";
import { BrowserRouter as Router, Routes , Route } from "react-router-dom";
//import SignIn from "./components/Signin";
//import SignUp from "./components/Signup";
//import { auth } from "./firebase";
//import { useStateValue } from "./StateProvider";
//import { actionTypes } from "./reducer";
//import Checkout from "./components/ProcessOrder/Checkout";
//import Products from "./components/Products";
import CheckoutPage from "./pages/CheckoutPage";
import Products from "./components/Products"
import Login from "./pages/Login"
import HomeProducts from "./components/HomeProducts";

function App() {
  return (
    <Router>
      <div className='app'>
        <Navbar />
        <Routes>
          {/*<Route path='/' element={<Products/>}></Route>*/}
          <Route path="/" element={<HomeProducts/>}></Route>
          <Route path='/checkout-page' element={<CheckoutPage/>}></Route>
          <Route path="/login" element={<Login/>}></Route>
        </Routes>
      </div>
    </Router>
  );
}

export default App;