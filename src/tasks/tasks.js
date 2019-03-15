import React, { Component } from 'react';
import './tasks.css';
import { Helmet } from "react-helmet";
import AddTask from './addTask';
import TaskList from './taskList';

class Tasks extends Component {

    constructor(props, context) {
        super(props, context);
    }
    render() {
        const tasks = this.props.tasks;
        return (
            <div id="main">
                <AddTask onNewTaskAdded={this.props.onNewTaskAdded} />
                <TaskList tasks={tasks} onTaskCompleted={this.props.onTaskCompleted} />
            </div>
        );
    }
}

export default Tasks;