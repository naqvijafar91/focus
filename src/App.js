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

    this.state = {
      data: [
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
      ],
      currentFolderIndexSelected: 0
    };

    this.onTaskCompleted = this.onTaskCompleted.bind(this);
    this.onNewTaskAdded = this.onNewTaskAdded.bind(this);
    this.onNewFolderSelected = this.onNewFolderSelected.bind(this);
    this.handleFolderNameChange = this.handleFolderNameChange.bind(this);
    this.updateFolderName = this.updateFolderName.bind(this);
  }

  handleFolderNameChange(folderID, newValue) {
    //@Todo: Perform an API request to the backend
    const updatedData = this.state.data.map((folder) => {
        if (folder.id == folderID)
            folder.name = newValue;
        return folder;
    });
    this.setState({ data: updatedData });
  }

  //@Todo:Perform an api hit
  updateFolderName(folderID) {
    console.log('Performing API HIT to update folder name .... with ID '+folderID);
  }

  onNewFolderSelected(folderID) {
    for(var i=0;i<this.state.data.length;i++) {
      const folder = this.state.data[i];
      if(folder.id == folderID) {
        this.setState({currentFolderIndexSelected : i});
        return;
      }
    }
  }

  onTaskCompleted(taskID) {
    // Loop through the tasks array and remove task with this id
    const updatedTasksForCurrentSelectedFolder = this.state.data[this.state.currentFolderIndexSelected].tasks.filter((taskItem) => {
      if (taskItem.id == taskID)
        return false;
      return true;
    });

    //Update our state
    const newState = Object.assign({}, this.state);
    newState.data[newState.currentFolderIndexSelected].tasks = updatedTasksForCurrentSelectedFolder;
    this.setState(newState);
  }

  onNewTaskAdded(taskToBeAdded) {
    console.log(taskToBeAdded + ' From App.js adding newTask');
    //Update our state
    const newState = Object.assign({}, this.state);
    newState.data[newState.currentFolderIndexSelected].tasks = [...newState.data[newState.currentFolderIndexSelected].tasks, { id: 100, task: taskToBeAdded }];
    this.setState(newState);
  }
  render() {
    const currentSelectedFolderTasks = this.state.data[this.state.currentFolderIndexSelected].tasks;
    return (
      <div id="container">
        <Helmet>
          <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css" />
        </Helmet>
        <Folders data={this.state.data}
        updateFolderName={this.updateFolderName}
        handleFolderNameChange={this.handleFolderNameChange}
         onNewFolderSelected={this.onNewFolderSelected} />
        <Tasks tasks={currentSelectedFolderTasks}
          onNewTaskAdded={this.onNewTaskAdded}
          onTaskCompleted={this.onTaskCompleted} />
      </div>

    );
  }
}

export default App;
