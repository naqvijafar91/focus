import React, { Component } from 'react';
import './tasks.css';
import { Helmet } from "react-helmet";
import AddTask from './addTask';
import TaskList from './taskList';
import Logout from './logout';
import UserStore from './../LoginPage/userStore';
class Tasks extends Component {

    constructor(props, context) {
        super(props, context);
    }

    render() {
        const tasks = this.props.tasks;
        return (
            <div id="main">
                <Logout onLogout={this.props.onLogout}/>
                <AddTask onNewTaskAdded={this.props.onNewTaskAdded} />
                <TaskList tasks={tasks} 
                onTaskDueDateChanged={this.props.onTaskDueDateChanged}
                onTaskCompleted={this.props.onTaskCompleted} />
            </div>
        );
    }
}

export default Tasks;