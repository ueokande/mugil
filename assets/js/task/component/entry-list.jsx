import React, { Component } from 'react';
import {
  Table,
  TableBody,
  TableRow,
  TableRowColumn,
} from 'material-ui/Table';
import PropTypes from 'prop-types';

function humanReadableDuration(nanoseconds) {
  let minutes = nanoseconds / 60 / 1e9;
  return `${minutes} min.`
}

export default class EntryList extends Component {
  render() {
    return (
      <Table onRowSelection={this.handleRowSelection}>
        <TableBody>
          {
            this.props.entries.map((entry) => {
              return (
                <TableRow key={entry.id} selected={entry.done}>
                  <TableRowColumn>{entry.priority}</TableRowColumn>
                  <TableRowColumn>{humanReadableDuration(entry.estimated_time)}</TableRowColumn>
                  <TableRowColumn>{entry.description}</TableRowColumn>
                </TableRow>
              );
            })
          }
        </TableBody>
      </Table>
    )
  }
}

EntryList.defaultProps = {
  entries: []
};

EntryList.propTypes = {
  entries: PropTypes.arrayOf(PropTypes.object)
};
