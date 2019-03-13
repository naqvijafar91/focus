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
        this.onNewTaskAdded = this.onNewTaskAdded.bind(this);
    }

    onNewTaskAdded(taskToBeAdded) {
        console.log(taskToBeAdded);
        this.setState({
            tasks : [...this.state.tasks,taskToBeAdded]
        });
       
    }
    render() {
        const tasks = this.state.tasks;
        return (
            <div id="main">
            <Helmet>
                <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css"/>
            </Helmet>
                <AddTask onNewTaskAdded={this.onNewTaskAdded}/>
                <TaskList tasks={tasks}/>
            </div>
        );
    }
}

export default Tasks;