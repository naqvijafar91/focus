import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Redirect } from "react-router-dom";
import LoginPage from './LoginPage/LoginPage';
import App from './App';
import './main.css';
import UserStore from './LoginPage/userStore';

class Main extends Component {
    constructor(props,context) {
        super(props,context);
    }

    render() {
        const isLoggedIn = false;
        return(
        <Router>
            <div id="container">
                <PrivateRoute path="/" exact component={App} />
                <Route path="/signin" component={LoginPage} />
            </div>
        </Router>
        );
    }
}

const PrivateRoute = ({ component: Component, ...rest }) => (
    
    <Route
      {...rest}
      render={props =>
        UserStore.isAuthenticated? (
          <Component {...props} />
        ) : (
          <Redirect
            to={{
              pathname: "/signin",
              state: { from: props.location }
            }}
          />
        )
      }
    />
  );

export default Main;