import React from "react";
import {Route,Switch} from "react-router-dom";
import Home from "./pages/Home";
import Profile from "./components/Home/Profile";
import Login from "./pages/Login";
import Cart from "./pages/Cart";
import Payment from "./pages/Payment";
import Register from "./pages/Register";
import PaymentResponse from "./pages/PaymentResponse";
import Products from "./pages/Products";
import Dashboard from "./pages/Dashboard";

import ServicePage from "./pages/ServicePage";
import ContactPage from "./pages/ContactPage";

function App() {
  return (
      <Switch>
        <Route exact path='/'>
            <Login/>
        </Route>
          <Route path='/login'>
                <Login/>
          </Route>
          <Route path='/Register'>
              <Register/>
          </Route>
          <Route path='/profile'>
              <Profile/>
          </Route>
          
        
          <Route path='/service'>
              <ServicePage/>
          </Route>
          <Route path='/contact'>
              <ContactPage/>
          </Route>
          <Route path='/cart'>
              <Cart/>
          </Route>
          <Route path='/payment'>
              <Payment/>
          </Route>
          <Route path='/products'>
              <Products/>
          </Route>
          <Route path='/response'>
              <PaymentResponse/>
          </Route>
          <Route path='/home'>
              <Home/>
          </Route>
          <Route path='/dashboard'>
              <Dashboard/>
          </Route>
          
      </Switch>
  );
}

export default App;
