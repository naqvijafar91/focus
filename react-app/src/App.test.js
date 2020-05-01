import React from 'react';
import ReactDOM from 'react-dom';
import UserStore from './LoginPage/userStore';
import App from './App';
UserStore.saveUser({token:"sample-test"});
it('renders without crashing', () => {
  const div = document.createElement('div');
  ReactDOM.render(<App />, div);
  ReactDOM.unmountComponentAtNode(div);
});
