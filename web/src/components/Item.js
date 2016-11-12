import React, { Component } from 'react';

class Item extends Component {
    render() {
        const classNames = ['item', this.props.item.highlight ? 'highlight' : 'dim'];

        return (
            <div className={ classNames.join(' ') }>
                <span className="itemCount">{this.props.item.Quantity}x</span> 
                <span className="itemName">{ this.props.item.Name }</span>
            </div>
        );
    }
}

export default Item;
