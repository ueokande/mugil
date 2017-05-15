import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import LoginForm from './login-form';

class Login extends Component {
  render() {
    return (
      <div>
        <hgroup>
          <h1>Mugil tasks</h1>
          <h3>Accelerate your task processing</h3>
        </hgroup>
        <LoginForm />
      </div>
    );
  }
}

window.addEventListener('load', () => {
  var login = document.getElementById('mugil-login');
  if (login !== null) {
    ReactDOM.render(
      <MuiThemeProvider>
        <Login />
      </MuiThemeProvider>,
      login
    );
  }
});
