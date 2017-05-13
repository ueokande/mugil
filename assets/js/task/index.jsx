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
    // TODO fetch task list from server
    this.setState({
      entries: [{
        id: 1,
        completed: false,
        priority: "A",
        time: 1.8e+12,
        description: "wash my hands",
      }, {
        id: 2,
        completed: false,
        priority: "B",
        time: 3.6e+12,
        description: "watch TV",
      }, {
        id: 3,
        completed: true,
        priority: "S",
        time: 3.6e+12,
        description: "Buy drugs",
      }, {
        id: 4,
        completed: false,
        priority: "A",
        time: 3.6e+12,
        description: "Clean my room",
      }, {
        id: 5,
        completed: false,
        priority: "B",
        time: 3.6e+12,
        description: "wakl with my dog",
      }]
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
