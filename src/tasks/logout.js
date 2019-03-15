import React,{Component} from 'react';
import './logout.css';

class Logout extends Component {

    constructor(props,context) {
        super(props,context);
    }
    render() {
        return (
            <div id="logout-text" onClick={()=>{this.props.onLogout()}}>
                <p>Logout</p>
            </div>
        );
    }
}

export default Logout;