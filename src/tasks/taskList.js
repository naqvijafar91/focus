import React,{Component} from 'react';
import './taskList.css';

class TaskList extends Component {

    constructor(props,context) {
        super(props,context);
        this.state = {
            listItems : ['']
        }
        this.updateTaskItems = this.updateTaskItems.bind(this);
    }

    updateTaskItems() {
        var listItems = this.props.tasks.map((taskItem) =>
        <li key={taskItem.toString()}>
          <div>
                        <input type="checkbox" name="vehicle" value="Bike" />
                        <span class="todos">{taskItem}</span>
                        <i class="due-date-inside-todo fa fa-calendar-o"></i>
                        <span class="due-date-text">02-04-2019</span>
                        <span class="time-left-for-task">~30m</span>
            </div>
        </li>);
        this.setState({listItems : listItems},()=>console.log('List State updated with'+this.state.listItems));
    }

    componentDidUpdate(prevProps) {
        if(this.props.tasks!=prevProps.tasks)  {
            this.updateTaskItems();
        }
       
    } 
    componentDidMount() {
        this.updateTaskItems();
    }
    render() {
        return (
            <div id="lists">
            <ul>{this.state.listItems}</ul>
        </div>
        );
    }
}

//@todo: Use this component instead 
class TaskItem extends Component {
    constructor(props,context) {
        super(props,context);
    }
    render() {
        return (
            <li>
            <div>
                <input type="checkbox" name="vehicle" value="Bike" />
                <span class="todos">Task one two three four</span>
                <i class="due-date-inside-todo fa fa-calendar-o"></i>
                <span class="due-date-text">02-04-2019</span>
                <span class="time-left-for-task">~30m</span>
            </div>
        </li>
        );
    }
}

export default TaskList;