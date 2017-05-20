import {combineReducers} from "redux"

import form from './form'
import tasks from './tasks'

export default combineReducers({
  form,
  tasks
});
