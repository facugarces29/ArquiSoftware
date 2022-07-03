import React from 'react';
import './App.css';
//import Product from './components/Product';
import Navbar from './components/Navbar';
//import HomeProducts from './pages/HomeProducts';
//import Products from './components/Products';
import CheckoutPage from './pages/CheckoutPage';

function App() {

  return (
    <div classNameName='app'>
      <Navbar></Navbar>   
      <CheckoutPage></CheckoutPage>
    </div>
    
  );
}

export default App;
