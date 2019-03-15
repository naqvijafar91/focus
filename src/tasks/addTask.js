import React,{Component} from 'react';
import './addTask.css';

class AddTask extends Component {

    constructor(props,context) {
        super(props,context);
        this.state = {taskToBeAdded: ''};
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }
    handleChange(event) {
        this.setState({taskToBeAdded: event.target.value});
      }
    
    handleSubmit(event) {
        event.preventDefault();
        this.props.onNewTaskAdded(this.state.taskToBeAdded);
        this.setState({taskToBeAdded:''});
      }
    render() {
        return(
            <div id="pusher">
            <form id="pusher-form" onSubmit={this.handleSubmit}>
                <input className="pusher-input" value={this.state.taskToBeAdded} onChange={this.handleChange}
                placeholder="What do you want to do?" type="text" name="lname" />
                <i className="due-date fa fa-calendar-o"></i>
                <br />
            </form>
        </div>
        );
    }
}

export default AddTask;