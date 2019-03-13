import React, { Component } from 'react';
import './folders.css';
class Folders extends Component {

    render() {
        return (
        <div id="folders">
            <span id="add-folder">+</span>
            <ul>
                <li>
                    <div>
                        <span>Inbox</span>
                        <span class="remaining-tasks">13</span>
                    </div>
                </li>
                <li>
                    <div>
                        <span>Work</span>
                        <span class="remaining-tasks">100</span>
                    </div>
                </li>
                <li>
                    <div>
                        <span>Grocery</span>
                        <span class="remaining-tasks">10</span>
                    </div>
                </li>
            </ul>
        </div>
        );
    }
}

export default Folders;