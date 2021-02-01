// @flow
export type AccountBalanceState = {
  symbol: string,
  balance: number,
  subscribed: boolean,
};

export type AccountBalancesState = { +[string]: AccountBalanceState };

export type AccountBalance = {
  symbol: string,
  balance: number,
};

export type AccountBalances = Array<AccountBalance>;
export type AccountBalancesMap = { [string]: AccountBalance }

export type SubscribeAccountBalanceAction = {
  type: 'accountBalances/SUBSCRIBE_BALANCE',
  payload: { symbol: string },
};

export type UpdateAccountBalanceAction = {
  type: 'accountBalances/UPDATE_BALANCE',
  payload: AccountBalance,
};

export type UnsubscribeAccountBalanceAction = {
  type: 'accountBalances/UNSUBSCRIBE_BALANCE',
  payload: { symbol: string },
};

export type UpdateAccountBalancesAction = {
  type: 'accountBalances/UPDATE_BALANCES',
  payload: { balances: AccountBalances },
};

export type ClearAccountBalancesAction = {
  type: 'accountBalances/CLEAR_BALANCES',
};

export type AccountBalancesEvent = any => AccountBalancesState => AccountBalancesState;

export type AccountBalancesAction =
  | SubscribeAccountBalanceAction
  | UpdateAccountBalanceAction
  | UpdateAccountBalancesAction
  | UnsubscribeAccountBalanceAction;
