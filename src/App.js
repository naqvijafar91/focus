import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import { BrowserRouter as Router, Route, Redirect } from "react-router-dom";
import FolderItem from './folders/folderItem';
import Folders from './folders/folders';
import Tasks from './tasks/tasks';
import Helmet from 'react-helmet';

class App extends Component {
  render() {
    return (
      <div id="container">
        <Helmet>
          <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css" />
        </Helmet>
        <Folders />
        <Tasks />
      </div>

    );
  }
}

export default App;
