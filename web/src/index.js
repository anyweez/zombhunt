import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';

import './index.css';
import store from './store';

import { GetInventories } from './api/inventory';

function render() {
  ReactDOM.render(
    <App state={store} />,
    document.getElementById('root')
  );
}

render();
store.subscribe(render);

// Load initial data
GetInventories();
// setInterval(() => GetInventories(), 10000);