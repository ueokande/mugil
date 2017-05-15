import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import LoginForm from './login-form';
import { getCsrfToken } from '../shared/csrf'

class Login extends Component {
  handleLogin(email, password) {
    fetch('/login', {
      method: 'POST',
      credentials: 'same-origin',
      headers: {
        'Content-Type': 'application/json',
        'X-CSRF-TOKEN': getCsrfToken()
      },
      body: JSON.stringify({
        email: email,
        password: password,
      })
    })
    .then((response) => {
      if (response.status >= 200 && response.status < 300) {
        return response
      } else {
        var error = new Error(response.statusText)
        error.response = response
        throw error
      }

    })
    .then((response) => {
      return response.json()
    })
    .then((entries) => {
      window.location.href = "/"
    });
  }

  render() {
    return (
      <div>
        <hgroup>
          <h1>Mugil tasks</h1>
          <h3>Accelerate your task processing</h3>
        </hgroup>
        <LoginForm
          login={this.handleLogin}
        />
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
