import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import EntryList from './entry-list';
import AppBar from 'material-ui/AppBar';

export default class TaskIndex extends Component {
  constructor() {
    super();
    this.state = {
      entries: []
    };
  }

  componentDidMount() {
    fetch('/tasks.json?date=2017-05-09', {
      method: 'GET',
      credentials: 'same-origin',
      headers: {
	'Content-Type': 'application/json'
      }
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
      this.setState({
	      entries: entries
      })
    });
  }

  render() {
    return (
      <div>
        <h1></h1>
        <EntryList entries={this.state.entries}></EntryList>
      </div>
    )
  }
}

window.addEventListener('load', () => {
  var parent = document.getElementById('mugil-task-index');
  if (parent !== null) {
    ReactDOM.render(
      <MuiThemeProvider>
        <div>
          <AppBar
            title="Mugil tasks"
            iconClassNameRight="muidocs-icon-navigation-expand-more"
          />
          <TaskIndex></TaskIndex>
        </div>
      </MuiThemeProvider>,
      parent 
    );
  }
});
