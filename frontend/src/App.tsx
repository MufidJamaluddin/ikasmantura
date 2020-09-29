import React from 'react';
import './App.css';
import {BrowserRouter, Route, Switch} from "react-router-dom";
import Template from "./home/Template";
import AdminApp from "./panel";


function App() {

  return (
    <BrowserRouter>
      <Switch>
        <Route path="/panel" component={AdminApp} exact={false}/>
        <Route path="/" component={Template}/>
      </Switch>
    </BrowserRouter>
  );
}

export default App;
