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

function App() {
  return (
    <Router>
      <div className='app'>
        <Navbar />
        <Routes>
          <Route path='/' element={<Products/>}></Route>
          <Route path='/checkout' element={<CheckoutPage/>}></Route>
        </Routes>
      </div>
    </Router>
  );
}

export default App;