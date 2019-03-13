import React,{Component} from 'react';
import './addTask.css';

class AddTask extends Component {
    render() {
        return(
            <div id="pusher">
            <form id="pusher-form">
                <input class="pusher-input" placeholder="What do you want to do?" type="text" name="lname" />
                <i class="due-date fa fa-calendar-o"></i>
                <br />
            </form>
        </div>
        );
    }
}

export default AddTask;