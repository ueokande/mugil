import React, { Component } from 'react';
import Dialog from 'material-ui/Dialog';
import MenuItem from 'material-ui/MenuItem';
import PropTypes from 'prop-types';
import SelectField from 'material-ui/SelectField';
import TextField from 'material-ui/TextField';
import FlatButton from 'material-ui/FlatButton';

export default class TaskForm extends Component {
  constructor() {
    super();
    this.state = {
      priority: "A",
      estimatedTime: 30 * 60 * 1e9,
    }
  }


  handlePriorityChange(event, index, value) {
    this.setState({
      priority: value
    });
  }

  handleEstimatedTimeChange(event, index, value) {
    this.setState({
      estimatedTime: value
    });
  }

  handleCreate() {
    let priority = this.state.priority;
    let estimatedTime  = this.state.estimatedTime;
    let description  = this.description.getValue();
    this.props.create(priority, estimatedTime, description)
  }

  render() {
    return (
      <Dialog open={this.props.open} >
        <SelectField
          floatingLabelText="Priority"
          name="priority"
          value={this.state.priority}
          onChange={(e, i, v) => this.handlePriorityChange(e, i, v)}
        >
          <MenuItem value="S" primaryText="S" />
          <MenuItem value="A" primaryText="A" />
          <MenuItem value="B" primaryText="B" />
          <MenuItem value="C" primaryText="C" />
        </SelectField>

        <SelectField
          floatingLabelText="Estimated Time"
          name="estimated_time"
          value={this.state.estimatedTime}
          onChange={(e, i, v) => this.handleEstimatedTimeChange(e, i, v)}
        >
          <MenuItem value={15 * 60 * 1e9} primaryText="15 min." />
          <MenuItem value={30 * 60 * 1e9} primaryText="30 min." />
          <MenuItem value={60 * 60 * 1e9} primaryText="60 min." />
          <MenuItem value={90 * 60 * 1e9} primaryText="1h 30 min." />
          <MenuItem value={120 * 60 * 1e9} primaryText="2h" />
          <MenuItem value={180 * 60 * 1e9} primaryText="3h" />
          <MenuItem value={240 * 60 * 1e9} primaryText="4h" />
        </SelectField>
        <br />

        <TextField
          ref={(e) => { this.description = e; }}
          hintText="Clean my room"
          floatingLabelText="Description"
          name="description"
          fullWidth={true}
        /><br />

        <div style={{ textAlign: "right" }}>
          <FlatButton label="Cancel" onTouchTap={this.props.cancel} />
          <FlatButton label="Create" onTouchTap={() => this.handleCreate()} />
        </div>
      </Dialog>
    );
  }
}

TaskForm.defaultProps = {
};

TaskForm.propTypes = {
  open: PropTypes.bool.isRequired,
  create: PropTypes.func.isRequired,
  cancel: PropTypes.func.isRequired,
};
