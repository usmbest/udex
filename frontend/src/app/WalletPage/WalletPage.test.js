import React from 'react';
import { shallow } from 'enzyme';
import WalletPage from './WalletPage';

const tokenTableData = [
  { symbol: 'EOS', balance: '10.0000' },
  { symbol: 'ZRX', balance: '1.00000' },
  { symbol: 'EOS', balance: '5.00000' },
  { symbol: 'EOS', balance: '8.00000' },
];

it('renders without crashing', () => {
  shallow(<WalletPage queryAccountData={jest.fn()} />);
});

it('calls queryAccountData on mount', () => {
  const queryAccountData = jest.fn();
  shallow(
    <WalletPage
      authenticated={true}
      queryAccountData={queryAccountData}
      tokenTableData={tokenTableData}
      loading={false}
    />
  );

  expect(queryAccountData).toBeCalled();
});
