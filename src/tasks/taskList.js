import React, { Component } from 'react';
import './taskList.css';

class TaskList extends Component {

    constructor(props, context) {
        super(props, context);
        this.handleTaskCompleted = this.handleTaskCompleted.bind(this);
    }
    handleTaskCompleted(event) {
        if(event.target.checked) {
            // @todo:Hit api and remove it from UI
            this.props.onTaskCompleted(event.target.name);
        }
    }
    render() {
        var listItems = this.props.tasks.map((taskItem) =>
            <li key={taskItem.id}>
                <div>
                    <input type="checkbox" name={taskItem.id}
                         onChange={this.handleTaskCompleted}/>
                    <div className="todos">{taskItem.task}</div>
                    <i className="due-date-inside-todo fa fa-calendar"></i>
                    <span className="due-date-text">02-04-2019</span>
                    <span className="time-left-for-task">~30m</span>
                </div>
            </li>);

        return (
            <div id="lists">
                <ul>{listItems}</ul>
            </div>
        );
    }
}

//@todo: Use this component instead 
class TaskItem extends Component {
    constructor(props, context) {
        super(props, context);
    }
    render() {
        return (
            <li>
                <div>
                    <input type="checkbox" name="vehicle" value="Bike" />
                    <span class="todos">Task one two three four</span>
                    <i class="due-date-inside-todo fa fa-calendar"></i>
                    <span class="due-date-text">02-04-2019</span>
                    <span class="time-left-for-task">~30m</span>
                </div>
            </li>
        );
    }
}

export default TaskList;