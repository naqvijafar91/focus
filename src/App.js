import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import { BrowserRouter as Router, Route, Redirect } from "react-router-dom";
import FolderItem from './folders/folderItem';
import Folders from './folders/folders';
import Tasks from './tasks/tasks';
import Helmet from 'react-helmet';

class App extends Component {
  constructor(props, context) {
    super(props, context);

    this.state({
      folders: [
        {
          id: 1,
          name: 'Inbox',
          remaining_tasks: 13,
          tasks: [{ id: 1, task: 'Task 1 to be done' },
          { id: 2, task: 'Task 2' }, { id: 3, task: 'Task 3' },
          { id: 4, task: 'Task 4' }]
        },
        {
          id: 2,
          name: 'Grocery',
          remaining_tasks: 1234,
          tasks: [{ id: 5, task: 'Task 1 to be done g' },
          { id: 6, task: 'Task 2 g' }, { id: 7, task: 'Task 3 g' },
          { id: 8, task: 'Task 4 g' }]
        },
        {
          id: 3,
          name: 'Work',
          remaining_tasks: 2,
          tasks: [{ id: 9, task: 'Task 1 to be done w' },
          { id: 10, task: 'Task 2 w' }, { id: 11, task: 'Task 3 w' },
          { id: 12, task: 'Task 4 w' }]
        }
      ]
    });
  }
  render() {
    const currentSelectedFolderTasks = this.state.data[0].tasks;
    return (
      <div id="container">
        <Helmet>
          <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css" />
        </Helmet>
        <Folders />
        <Tasks tasks={currentSelectedFolderTasks}/>
      </div>

    );
  }
}

export default App;
