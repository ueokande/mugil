import React, { Component } from 'react';
import TextField from 'material-ui/TextField';
import RaisedButton from 'material-ui/RaisedButton';
import CsrfToken from '../ui/csrf-token';
import PropTypes from 'prop-types';


export default class LoginForm extends Component {
  handleLogin() {
    let email = this.email.getValue()
    let password = this.password.getValue()
    this.props.login(email, password)
  }

  render() {
    return (
      <div className='login-form'>
        <TextField
          ref={(e) => { this.email = e }}
          hintText="Email"
          floatingLabelText="Email"
          name="email"
          fullWidth={true}
        /><br />

        <TextField
          ref={(e) => { this.password = e }}
          hintText="Password Field"
          floatingLabelText="Password"
          type="password"
          name="password"
          fullWidth={true}
        /><br />

        <RaisedButton
          label="Login"
          onTouchTap={(e) => { this.handleLogin() }}
          primary={true}
          fullWidth={true}
        />
      </div>
    )
  }
}

LoginForm.defaultProps = {
};

LoginForm.propTypes = {
  login: PropTypes.func.isRequired,
}
