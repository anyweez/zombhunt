import Reducers from '../actions/recipe';

const initial = {
    recipe: null,
    count: 1,
    inventories: [],
};

/**
 * Set the 'highlight' property of each item based on the chosen recipe.
 */
function highlight(inventory, recipe) {
    inventory.Inventory = inventory.Inventory.map(item => {
        item.highlight = recipe.Ingredients.filter(ingredient => item.Name === ingredient.Name).length > 0;
        return item;
    });

    return inventory;
}

function updateIngredientCount(recipe, multiplier = 1) {
    recipe.Quantity *= multiplier;

    if (recipe.Ingredients !== null && recipe.Ingredients.length > 0) {
        recipe.Ingredients = recipe.Ingredients.map(ingredient => updateIngredientCount(ingredient, recipe.Quantity));
    }

    return recipe;
}

function appendTotals(recipe, inventories) {
    if (recipe.Ingredients === null) return recipe;

    console.log(recipe);
    recipe.Ingredients = recipe.Ingredients.map(ingredient => {
        // Each ingredient gets a count of the number of that ingredient available across all inventories.
        ingredient.Available = inventories
            .map(inventory => {
                return inventory.Inventory
                    .map(item => item.Name === ingredient.Name ? item.Quantity : 0)
                    .reduce((total, next) => total + next, 0);
            }).reduce((total, next) => total + next, 0);

        return ingredient;
    })

    // Recurse.
    recipe.Ingredients = recipe.Ingredients.map(ingredient => appendTotals(ingredient, inventories));

    return recipe;
}

export default (state = initial, action) => {
    let next = Object.assign({}, state);

    switch (action.type) {
        case Reducers.RECIPE_UPDATE_QUANTITY:
            next.count = action.target.length === 0 ? 0 : parseInt(action.target, 10);

            break;

        case Reducers.INVENTORY_UPDATE:
            next.inventories = action.target;

            if (next.recipe !== null) {
                next.inventories = next.inventories.map(inventory => highlight(inventory, next.recipe));
                next.recipe = appendTotals(next.recipe, next.inventories);
            }

            break;

        case Reducers.RECIPE_UPDATE_SELECTED:
            next.recipe = appendTotals(updateIngredientCount(action.target), next.inventories);

            next.inventories = next.inventories.map(inventory => highlight(inventory, next.recipe));
            break;

        default:
            console.log('default event');
            break;
    }

    return next;
};