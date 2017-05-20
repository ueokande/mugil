import React, { Component } from 'react';
import AppBar from 'material-ui/AppBar';
import ContentAdd from 'material-ui/svg-icons/content/add';
import FloatingActionButton from 'material-ui/FloatingActionButton';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import LinearProgress from 'material-ui/LinearProgress';

import TaskFormDialog from './task-form-dialog';
import EntryList from './entry-list';

import * as form from '../action/form'
import * as tasks from '../action/tasks'

const floatingActionButtonStyle = {
  margin: 0,
  top: 'auto',
  right: 20,
  bottom: 20,
  left: 'auto',
  position: 'fixed',
};

class TaskIndex extends Component {
  constructor() {
    super();
  }

  componentDidMount() {
    let date = new Date().toISOString().split('T')[0]
    this.props.dispatch(tasks.fetchTasks(date));
  }

  handleCreate(priority, estimatedTime, description) {
    this.props.dispatch(tasks.postTask(priority, estimatedTime, description, () => {
      this.props.dispatch(form.close());
    }));
  }

  handleNewFormOpen() {
    this.props.dispatch(form.openNew());
  }

  handleCancel() {
    this.props.dispatch(form.close());
  }

  render() {
    return (
      <MuiThemeProvider>
        <div>
          <AppBar
            title="Mugil tasks"
            iconClassNameRight="muidocs-icon-navigation-expand-more"
          />
          <LinearProgress color="#E91E63" style={{
            display: this.props.loading ? "block" : "none"
          }}/>
          <EntryList entries={this.props.entries}></EntryList>
          <FloatingActionButton
            style={floatingActionButtonStyle}
            onTouchTap={() => this.handleNewFormOpen()}
          >
            <ContentAdd />
          </FloatingActionButton>
          <TaskFormDialog
            open={this.props.formOpen}
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

EntryList.propTypes = {
  formOpen: PropTypes.bool,
  entries: PropTypes.arrayOf(PropTypes.object)
};

export default connect(({form, tasks}) => {
  return {
    formOpen: form.open,
    entries: tasks.entries,
    loading: tasks.loading,
  }
})(TaskIndex)
