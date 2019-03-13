import React, { Component } from 'react';
import './tasks.css';
import {Helmet} from "react-helmet";
import AddTask from './addTask';

class Tasks extends Component {
    render() {
        return (
            <div id="main">
            <Helmet>
                <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css"/>
            </Helmet>
                <AddTask/>
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
            </div>
        );
    }
}

export default Tasks;