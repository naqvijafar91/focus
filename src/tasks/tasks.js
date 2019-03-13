import React, { Component } from 'react';
import './tasks.css';
import {Helmet} from "react-helmet";
import AddTask from './addTask';
import TaskList from './taskList';

class Tasks extends Component {

    constructor(props,context) {
        super(props,context);
        this.state={
            tasks : ['Task 1','Task 2','Task 3','Task 4']
        }
    }

    onNewTaskAdded(taskToBeAdded) {
        console.log(taskToBeAdded);
    }
    render() {
        return (
            <div id="main">
            <Helmet>
                <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css"/>
            </Helmet>
                <AddTask onNewTaskAdded={this.onNewTaskAdded}/>
                <TaskList/>
            </div>
        );
    }
}

export default Tasks;