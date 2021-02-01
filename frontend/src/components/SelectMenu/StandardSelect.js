import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { Button, MenuItem } from '@blueprintjs/core';
import { Select } from '@blueprintjs/select';

class StandardSelect extends Component {
  renderItem(item, { handleClick, modifiers }) {
    return (
      <MenuItem
        active={modifiers.active}
        disabled={modifiers.disabled}
        onClick={rank => handleClick(item.rank)}
        text={`${item.name}`}
        key={item.rank}
      />
    );
  }

  render() {
    const { items, item, icon, handleChange, minimal } = this.props;
    return (
      <Select
        items={items}
        filterable={false}
        itemRenderer={this.renderItem}
        noResults={<MenuItem disabled text="No results." />}
        onItemSelect={handleChange}
        popoverProps={{ minimal: true }}
      >
        <Button icon={item.icon ? item.icon : icon} text={!minimal && item.name} righticonname="double-caret-vertical" />
      </Select>
    );
  }
}

StandardSelect.propTypes = {
  item: PropTypes.object,
  items: PropTypes.array,
  handleChange: PropTypes.func,
};

export default StandardSelect;
