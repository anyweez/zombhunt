import React, { Component } from 'react';
import './App.css';

import RecipeSearch from './components/RecipeSearch';
import Recipe from './components/Recipe';
import Inventory from './components/Inventory';

class App extends Component {
  render() {
    const state = this.props.state.getState();
    const inventories = state.inventories.map((user, i) => (<Inventory inv={user} key={i} />))

    return ( 
      <div id="app">
        <header>
          <h1>Zomb Hunt</h1>
          <RecipeSearch state={this.props.state} />
        </header>
        <main>
          <Recipe state={this.props.state} />
          <section id="inventories">
            {inventories}
          </section>
        </main>
        <footer>
        </footer>
      </div>
    );
  }
}


export default App;
