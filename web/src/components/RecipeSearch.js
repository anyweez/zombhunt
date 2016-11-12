import React, { Component } from 'react';
import Reducers from '../actions/recipe';

import { SearchRecipes } from '../api/recipes';

class RecipeSearch extends Component {
    constructor() {
        super();

        this.state = {
            term: '',
            // count: 0,
            results: [],
        };
    }

    search(term) {
        this.setState({ term }, () => {
            if (this.state.term.length > 2) {
                SearchRecipes(this.state.term, results => this.setState({ results }));
            } else {
                this.setState({ results: [] });
            }
        });
    }

    updateCount(count) {
        this.setState({ count });
        this.props.state.dispatch({ type: Reducers.RECIPE_UPDATE_QUANTITY, target: count });
    }

    selectRecipe(recipe) {
        this.setState({ results: [] });
        this.props.state.dispatch({ type: Reducers.RECIPE_UPDATE_SELECTED, target: recipe });
    }

    render() {
        const state = this.props.state.getState();
        const autocomplete = this.state.results.map((result, i) => (<li key={i} onClick={() => this.selectRecipe(result)}>{ result.Name }</li>));

        return (
            <div>
                <div className="autocomplete">
                    <input type="text" className="r-search" value={this.state.term} onChange={event => this.search(event.target.value)} placeholder="Search for recipe" />
                    <ul>
                        { autocomplete }
                    </ul>
                </div>
                <input type="text" className="r-quantity" value={state.count} onChange={event => this.updateCount(event.target.value)} placeholder="#" />
            </div>
        );
    }
}

export default RecipeSearch;
