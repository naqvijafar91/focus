import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import { BrowserRouter as Router, Route, Redirect } from "react-router-dom";
import FolderItem from './folders/folderItem';
import Folders from './folders/folders';
import Tasks from './tasks/tasks';
import Helmet from 'react-helmet';
import UserStore from './LoginPage/userStore';
import axios from 'axios';
import ServerURLFetcher from './ServerURLFetcher';

class App extends Component {
  constructor(props, context) {
    super(props, context);
    this.baseURL = ServerURLFetcher.getURL();
    this.state = {
      data: [],
      currentFolderIndexSelected: 0
    };
    // {
    //   id: 1,
    //   name: 'Inbox',
    //   remaining_tasks: 13,
    //   tasks: [{ id: 1, description: 'Task 1 to be done', dueDate: '' },
    //   { id: 2, description: 'Task 2', dueDate: new Date() },
    //   { id: 3, description: 'Task 3', dueDate: new Date() },
    //   { id: 4, description: 'Task 4', dueDate: '' }]
    // },
    // {

    //   id: 2,
    //   name: 'Grocery',
    //   remaining_tasks: 1234,
    //   tasks: [{ id: 5, description: 'Task 1 to be done g' },
    //   { id: 6, description: 'Task 2 g' }, { id: 7, description: 'Task 3 g' },
    //   { id: 8, description: 'Task 4 g' }]
    // },
    // {
    //   id: 3,
    //   name: 'Work',
    //   remaining_tasks: 2,
    //   tasks: [{ id: 9, description: 'Task 1 to be done w' },
    //   { id: 10, description: 'Task 2 w' }, { id: 11, description: 'Task 3 w' },
    //   { id: 12, description: 'Task 4 w' }]
    // }
    this.onTaskCompleted = this.onTaskCompleted.bind(this);
    this.onTaskDueDateChanged = this.onTaskDueDateChanged.bind(this);
    this.onNewTaskAdded = this.onNewTaskAdded.bind(this);
    this.onNewFolderSelected = this.onNewFolderSelected.bind(this);
    this.handleFolderNameChange = this.handleFolderNameChange.bind(this);
    this.updateFolderName = this.updateFolderName.bind(this);
    this.addDummyFolderItem = this.addDummyFolderItem.bind(this);
    this.onLogout = this.onLogout.bind(this);
    this.fetchLatestDataFromServer = this.fetchLatestDataFromServer.bind(this);
    this.parseCompleteServerResponse = this.parseCompleteServerResponse.bind(this);
    this.updateTask = this.updateTask.bind(this);
    this.parseDate = this.parseDate.bind(this);
    this.extractDateString = this.extractDateString.bind(this);
    this.fetchLatestDataFromServer();
  }

  /**
   * Parses server response by converting string to dates
   * @param {JSONObject} response 
   * @returns {JSONObject} parsedResponse
   */
  parseCompleteServerResponse(response) {
    let self = this;
    let folders = response.data;
    if (!folders) {
      folders = [];
    }
    for (let i = 0; i < folders.length; i++) {
      let folder = folders[i];
      for (let j = 0; j < folder.tasks.length; j++) {
        let task = folder.tasks[j];
        // parse due_date and completed_date
        task.due_date = self.parseDate(task.due_date);
        task.completed_date = self.parseDate(task.completed_date);
      }
    }
    return response;
  }

  parseDate(dateString) {
    if(!dateString) {
      return;
    }
    let parts = dateString.split("-");
    return new Date(parseInt(parts[2], 10),
      parseInt(parts[1], 10) - 1,
      parseInt(parts[0], 10));
  }

  /**
   * Converts date object to its string representation
   * @param {Date} dateObj 
   */
  extractDateString(dateObj) {
    try {
      let day = dateObj.getDate();
      if(day<10) {
        day = "0"+day;
      }
      let month = dateObj.getMonth()+1;
      if(month<10) {
        month = "0"+month;
      }
      const year = dateObj.getFullYear();
      const dateStr = day+"-"+month+"-"+year;
      return dateStr;
    } catch(err) {
      return null;
    }
    return null;
  }

  fetchLatestDataFromServer() {
    let self = this;
    axios({
      method: 'get',
      url: 'http://localhost:8080',
      headers: {
        'Authorization': 'Bearer ' + UserStore.getUser().token,
        'ss': 'sssss'
      }
    }).then(function (response) {
      console.log(response.data);
      self.setState({ data: self.parseCompleteServerResponse(response.data).data });
    }).catch(function (err) {
      console.log(err);
      self.setState({ data: [] });
      alert(err);
    })
  }

  handleFolderNameChange(folderID, newValue) {
    //Question: Should we perform an API request to the backend?
    const updatedData = this.state.data.map((folder) => {
      if (folder.id == folderID)
        folder.name = newValue;
      return folder;
    });
    this.setState({ data: updatedData });
  }

  //@Todo:Perform an api hit
  updateFolderName(folderID) {
    console.log('Performing API HIT to update folder name .... with ID ' + folderID);
    let updatedFolder = null;
    for (var i = 0; i < this.state.data.length; i++) {
      const folder = this.state.data[i];
      if (folder.id == folderID) {
        updatedFolder = folder;
      }
    }
    axios({
      method: 'put',
      url: this.baseURL + '/folder',
      headers: {
        'Authorization': 'Bearer ' + UserStore.getUser().token
      },
      data: updatedFolder
    }).then(function (resp) {
      console.log('Folder updated');
    }).catch(function (err) {
      alert(err);
    });
  }

  onNewFolderSelected(folderID) {
    for (var i = 0; i < this.state.data.length; i++) {
      const folder = this.state.data[i];
      if (folder.id == folderID) {
        this.setState({ currentFolderIndexSelected: i });
        return;
      }
    }
  }

