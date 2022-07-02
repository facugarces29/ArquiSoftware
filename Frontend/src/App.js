import React from 'react';
import './App.css';
//import Product from './components/Product';
import Navbar from './components/Navbar';
import HomeProducts from './pages/HomeProducts';

function App() {

  return (
    <div classNameName='app'>
      <Navbar></Navbar>
      <HomeProducts></HomeProducts>    
    </div>
    
  );
}

export default App;
