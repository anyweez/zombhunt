import store from '../store';
import Reducers from '../actions/recipe';

export function GetInventories() {
    fetch('http://localhost:8080/inventories')
        .then(res => res.json())
        .then(result => {
            result = result.map(user => {
                user.Inventory = Object.keys(user.Inventory.Items).map(key => user.Inventory.Items[key]);
                return user;
            });

            store.dispatch({ type: Reducers.INVENTORY_UPDATE, target: result });
        });  
};