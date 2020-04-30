import React, { Component } from 'react';
import './taskList.css';
import Calendar from 'react-calendar';

class TaskList extends Component {

    constructor(props, context) {
        super(props, context);
        this.handleTaskCompleted = this.handleTaskCompleted.bind(this);
        this.onDueDateChanged = this.onDueDateChanged.bind(this);
        this.showCalendar = this.showCalendar.bind(this);
        this.hideCalendar = this.hideCalendar.bind(this);
        this.state = {
            showCalendarWithTaskID : ''
        };

    }
    handleTaskCompleted(event) {
        if (event.target.checked) {
            // @todo:Hit api and remove it from UI
            this.props.onTaskCompleted(event.target.name);
        }
    }

    showCalendar(taskID) {
        this.setState({showCalendarWithTaskID:taskID});
    }

    hideCalendar(taskID) {
        this.setState({showCalendarWithTaskID:''});
    }

    toggleCalendarVisibility(taskID) {
        this.state.showCalendarWithTaskID == taskID ? this.hideCalendar(taskID) : this.showCalendar(taskID);
    }

    onDueDateChanged(taskID,date) {
        // Notify the parent to update due date with that taskID
        this.hideCalendar();
        this.props.onTaskDueDateChanged(taskID,date);
    }
    
    render() {
        var listItems = this.props.tasks.map((taskItem) => {
            const displayDueDate = !taskItem.dueDate || taskItem.dueDate == '' ? false : true; 
            const date = taskItem.dueDate;
            const formattedDueDate = displayDueDate ? taskItem.dueDate.toDateString() : '';
            return <li key={taskItem.id}>
                <div>
                    <input type="checkbox" name={taskItem.id}
                        onChange={this.handleTaskCompleted} />
                    <div className="todos">{taskItem.description}</div>
                    <Calendar className={this.state.showCalendarWithTaskID == taskItem.id ?'add-task-calendar':'hidden'}
                     onChange={(updatedDate)=>this.onDueDateChanged(taskItem.id,updatedDate)}
                        value={date==''?new Date():date}/>
                    <i onClick={()=>this.toggleCalendarVisibility(taskItem.id)} className={displayDueDate?"due-date-inside-todo fa fa-calendar":'hidden'}></i>
                    <i onClick={()=>this.toggleCalendarVisibility(taskItem.id)} className={displayDueDate?'hidden':"due-date-inside-todo fa fa-calendar-o"}></i>
                    <span className="due-date-text">{formattedDueDate}</span>
                    {/* <span className="time-left-for-task">~30m</span> */}
                </div>
            </li>
        });

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