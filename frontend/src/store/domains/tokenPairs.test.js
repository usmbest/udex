import * as eventCreators from './tokenPairs';
import { generateTokenPairs } from '../../utils/tokens';
import getTokenPairsDomain from './tokenPairs';

//createInitialState is not an eventCreator. We simply import it in order to create a new
//create an initial state. The default initial state used in the application has to many
//tokens to be used for tests. Therefore we recreate an initial state with less tokens
//to test the token pair model
const symbols = ['ETH', 'EOS', 'WETH', 'ZRX'];

const quoteTokens = [
  { symbol: 'WETH', asset: '0x3', decimals: 18 },
  { symbol: 'DAI', asset: '0x2', decimals: 18 },
]

const tokensBySymbol = {
  EOS: { symbol: 'EOS', asset: '0x1', decimals: 18 },
  WETH: { symbol: 'WETH', asset: '0x3', decimals: 18 },
  ZRX: { symbol: 'ZRX', asset: '0x4', decimals: 18 },
};

const tokens = Object.values(tokensBySymbol);

const initialTokenPairState = {
  byPair: generateTokenPairs(quoteTokens, tokens),
  data: {},
  favorites: [],
  currentPair: 'WETH/USDC',
  sortedPairs: [],
};

function getDomain(events) {
  const state = events.reduce((state, event) => event(state), undefined);
  return getTokenPairsDomain(state);
}

