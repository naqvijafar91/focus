import React, { Component } from 'react';

class FolderItem extends Component {
    
    render() {
        return (
            <li>
            <div>
                <span>Inbox</span>
                <span class="remaining-tasks">13</span>
            </div>
        </li>    
    )
    }
}


export default FolderItem;