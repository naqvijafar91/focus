import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import { BrowserRouter as Router, Route, Redirect } from "react-router-dom";
import FolderItem from './folders/folderItem';

class App extends Component {
  render() {
    return (
      <div id="container">

      <FolderItem/>
      
      </div>
  
  );
  }
}

export default App;
