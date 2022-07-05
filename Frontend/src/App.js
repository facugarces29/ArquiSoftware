import React from 'react';
import './App.css';
//import HomeProducts from './components/HomeProducts';
//import Product from './components/Product';
import Navbar from './components/Navbar';
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
