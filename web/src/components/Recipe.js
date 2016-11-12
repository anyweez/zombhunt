import React, { Component } from 'react';

class Recipe extends Component {
    render() {
        const state = this.props.state.getState();
        const tiers = [];

        if (state.recipe !== null) {
            tiers.push(state.recipe.Ingredients.map((ingredient, i) => {
                const classNames = ['item', ingredient.Available > ingredient.Quantity * state.count ? 'meets-requirements' : 'beneath-requirements'];

                return (
                    <li className={classNames.join(' ')} key={i}>
                        <span className="itemCount">{ ingredient.Available } of { ingredient.Quantity * state.count }</span> 
                        <span className="itemName">{ ingredient.Name }</span>
                    </li>
                );
            }));

            const tier1 = state.recipe.Ingredients
                .map(ingredient => ingredient.Ingredients)
                .reduce((full, next) => full.concat(next), [])
                .filter(ingredient => ingredient !== null);

            tiers.push(tier1.map((ingredient, i) => {
                const classNames = ['item', ingredient.Available > ingredient.Quantity * state.count ? 'meets-requirements' : 'beneath-requirements'];

                return (
                    <li className={classNames.join(' ')} key={i}>
                        <span className="itemCount">{ ingredient.Available } of { ingredient.Quantity * state.count }</span> 
                        <span className="itemName">{ ingredient.Name }</span>
                    </li>
                );
            }));
        }

        const itemLists = tiers
            .map((tier, i) => (<ul key={i} className="items">{ tier }</ul>));

        return (
            <section className="recipe">
                <h2>{ (state.recipe !== null) ? state.recipe.Name : 'Recipe' }</h2>
                { itemLists }
            </section>
        );
    }
}

export default Recipe;
