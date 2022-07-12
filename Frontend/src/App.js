import './App.css';
import { BrowserRouter as Router, Routes , Route } from 'react-router-dom';
import Navbar from './components/Navbar';
import CheckoutPage from './pages/CheckoutPage';
import LoginPage from './pages/LoginPage'
import HomePage from './pages/HomePage';
import StickyFooter from './components/StickyFooter';
import OrdersPage from './pages/OrdersPage'
import SearchPage from './pages/SearchPage'
import ReloadPage from './pages/ReloadPage'
import PlaceOrderPage from './pages/PlaceOrderPage'

function App() {
  return (
    <Router forceRefresh={true}>
      <div className='app'>
        <Navbar />
        <Routes>
          <Route path='/' element={<HomePage />} />
          <Route path='/checkout' element={<CheckoutPage />} />
          <Route path='/login' element={<LoginPage />} />
          <Route path='/orders' element={<OrdersPage />} />
          <Route path='/search/:id' element={<SearchPage/>} />
          <Route path='/reload/:id' element={<ReloadPage/>} />
          <Route path='/place-order' element={<PlaceOrderPage />} />
        </Routes>
      </div>
      <StickyFooter />
    </Router>
  );
}

export default App;