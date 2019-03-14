import React, { Component } from 'react';
import './folders.css';
class Folders extends Component {

    constructor(props, context) {
        super(props, context);
        this.state = {
            folders: [
                {
                    id: 1,
                    name: 'Inbox',
                    remaining_tasks: 13
                },
                {
                    id: 2,
                    name: 'Grocery',
                    remaining_tasks: 1234
                },
                {
                    id: 3,
                    name: 'Work',
                    remaining_tasks: 2
                }
            ]
        };
    }

    enableFolderNameEditor() {
        // Hide that span and make input for that folder visible
    }

    render() {
        const folders = this.state.folders.map((folder) => {
            return <li>
                <div>
                    <span>{folder.name}</span>
                    <i class="fa fa-pencil rename-folder-icon" aria-hidden="true"></i>
                    <span class="remaining-tasks">{folder.remaining_tasks}</span>
                </div>
            </li>
        });

        return (
            <div id="folders">
                <span id="add-folder">+</span>
                <ul>{folders}</ul>
            </div>
        );
    }
}

export default Folders;