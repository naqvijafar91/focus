import React, { Component } from 'react';
import './newLoginPage.css';
import UserStore from './userStore';

class LoginPage extends Component {

    constructor(props, context) {
        super(props, context);
        this.state = {
            login_email: '',
            login_password: '',
            registration_email: '',
            registration_password: '',
            registration_confirmed_password: ''
        }

        this.handleLoginEmailChange = this.handleLoginEmailChange.bind(this);
        this.handleLoginPasswordChange = this.handleLoginPasswordChange.bind(this);
        this.handleRegistrationEmailChange = this.handleRegistrationEmailChange.bind(this);
        this.handleRegistrationPasswordChange = this.handleRegistrationPasswordChange.bind(this);
        this.handleSubmitLogin = this.handleSubmitLogin.bind(this);
    }

    handleLoginEmailChange(event) {
        this.setState({ login_email: event.target.value });
    }

    handleLoginPasswordChange(event) {
        this.setState({ login_password: event.target.value });
    }

    handleRegistrationEmailChange(event) {
        this.setState({ registration_email: event.target.value });
    }

    handleRegistrationPasswordChange(event) {
        this.setState({ registration_password: event.target.value });
    }

    handleSubmitLogin(event) {
        event.preventDefault();
        //@Todo:Hit api 
        UserStore.saveUser({ user: '', 'token': '' });
        this.props.history.push('/');
    }

    handleSubmitRegister(event) {
        event.preventDefault();
        //@Todo:Hit api 
        UserStore.saveUser({ user: '', 'token': '' });
        this.props.history.push('/');
    }

    render() {
        return (
            <div id="login-page-container">
                <div class="login-page">
                    <div class="form">
                        <form class="register-form" onSubmit={this.handleSubmitRegister}>
                            <input type="email" placeholder="email" value={this.state.registration_email}
                                onChange={this.handleRegistrationEmailChange} />
                            <input type="password" placeholder="password" value={this.state.registration_password}
                                onChange={this.handleRegistrationPasswordChange} />
                            <button>create</button>
                            <p class="message">Already registered? <a href="#">Sign In</a></p>
                        </form>
                        <form class="login-form" onSubmit={this.handleSubmitLogin}>
                            <input type="email" placeholder="email" value={this.login_email}
                                onChange={this.handleLoginEmailChange} />
                            <input type="password" placeholder="password" value={this.login_password}
                                onChange={this.handleLoginPasswordChange} />
                            <button>login</button>
                            <p class="message">Not registered? <a href="#">Create an account</a></p>
                        </form>
                    </div>
                </div>
            </div>
        );
    }
}

export default LoginPage;