import React from 'react';
import TextField from 'material-ui/TextField';
import RaisedButton from 'material-ui/RaisedButton';
import CsrfToken from '../shared/csrf-token';

const LoginForm = () => (
  <form action="/login" method="POST">

    <CsrfToken />

    <TextField
      hintText="Email"
      floatingLabelText="Email"
      name="email"
    /><br />

    <TextField
      hintText="Password Field"
      floatingLabelText="Password"
      type="password"
      name="password"
    /><br />

    <RaisedButton
      type="submit"
      label="Login"
      primary={true}
      fullWidth={true}
    />
  </form>
);

export default LoginForm;
