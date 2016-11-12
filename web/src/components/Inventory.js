import React, { Component } from 'react';

import Item from './Item';

class Inventory extends Component {
    render() {
        const displayItems = this.props.inv.Inventory.map((item, i) => (<Item key={i} item={item} />));

        return (
            <section className="inventory">
                <div className="intro">
                    <img alt="profile" src={this.props.inv.AvatarUrl} />
                    <h3>{this.props.inv.Name}</h3>
                    <h4>{this.props.inv.Inventory.length} items</h4>
                </div>
                <div className="items">
                    { displayItems }
                </div>
            </section>
        );
    }
}

export default Inventory;
