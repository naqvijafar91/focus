import React, { Component } from 'react';
import './folders.css';
class Folders extends Component {

    render() {
        const editIconStyle = {
            fontSize: '0.75em'
          };
        return (
        <div id="folders">
            <span id="add-folder">+</span>
            <ul>
                <li>
                    <div>
                        <span>Inbox</span>
                        <i class="fa fa-pencil rename-folder-icon" aria-hidden="true"
                        style={editIconStyle}></i>
                        <span class="remaining-tasks">13</span>
                    </div>
                </li>
                <li>
                    <div>
                        <span>Work</span>
                        <i class="fa fa-pencil rename-folder-icon" aria-hidden="true"
                        style={editIconStyle}></i>
                        <span class="remaining-tasks">103</span>
                    </div>
                </li>
                <li>
                    <div>
                        <span>Grocery</span>
                        <i class="fa fa-pencil rename-folder-icon" aria-hidden="true"
                        style={editIconStyle}></i>
                        <span class="remaining-tasks">3</span>
                    </div>
                </li>
            </ul>
        </div>
        );
    }
}

export default Folders;