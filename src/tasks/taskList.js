import React,{Component} from 'react';
import './taskList.css';

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

class TaskList extends Component {

    constructor(props,context) {
        super(props,context);

    }

    render() {
        return (
            <div id="lists">
            <ul>
                <li>
                    <div>
                        <input type="checkbox" name="vehicle" value="Bike" />
                        <span class="todos">Task one two three four</span>
                        <i class="due-date-inside-todo fa fa-calendar-o"></i>
                        <span class="due-date-text">02-04-2019</span>
                        <span class="time-left-for-task">~30m</span>
                    </div>
                </li>
                <li>
                    <div>
                        <input type="checkbox" name="vehicle" value="Bike" />
                        <span class="todos">Task one two three four</span>
                        <i class="due-date-inside-todo fa fa-calendar-o"></i>
                        <span class="due-date-text">02-04-2019</span>
                        <span class="time-left-for-task">~30m</span>
                    </div>
                </li>
                <li>
                    <div>
                        <input type="checkbox" name="vehicle" value="Bike" />
                        <span class="todos">Task one two three four</span>
                        <i class="due-date-inside-todo fa fa-calendar-o"></i>
                        <span class="due-date-text">02-04-2019</span>
                        <span class="time-left-for-task">~30m</span>
                    </div>
                </li>
                <li>
                    <div>
                        <input type="checkbox" name="vehicle" value="Bike" />
                        <span class="todos">Task one two three four</span>
                        <i class="due-date-inside-todo fa fa-calendar-o"></i>
                        <span class="due-date-text">02-04-2019</span>
                        <span class="time-left-for-task">~30m</span>
                    </div>
                </li>
                <li>
                    <div>
                        <input type="checkbox" name="vehicle" value="Bike" />
                        <span class="todos">Task one two three four</span>
                        <i class="due-date-inside-todo fa fa-calendar-o"></i>
                        <span class="due-date-text">02-04-2019</span>
                        <span class="time-left-for-task">~30m</span>
                    </div>
                </li>


            </ul>
        </div>
        );
    }
}

export default TaskList;