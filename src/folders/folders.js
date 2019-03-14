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

        this.enableFolderNameEditor = this.enableFolderNameEditorMode.bind(this);
        this.disableFolderEditorMode = this.disableFolderEditorMode.bind(this);
        this.handleFolderNameSubmit = this.handleFolderNameSubmit.bind(this);
    }

    enableFolderNameEditorMode(folderID) {
        // Hide that span and make input for that folder visible
        this.setState({ ['hideName' + folderID]: true, ['showInput' + folderID]: true });
    }

    disableFolderEditorMode() {
        this.state.folders.map((folder)=>{
            this.setState({ ['hideName' + folder.id]: false, ['showInput' + folder.id]: false });
        });
    }

    handleFolderNameChange(folderID, newValue) {
        //@Todo: Perform an API request to the backend
        const updatedFolders = this.state.folders.map((folder) => {
            if (folder.id == folderID)
                folder.name = newValue;
            return folder;
        });
        this.setState({ folders: updatedFolders });
    }

    handleFolderNameSubmit(event) {
        event.preventDefault();
        // Now disable edit mode 
        this.disableFolderEditorMode();
    }

    render() {
        const folders = this.state.folders.map((folder) => {
            return <li key={folder.id}>
                <div>
                    <span className={this.state['hideName' + folder.id] ? 'hidden' : 'folder-name-text'}>{folder.name}</span>
                    <form onSubmit={this.handleFolderNameSubmit} className={this.state['showInput' + folder.id]
                            ? 'edit-folder-name-input' : 'hidden'}>
                        <input onChange={(event) => this.handleFolderNameChange(folder.id, event.target.value)} value={folder.name}  type="text" />
                    </form>
                    <i onClick={() => this.enableFolderNameEditorMode(folder.id)} class="fa fa-pencil rename-folder-icon" aria-hidden="true"></i>
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