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

        this.enableFolderNameEditor = this.enableFolderNameEditor.bind(this);
    }

    enableFolderNameEditor(folderID) {
        // Hide that span and make input for that folder visible
        console.log(folderID);
        this.setState({['hideName'+folderID]:true,['showInput'+folderID]:true});
    }

    render() {
        const folders = this.state.folders.map((folder) => {
            return <li key={folder.id}>
                <div>
                    <span className={this.state['hideName'+folder.id]?'hidden':'folder-name-text'}>{folder.name}</span>
                    <input value={folder.name} className={this.state['showInput'+folder.id]?'edit-folder-name-input':'hidden'} type="text"/>
                    <i onClick={()=>this.enableFolderNameEditor(folder.id)}class="fa fa-pencil rename-folder-icon" aria-hidden="true"></i>
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