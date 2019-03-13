import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import { BrowserRouter as Router, Route, Redirect } from "react-router-dom";
import FolderItem from './folders/folderItem';
import Folders from './folders/folders';
import Tasks from './tasks/tasks';

class App extends Component {
  render() {
    return (
      <div id="container">

      <Folders/>
      <Tasks/>
      </div>
  
  );
  }
}

export default App;
