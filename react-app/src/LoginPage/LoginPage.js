import React, { Component } from 'react';
import './newLoginPage.css';
import UserStore from './userStore';
import axios from 'axios';
import ServerURLFetcher from './../ServerURLFetcher';

class LoginPage extends Component {

    constructor(props, context) {
        super(props, context);
        this.state = {
            login_password: '',
            registration_email: '',
            showRegistrationForm: true,
            loginCodeSent: false
        }
        this.handleLoginPasswordChange = this.handleLoginPasswordChange.bind(this);
        this.handleRegistrationEmailChange = this.handleRegistrationEmailChange.bind(this);
        this.handleSubmitLogin = this.handleSubmitLogin.bind(this);
        this.handleSubmitRegister = this.handleSubmitRegister.bind(this);
        this.showRegistrationForm = this.showRegistrationForm.bind(this);
        this.hideRegistrationForm = this.hideRegistrationForm.bind(this);
    }

    handleLoginPasswordChange(event) {
        this.setState({ login_password: event.target.value });
    }

    handleRegistrationEmailChange(event) {
        this.setState({ registration_email: event.target.value });
    }

    handleSubmitLogin(event) {
        event.preventDefault();
        let self = this;
        axios({
            method: 'post',
            url: ServerURLFetcher.getURL() + '/user/verify',
            data: {
                email: this.state.registration_email,
                login_code: this.state.login_password
            },
            json: true
        }).then(function (response) {
            UserStore.saveUser({ 'user': response.data.user, 'token': response.data.token });
            self.props.history.push('/');
        }).catch(function (err) {
            alert(err);
        });
    }

    /**
     * Just send the email to the server and if successfull show the login code section
     * @param {event} event 
     */
    handleSubmitRegister(event) {
        event.preventDefault();
        let self = this;
        self.setState({ loginCodeSent: true })
        axios({
            method: 'post',
            url: ServerURLFetcher.getURL()+ '/user/generate',
            data: {
                email: this.state.registration_email
            },
            json: true
        }).then(function (response) {
            // Show the login code section
            self.hideRegistrationForm(event);
        }).catch(function (err) {
            self.setState({ loginCodeSent: false })
            alert(err);
        });
    }

    showRegistrationForm(event) {
        event.preventDefault();
        this.setState({ showRegistrationForm: true });
    }

    hideRegistrationForm(event) {
        event.preventDefault();
        this.setState({ showRegistrationForm: false });
    }

    render() {
        return (
            <div id="login-page-container">
                <div class="login-page">
                    <div class="form">
                        <form className={this.state.showRegistrationForm ? "register-form" : "hidden"} onSubmit={this.handleSubmitRegister}>
                            <p className={this.state.loginCodeSent ? "" : "hidden"}>Please wait while a login code is being sent to your email id...</p>
                            <input type="email" placeholder="email" value={this.state.registration_email}
                                onChange={this.handleRegistrationEmailChange} />
                            <button>Login</button>
                        </form>
                        <form className={this.state.showRegistrationForm ? 'hidden' : "login-form"} onSubmit={this.handleSubmitLogin}>
                            <p>Enter the login code sent on your email id.</p>
                            <input type="password" placeholder="Login Code" value={this.login_password}
                                onChange={this.handleLoginPasswordChange} />
                            <button>login</button>
                            <p class="message">Login With a Different Email Id? <a href="#" onClick={this.showRegistrationForm}>Change Email ID</a></p>
                        </form>
                    </div>
                </div>
            </div>
        );
    }
}

export default LoginPage;