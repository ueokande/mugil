import React, { Component } from 'react';
import AppBar from 'material-ui/AppBar';
import ContentAdd from 'material-ui/svg-icons/content/add';
import FloatingActionButton from 'material-ui/FloatingActionButton';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';

import TaskFormDialog from './task-form-dialog';
import EntryList from './entry-list';

import { getCsrfToken } from '../../shared/csrf'

const floatingActionButtonStyle = {
  margin: 0,
  top: 'auto',
  right: 20,
  bottom: 20,
  left: 'auto',
  position: 'fixed',
};

function checkStatus(response) {
  if (response.status >= 200 && response.status < 300) {
    return response
  } else {
    var error = new Error(response.statusText)
    error.response = response
    throw error
  }
}

export default class TaskIndex extends Component {
  constructor() {
    super();
    this.state = {
      formOpen: false,
      entries: []
    };
  }

  componentDidMount() {
    fetch('/tasks?date=2017-05-16', {
      method: 'GET',
      credentials: 'same-origin',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    .then(checkStatus)
    .then((response) => {
      return response.json()
    })
    .then((entries) => {
      this.setState({
        entries: entries
      })
    });
  }

  handleCreate(priority, estimatedTime, description) {
    fetch('/tasks', {
      method: 'POST',
      credentials: 'same-origin',
      headers: {
        'Content-Type': 'application/json',
        'X-CSRF-TOKEN': getCsrfToken()
      },
      body: JSON.stringify({
        priority: priority,
        estimated_time: estimatedTime,
        description: description
      })
    })
    .then(checkStatus)
    .then(() => {
      this.setState({
        formOpen: false,
      })
    })
  }

  handleNewFormOpen() {
    this.setState({
      formOpen: true
    })
  }

  handleCancel() {
    this.setState({
      formOpen: false
    })
  }

  render() {
    return (
      <MuiThemeProvider>
        <div>
          <AppBar
            title="Mugil tasks"
            iconClassNameRight="muidocs-icon-navigation-expand-more"
          />
          <EntryList entries={this.state.entries}></EntryList>
          <FloatingActionButton
            style={floatingActionButtonStyle}
            onTouchTap={() => this.handleNewFormOpen()}
          >
            <ContentAdd />
          </FloatingActionButton>
          <TaskFormDialog
            open={this.state.formOpen}
            priority="B"
            estimatedTime={30 * 60 * 1e9}
            create={(p, e, d) => this.handleCreate(p, e, d)}
            cancel={() => this.handleCancel()}
          />
        </div>
      </MuiThemeProvider>
    )
  }
}

