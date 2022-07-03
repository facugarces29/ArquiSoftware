import React from 'react';
import './App.css';
import HomeProducts from './components/HomeProducts';
//import Product from './components/Product';
import Navbar from './components/Navbar';
//import HomeProducts from './pages/HomeProducts';
//import Products from './components/Products';
//import CheckoutPage from './pages/CheckoutPage';

function App() {

  return (
    <div classNameName='app'>
      <Navbar></Navbar>   
      <HomeProducts></HomeProducts>
    </div>
    
  );
}

export default App;