  onTaskCompleted(taskID) {
    let taskToBeUpdated = null;
    let self = this;
    // Loop through the tasks array and fremove task with this id
    const updatedTasksForCurrentSelectedFolder = this.state.data[this.state.currentFolderIndexSelected].tasks.filter((taskItem) => {
      if (taskItem.id == taskID) {
        taskItem.completed_date = new Date();
        taskToBeUpdated = taskItem;
        return false;
      }
      return true;
    });

    this.updateTask(taskToBeUpdated)
      .then(function (done) {
        //Update our state
        const newState = Object.assign({}, self.state);
        // Subtract remaining tasks
        newState.data[newState.currentFolderIndexSelected].remaining_tasks--;
        newState.data[newState.currentFolderIndexSelected].tasks = updatedTasksForCurrentSelectedFolder;
        self.setState(newState);
      }).catch(function (err) {
        alert(err);
      });
  }

  onTaskDueDateChanged(taskID, dueDate) {
    let taskToBeUpdated = null;
    let self = this;
    // Loop through the tasks array and fetch and update task with this id
    // Also fetch the task object for hitting the API
    const updatedTasksForCurrentSelectedFolder = this.state.data[this.state.currentFolderIndexSelected].tasks.map((taskItem) => {
      if (taskItem.id == taskID) {
        taskItem.due_date = dueDate;
        taskToBeUpdated = taskItem;
      }
      return taskItem;
    });
    console.log(updatedTasksForCurrentSelectedFolder);
    this.updateTask(taskToBeUpdated)
      .then(function (done) {
        //Update our state
        const newState = Object.assign({}, self.state);
        newState.data[newState.currentFolderIndexSelected].tasks = updatedTasksForCurrentSelectedFolder;
        self.setState(newState);
      }).catch(function (error) {
        alert(error);
      });
  }

  /**
   * Updates a task object on the server, this function also parses the date strings into correct string format
   * representations
   * @param {Task} updatedTask 
   */
  updateTask(updatedTask) {
    // First modify 
    let clonedTask = {
      id : updatedTask.id,
      description : updatedTask.description,
      folder_id: updatedTask.folder_id,
      due_date : this.extractDateString(updatedTask.due_date),
      completed_date: this.extractDateString(updatedTask.completed_date)
    };
    return axios({
      method: "put",
      url: this.baseURL + "/task",
      headers: {
        'Authorization': 'Bearer ' + UserStore.getUser().token
      },
      data: clonedTask
    });
  }

  onNewTaskAdded(taskToBeAdded, dueDate) {
    let self = this;
    console.log(taskToBeAdded + ' From App.js adding newTask');
    axios({
      url: this.baseURL + '/task',
      method: "post",
      headers: {
        'Authorization': 'Bearer ' + UserStore.getUser().token
      },
      data: {
        "description": taskToBeAdded,
        "folder_id": self.state.data[self.state.currentFolderIndexSelected].id,
        // @Todo: Uncomment this once backend is good
        "due_date" : self.extractDateString(dueDate)
      }
    }).then(function (response) {
      //Update our state
      const newState = Object.assign({}, self.state);
      // Increment remaining tasks
      newState.data[newState.currentFolderIndexSelected].remaining_tasks++;
      newState.data[newState.currentFolderIndexSelected].tasks = [...newState.data[newState.currentFolderIndexSelected].tasks, { id: response.data.id, description: taskToBeAdded, due_date: dueDate }];
      self.setState(newState);
    }).catch(function (err) {
      alert(err);
    });
  }

  /**
   * This function is used as a 1st step to add a new folder
   * @Todo : Perform API hit to create a dummy folder named New Folder
   */
  addDummyFolderItem(notifyChildComponentWithNewID) {
    let self = this;
    axios({
      method: 'post',
      url: this.baseURL + '/folder',
      headers: {
        'Authorization': 'Bearer ' + UserStore.getUser().token
      },
      data: {
        "name": "New Folder"
      }
    }).then(function (response) {
      console.log(response);
      const newState = Object.assign({}, self.state);
      if(!newState.data) {
        newState.data = [];
      }
      newState.data.push({
        id: response.data.id,
        name: 'New Folder',
        remaining_tasks: 0,
        tasks: []
      });
      self.setState(newState, function () {
        notifyChildComponentWithNewID(response.data.id);
      });
    }).catch(function (err) {
      alert(err);
    });

  }

  onLogout() {
    UserStore.deleteUser();
    this.props.history.push('/signin');
  }

  render() {
    let data = this.state.data;
    if (!data) {
      data = [];
    }
    const currentSelectedFolderObject = data[this.state.currentFolderIndexSelected];
    let currentSelectedFolderTasks = [];
    let currentSelectedFolderID = "";
    if (currentSelectedFolderObject) {
      currentSelectedFolderTasks = currentSelectedFolderObject.tasks;
      currentSelectedFolderID = currentSelectedFolderObject.id;
    }
    return (
      <div id="app-container">
        <Helmet>
          <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css" />
        </Helmet>
        <Folders data={data}
          currentSelectedFolderID={currentSelectedFolderID}
          addDummyFolderItem={this.addDummyFolderItem}
          updateFolderName={this.updateFolderName}
          handleFolderNameChange={this.handleFolderNameChange}
          onNewFolderSelected={this.onNewFolderSelected} />
        <Tasks tasks={currentSelectedFolderTasks}
          onLogout={this.onLogout}
          onTaskDueDateChanged={this.onTaskDueDateChanged}
          onNewTaskAdded={this.onNewTaskAdded}
          onTaskCompleted={this.onTaskCompleted} />
      </div>

    );
  }
}

export default App;
