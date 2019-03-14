import React, { Component } from 'react';
import './tasks.css';
import { Helmet } from "react-helmet";
import AddTask from './addTask';
import TaskList from './taskList';

class Tasks extends Component {

    constructor(props, context) {
        super(props, context);
        this.state = {
            tasks: [{ id: 1, task: 'Task 1 to be done' },
            { id: 2, task: 'Task 2' }, { id: 3, task: 'Task 3' },
            { id: 4, task: 'Task 4' }]
        }
        this.onNewTaskAdded = this.onNewTaskAdded.bind(this);
        this.onTaskCompleted = this.onTaskCompleted.bind(this);
    }

    onTaskCompleted(taskID) {
        // Loop through the tasks array and remove task with this id
        const updatedTasks = this.state.tasks.filter((taskItem) => {
            if (taskItem.id == taskID)
                return false;
            return true;
        });
        this.setState({ tasks: updatedTasks });
    }

    onNewTaskAdded(taskToBeAdded) {
        console.log(taskToBeAdded);
        this.setState({
            tasks: [...this.state.tasks, { id: 100, task: taskToBeAdded }]
        });

    }
    render() {
        const tasks = this.state.tasks;
        return (
            <div id="main">
                <AddTask onNewTaskAdded={this.onNewTaskAdded} />
                <TaskList tasks={tasks} onTaskCompleted={this.onTaskCompleted} />
            </div>
        );
    }
}

export default Tasks;