describe('Token Pair Domain', () => {
  it('handles initialized event properly', () => {
    const tokenPairsDomain = getDomain([eventCreators.initialized(initialTokenPairState)]);
    const expectedPairs = ['EOS/WETH', 'EOS/DAI', 'WETH/DAI', 'ZRX/WETH', 'ZRX/DAI'];

    const expectedByPairsByCode = {
      'EOS/DAI': {
        pair: 'EOS/DAI',
        baseTokenSymbol: 'EOS',
        quoteTokenSymbol: 'DAI',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x1',
        quoteAsset: '0x2',
      },
      'EOS/WETH': {
        pair: 'EOS/WETH',
        baseTokenSymbol: 'EOS',
        quoteTokenSymbol: 'WETH',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x1',
        quoteAsset: '0x3',
      },
      'ZRX/DAI': {
        pair: 'ZRX/DAI',
        baseTokenSymbol: 'ZRX',
        quoteTokenSymbol: 'DAI',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x4',
        quoteAsset: '0x2',
      },
      'ZRX/WETH': {
        pair: 'ZRX/WETH',
        baseTokenSymbol: 'ZRX',
        quoteTokenSymbol: 'WETH',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x4',
        quoteAsset: '0x3',
      },
      'WETH/DAI': {
        pair: 'WETH/DAI',
        baseTokenSymbol: 'WETH',
        quoteTokenSymbol: 'DAI',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x3',
        quoteAsset: '0x2',
      },
    };

    expect(tokenPairsDomain.getPairs()).toEqual(expectedPairs);
    expect(tokenPairsDomain.getPairsByCode()).toEqual(expectedByPairsByCode);
  });

  it('handles tokenPairsUpdated event properly', () => {

    const pairs = [
      {
        pair: 'EOS/DAI',
        baseTokenSymbol: 'EOS',
        quoteTokenSymbol: 'DAI',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x1',
        quoteAsset: '0x2',
      },
      {
        pair: 'EOS/WETH',
        baseTokenSymbol: 'EOS',
        quoteTokenSymbol: 'WETH',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x1',
        quoteAsset: '0x3',
      },
      {
        pair: 'ZRX/DAI',
        baseTokenSymbol: 'ZRX',
        quoteTokenSymbol: 'DAI',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x4',
        quoteAsset: '0x2',
      },
      {
        pair: 'ZRX/WETH',
        baseTokenSymbol: 'ZRX',
        quoteTokenSymbol: 'WETH',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x4',
        quoteAsset: '0x3',
      },
      {
        pair: 'WETH/DAI',
        baseTokenSymbol: 'WETH',
        quoteTokenSymbol: 'DAI',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x3',
        quoteAsset: '0x2',
      }
    ];

    const tokenPairsDomain = getDomain([
      eventCreators.initialized(initialTokenPairState),
      eventCreators.tokenPairsUpdated(pairs),
    ]);

    const expectedPairs = ["EOS/WETH", "EOS/DAI", "WETH/DAI", "ZRX/WETH", "ZRX/DAI"]

    const expectedPairsBySymbol = {
      'EOS/DAI': {
        pair: 'EOS/DAI',
        baseTokenSymbol: 'EOS',
        quoteTokenSymbol: 'DAI',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x1',
        quoteAsset: '0x2',
      },
      'EOS/WETH': {
        pair: 'EOS/WETH',
        baseTokenSymbol: 'EOS',
        quoteTokenSymbol: 'WETH',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x1',
        quoteAsset: '0x3',
      },
      'ZRX/DAI': {
        pair: 'ZRX/DAI',
        baseTokenSymbol: 'ZRX',
        quoteTokenSymbol: 'DAI',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x4',
        quoteAsset: '0x2',
      },
      'ZRX/WETH': {
        pair: 'ZRX/WETH',
        baseTokenSymbol: 'ZRX',
        quoteTokenSymbol: 'WETH',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x4',
        quoteAsset: '0x3',
      },
      'WETH/DAI': {
        pair: 'WETH/DAI',
        baseTokenSymbol: 'WETH',
        quoteTokenSymbol: 'DAI',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x3',
        quoteAsset: '0x2',
      }
    };

    expect(tokenPairsDomain.getPairs()).toEqual(expectedPairs);
    expect(tokenPairsDomain.getPairsByCode()).toEqual(expectedPairsBySymbol);
  });

  
  it('handles tokenPairDataUpdated event', () => {
    const tokenPairData = [
      {
        pair: 'WETH/USDC',
        lastPrice: '7425.2945',
        change: '4.5421',
        high: '8782.7964',
        low: '6499.3696',
        volume: 720404,
      },
      {
        pair: 'WETH/DAI',
        lastPrice: '6018.7886',
        change: '1.6589',
        high: '3876.8717',
        low: '4613.5315',
        volume: 68946,
      }
    ]
    
    const expectedTokenPairArray = {
      'WETH/USDC': {
        pair: 'WETH/USDC',
        lastPrice: '7425.2945',
        change: '4.5421',
        high: '8782.7964',
        low: '6499.3696',
        volume: 720404,
      },
      'WETH/DAI': {
        pair: 'WETH/DAI',
        lastPrice: '6018.7886',
        change: '1.6589',
        high: '3876.8717',
        low: '4613.5315',
        volume: 68946,
      }
    };

    const domain = getDomain([eventCreators.initialized(), eventCreators.tokenPairDataUpdated(tokenPairData)]);
    expect(domain.getTokenPairsData()).toEqual(expectedTokenPairArray);
  });

  it('handles tokenPairFavorited event', () => {
    const domain = getDomain([
      eventCreators.initialized(initialTokenPairState),
      eventCreators.tokenPairFavorited('EOS/WETH', true),
    ]);

    expect(domain.getFavoritePairs()).toEqual(['EOS/WETH']);
  });

  it('handles tokenPairFavorited events', () => {
    const domain = getDomain([
      eventCreators.initialized(initialTokenPairState),
      eventCreators.tokenPairFavorited('EOS/WETH', true),
      eventCreators.tokenPairFavorited('EOS/DAI', true),
      eventCreators.tokenPairFavorited('EOS/ZRX', true),
      eventCreators.tokenPairFavorited('EOS/WETH', false),
    ]);

    expect(domain.getFavoritePairs()).toEqual(['EOS/DAI', 'EOS/ZRX']);
  });

  it('returns getTokensPairsWithData correctly', () => {
    const pairs = [
      {
        pair: 'EOS/DAI',
        baseTokenSymbol: 'EOS',
        quoteTokenSymbol: 'DAI',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x1',
        quoteAsset: '0x2',
      },
      {
        pair: 'EOS/WETH',
        baseTokenSymbol: 'EOS',
        quoteTokenSymbol: 'WETH',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x1',
        quoteAsset: '0x3',
      },
      {
        pair: 'ZRX/DAI',
        baseTokenSymbol: 'ZRX',
        quoteTokenSymbol: 'DAI',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x4',
        quoteAsset: '0x2',
      },
      {
        pair: 'ZRX/WETH',
        baseTokenSymbol: 'ZRX',
        quoteTokenSymbol: 'WETH',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x4',
        quoteAsset: '0x3',
      },
      {
        pair: 'WETH/DAI',
        baseTokenSymbol: 'WETH',
        quoteTokenSymbol: 'DAI',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x3',
        quoteAsset: '0x2',
      }
    ];

    const tokenPairData = [
      {
        pair: 'EOS/DAI',
        lastPrice: '7425.2945',
        change: '4.5421',
        high: '8782.7964',
        low: '6499.3696',
        volume: 720404,
        orderbookSize: '100000',
      },
      {
        pair: 'WETH/DAI',
        lastPrice: '6018.7886',
        change: '1.6589',
        high: '3876.8717',
        low: '4613.5315',
        volume: 68946,
        orderbookSize: '100000',
      }
    ]
  
    const tokenPairsDomain = getDomain([
      eventCreators.initialized(initialTokenPairState),
      eventCreators.tokenPairsUpdated(pairs),
      eventCreators.tokenPairDataUpdated(tokenPairData)
    ]);

    expect(tokenPairsDomain.getTokenPairsWithDataObject()).toEqual({
      'EOS/DAI': {
        pair: 'EOS/DAI',
        baseTokenSymbol: 'EOS',
        quoteTokenSymbol: 'DAI',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x1',
        quoteAsset: '0x2',
        lastPrice: '7425.2945',
        change: '4.5421',
        orderbookSize: '100000',
        high: '8782.7964',
        low: '6499.3696',
        volume: 720404,
      },
      'WETH/DAI': {
        pair: 'WETH/DAI',
        baseTokenSymbol: 'WETH',
        quoteTokenSymbol: 'DAI',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x3',
        quoteAsset: '0x2',
        lastPrice: '6018.7886',
        change: '1.6589',
        orderbookSize: '100000',
        high: '3876.8717',
        low: '4613.5315',
        volume: 68946,
      }
    })

    expect(tokenPairsDomain.getTokenPairsWithDataArray()).toEqual([
      {
        pair: 'EOS/DAI',
        baseTokenSymbol: 'EOS',
        quoteTokenSymbol: 'DAI',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x1',
        quoteAsset: '0x2',
        lastPrice: '7425.2945',
        change: '4.5421',
        orderbookSize: '100000',
        high: '8782.7964',
        low: '6499.3696',
        volume: 720404,
      },
      {
        pair: 'WETH/DAI',
        baseTokenSymbol: 'WETH',
        quoteTokenSymbol: 'DAI',
        baseTokenDecimals: 18,
        quoteTokenDecimals: 18,
        baseAsset: '0x3',
        quoteAsset: '0x2',
        lastPrice: '6018.7886',
        change: '1.6589',
        orderbookSize: '100000',
        high: '3876.8717',
        low: '4613.5315',
        volume: 68946,
      }
    ]
    )
  });
});
