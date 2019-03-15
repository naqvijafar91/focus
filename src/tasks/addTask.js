import React, { Component } from 'react';
import Calendar from 'react-calendar';
import './addTask.css';

class AddTask extends Component {

    constructor(props, context) {
        super(props, context);
        this.state = {
            taskToBeAdded: '',
            date: new Date(),
            showCalendar : false
        };
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.onDueDateChanged = this.onDueDateChanged.bind(this);
        this.showCalendar = this.showCalendar.bind(this);
        this.hideCalendar = this.hideCalendar.bind(this);

        this.setWrapperRef = this.setWrapperRef.bind(this);
        this.handleClickOutside = this.handleClickOutside.bind(this);
    }
    handleChange(event) {
        this.setState({ taskToBeAdded: event.target.value });
    }

    handleSubmit(event) {
        event.preventDefault();
        this.props.onNewTaskAdded(this.state.taskToBeAdded);
        this.setState({ taskToBeAdded: '' });
    }

    showCalendar() {
        this.setState({showCalendar:true});
    }

    hideCalendar() {
        this.setState({showCalendar:false});
    }

    toggleCalendarVisibility() {
        this.state.showCalendar ? this.hideCalendar() : this.showCalendar();
    }

    onDueDateChanged(date) {
        this.setState({ date });
        this.hideCalendar();
    }

    setWrapperRef(node) {
        this.wrapperRef = node;
    }

    handleClickOutside(event) {
        if (this.wrapperRef && !this.wrapperRef.contains(event.target)) {
          this.hideCalendar();
        }
    }

    componentDidMount() {
        document.addEventListener('mousedown', this.handleClickOutside);
    }
    
    componentWillUnmount() {
        document.removeEventListener('mousedown', this.handleClickOutside);
    }

    render() {
        return (
            <div id="pusher">
                <form ref={this.setWrapperRef} id="pusher-form" onSubmit={this.handleSubmit}>
                    <input className="pusher-input" value={this.state.taskToBeAdded} onChange={this.handleChange}
                        placeholder="What do you want to do?" type="text" name="lname" />
                    <i onClick={()=>this.toggleCalendarVisibility()} className="due-date fa fa-calendar-o"></i>
                    <Calendar className={this.state.showCalendar?'add-task-calendar':'hidden'}
                     onChange={this.onDueDateChanged}
                        value={this.state.date} onBlur={this.hideCalendar}/>
                    <br />
                </form>
            </div>
        );
    }
}

export default AddTask;