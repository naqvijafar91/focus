import React, { Component } from 'react';
import './folders.css';
class Folders extends Component {

    constructor(props, context) {
        super(props, context);
        this.state={};
        this.enableFolderNameEditorMode = this.enableFolderNameEditorMode.bind(this);
        this.disableFolderEditorMode = this.disableFolderEditorMode.bind(this);       
        this.handleFolderNameSubmit = this.handleFolderNameSubmit.bind(this);
        this.addNewFolder = this.addNewFolder.bind(this);
        this.showEditButtonForFolder = this.showEditButtonForFolder.bind(this);
        this.hideEditButtonForFolder = this.hideEditButtonForFolder.bind(this);
    }

    enableFolderNameEditorMode(folderID) {
        // Hide that span and make input for that folder visible
        this.setState({ ['hideName' + folderID]: true, ['showInput' + folderID]: true });
    }

    disableFolderEditorMode() {
        this.props.data.map((folder)=>{
            this.setState({['hideName' + folder.id]: false, ['showInput' + folder.id]: false });
        });
    }

    handleFolderNameSubmit(event,folderID) {
        event.preventDefault();
        // Now disable edit mode 
        this.disableFolderEditorMode();
        this.props.updateFolderName(folderID);
    }

    addNewFolder() {
        let self = this;
        // Notify parent to add a dummy folder item and 
        // enable editor mode for that id
        this.props.addDummyFolderItem(function(newFolderID){
            self.enableFolderNameEditorMode(newFolderID);
        });
     }

    showEditButtonForFolder(folderID) {
        this.setState({ ['showEditButton' + folderID]: true});
    }

    hideEditButtonForFolder(folderID) {
        this.setState({ ['showEditButton' + folderID]: false});
    }

    render() {
        const folders = this.props.data.map((folder) => {
            return <li key={folder.id} 
            className={this.props.currentSelectedFolderID == folder.id?'highlight-selected-folder':''}
            onMouseEnter={()=>this.showEditButtonForFolder(folder.id)}
            onMouseLeave={()=>this.hideEditButtonForFolder(folder.id)}
            onClick={()=>this.props.onNewFolderSelected(folder.id)}>
                <div>
                    <span className={this.state['hideName' + folder.id] ? 'hidden' : 'folder-name-text'}>{folder.name}</span>
                    <form onSubmit={(event)=>this.handleFolderNameSubmit(event,folder.id)} className={this.state['showInput' + folder.id]
                            ? 'edit-folder-name-input' : 'hidden'}>
                        <input onChange={(event) => this.props.handleFolderNameChange(folder.id, event.target.value)} value={folder.name}  type="text" />
                    </form>
                    <i onClick={() => this.enableFolderNameEditorMode(folder.id)} 
                        className={this.state['showEditButton'+folder.id]?"fa fa-pencil rename-folder-icon":'hidden'} aria-hidden="true"></i>
                    <span className="remaining-tasks">{folder.remaining_tasks}</span>
                </div>
            </li>
        });

        return (
            <div id="folders">
                <span onClick={()=>this.addNewFolder()} id="add-folder">+</span>
                <ul>{folders}</ul>
            </div>
        );
    }
}

export default Folders;