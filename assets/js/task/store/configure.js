import logger  from 'redux-logger';
import { createStore, applyMiddleware } from 'redux';
import rootReducer from '../reducer/index';

export default function createStoreWithMiddleware() {
  const store = applyMiddleware(
    logger
  )(createStore);
  return store(rootReducer);
}
