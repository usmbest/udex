# UDEX Backend

[![Chat on discord](https://img.shields.io/discord/534371689996222485.svg?logo=discord)](https://discord.gg/qHHST4e)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
![Contributions welcome](https://img.shields.io/badge/contributions-welcome-orange.svg)

# The UDEX Decentralized Exchange

The UDEX decentralized exchange is a hybrid decentralized exchange that aims at bringing together the ease of use of centralized exchanges along with the security and privacy features of decentralized exchanges. Orders are matched through an off-chain orderbook. After orders are matched and signed, the decentralized exchange operator has the sole ability to perform a transaction to the AA. This provides for the best UX as the exchange operator is the only party having to interact directly with the blockchain. Exchange users simply sign orders which are broadcasted then to the orderbook. This design enables users to queue and cancel their orders seamlessly.

Several matchers can operate exchanges based on UDEX technology at the same time. They share their orderbooks and exchange all new orders among themselves, thus improving liquidity for all UDEX exchanges. An order can be submitted through any UDEX exchange, however to be matched, both maker and taker orders have to indicate the same matcher. The exchange that was used to submit the order serves as an affliate and can charge a fee from its users.  Anyone can become a matcher or affiliate, or just watch the orders that are being exchanged among the matchers and detect any possible misbehavior by matchers.

# UDEX 백엔드

[! [Chat on discord] (https://img.shields.io/discord/534371689996222485.svg?logo=discord)] (https://discord.gg/qHHST4e)
[! [라이선스] (https://img.shields.io/badge/license-MIT-blue.svg)] (https://opensource.org/licenses/MIT)
! [기고 환영] (https://img.shields.io/badge/contributions-welcome-orange.svg)

# UDEX 탈 중앙화 거래소

UDEX 분산 거래소는 분산 거래소의 보안 및 개인 정보 보호 기능과 함께 중앙 거래소의 사용 편의성을 결합하는 것을 목표로하는 하이브리드 탈 중앙 거래소입니다. 주문은 오프 체인 주문서를 통해 매칭됩니다. 주문이 일치되고 서명 된 후 탈 중앙화 거래소 운영자는 AA에 대한 거래를 수행 할 수있는 유일한 능력을 갖습니다. 이것은 거래소 운영자가 블록 체인과 직접 상호 작용해야하는 유일한 당사자이기 때문에 최상의 UX를 제공합니다. 교환 사용자는 주문서에 브로드 캐스트 된 주문에 서명하기 만하면됩니다. 이 디자인을 통해 사용자는 주문을 원활하게 대기열에 추가하고 취소 할 수 있습니다.

여러 명의 매 처가 동시에 UDEX 기술을 기반으로 교환을 운영 할 수 있습니다. 그들은 주문서를 공유하고 모든 새로운 주문을 서로 교환하여 모든 UDEX 거래소의 유동성을 향상시킵니다. 주문은 UDEX 거래소를 통해 제출할 수 있지만 일치하려면 메이커 주문과 테이커 주문 모두 동일한 일치자를 표시해야합니다. 주문을 제출하는 데 사용 된 거래소는 제휴사 역할을하며 사용자에게 수수료를 부과 할 수 있습니다. 누구든지 매처 또는 제휴사가 될 수 있으며, 매처간에 교환되는 주문을보고 매처의 가능한 오작동을 감지 할 수 있습니다.

# Getting Started

## Requirements

- **mongoDB** version 3.6 or newer ([installation instructions for ubuntu](https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/))
- **rabbitmq** version 3.7.7 or newer ([installation instructions for ubuntu](https://computingforgeeks.com/how-to-install-latest-rabbitmq-server-on-ubuntu-18-04-lts/))
- **golang** latest ([installation instructions for ubuntu](https://github.com/golang/go/wiki/Ubuntu))

## Install

```
go get github.com/citypayorg/udex/backend
```

## Run
You don't run the backend directly. Run [UDEX wallet](https://github.com/citypayorg/udex/wallet) and it will launch the backend automatically.
백엔드를 직접 실행하지 않습니다. [UDEX 지갑] (https://github.com/citypayorg/udex/wallet)을 실행하면 자동으로 백엔드가 실행됩니다.

# API Endpoints

## Tokens
- `GET /tokens` : returns list of all the tokens from the database
- `GET /tokens/<asset>`: returns details of a token from db using token's asset ID
- `POST /tokens`: Create/Insert token in DB. Sample input:
```
{
	"name":"HotPotCoin",
	"symbol":"HPC",
	"decimal":18,
	"asset":"0x1888a8db0b7db59413ce07150b3373972bf818d3",
	"active":true,
	"quote":true
}
```

## Pairs
- `GET /pairs` : returns list of all the pairs from the database
- `GET /pairs/<baseToken>/<quoteToken>`: returns details of a pair from db using using asset IDs of its constituting tokens
- `GET /pairs/book/<pairName>`: Returns orderbook for the pair using pair name
- `POST /pairs/create`: Create/Insert pair in DB. Sample input:
```
{
    "asset":"5b3e82587b44576ba8000001",
}

```

## Address
- `POST /account/create?address=<addr>`: Create/Insert address and corresponding balance entry in DB.

## Balance
- `GET /account/balances/<addr>`: Fetch the balance details from db of the given address.

## Order
- `GET /orders?address=<addr>`: Fetch all the orders placed by the given address

## Trade
- `GET /trades/history/<pair>`: Fetch complete trade history of given pair using pair name
- `GET /trades?address=<addr>`: Fetch all the trades in which the given address is either maker or taker
- `GET /trades/ticks`: Fetch ohlcv data. Query Params:
```
// Query Params for /trades/ticks
pairName: names of pair separated by comma(,) ex: "hpc/aut,abc/xyz". (At least 1 Required)
unit: sec,min,hour,day,week,month,yr. (default:hour)
duration: in int. (default: 24)
from: unix timestamp of from time.(default: start of timestamp)
to: unix timestamp of to time. (default: current timestamp)

// / trades / ticks에 대한 쿼리 매개 변수
pairName : 쉼표 (,)로 구분 된 쌍의 이름 예 : "hpc / aut, abc / xyz". (최소 1 개 필요)
단위 : 초, 분, 시간, 일, 주, 월, 년. (기본값 : 시간)
기간 : int. (기본값 : 24)
from : 시작 시간의 유닉스 타임 스탬프 (기본값 : 타임 스탬프 시작)
to : 시간의 유닉스 타임 스탬프. (기본값 : 현재 타임 스탬프)
```

# Types

## Orders

Orders contain the information that is required to register an order in the orderbook as a "Maker".

- **id** is the primary ID of the order (possibly deprecated)
- **orderType** is either BUY or SELL. It is currently not parsed by the server and compute directly from tokenBuy, tokenSell, amountBuy, amountSell
- **exchangeAddress** is the exchange AA address
- **maker** is the maker (usually sender) Obyte address
- **tokenBuy** is the BUY token asset ID
- **tokenSell** is the SELL token asset ID
- **amountBuy** is the BUY amount (in BUY_TOKEN units)
- **amountSell** is the SELL amount (in SELL_TOKEN units)
- **expires** is the order expiration timestamp
- **nonce** is a random string or number to make sure order hashes are unique even if all other parameters are identical
- **pairID** is a hash of the corresponding
- **hash** is a hash of the order details (see details below)
- **signature** is a signature of the order hash. The signer must equal to the maker address for the order to be valid.
- **price** corresponds to the pricepoint computed by the matching engine (not parsed)
- **amount** corresponds to the amount computed by the matching engine (not parsed)

**Order Price and Amount**

There are two ways to describe the amount of tokens being bought/sold. The AA requires (sell_asset, buy_asset, sell_amount, price) while the
orderbook requires (pairID, amount, price).

The conversion between both systems can be found in the engine.ComputeOrderPrice
function

** 주문 가격 및 금액 **

매수 / 매도되는 토큰의 양을 설명하는 두 가지 방법이 있습니다. AA에는 (sell_asset, buy_asset, sell_amount, price)가 필요하지만
주문서에는 (pairID, 금액, 가격)이 필요합니다.

두 시스템 간의 변환은 엔진에서 찾을 수 있습니다.
함수

**Order Hash**

The order hash is a sha-256 hash of the following elements:
- Exchange address
- Token Buy asset ID
- Amount Buy
- Token Sell asset ID
- Amount Sell
- Nonce
- User Address


## Trades


- **orderHash** is the hash of the matching order
- **amount** is the amount of tokens that will be traded
- **taker** is the taker Obyte address
- **pairID** is a hash identifying the token pair that will be traded
- **hash** is a unique identifier hash of the trade details (see details below)

Trade Hash:

The trade hash is a sha-256 hash of the following elements:
- Order Hash
- Amount
- Taker Address


The (Order, Trade) tuple can then be used to perform an on-chain transaction for this trade.

거래 해시 :

거래 해시는 다음 요소의 sha-256 해시입니다.
-주문 해시
-금액
-테이커 주소


그런 다음 (주문, 거래) 튜플을 사용하여이 거래를위한 온 체인 거래를 수행 할 수 있습니다.

## Quote Tokens and Token Pairs

In the same way as traditional exchanges function with the idea of base
currencies and quote currencies, the UDEX decentralized exchange works with
base tokens and quote tokens under the following principles:

- Only the exchange operator can register a quote token
- Anybody can register a token pair (but the quote token needs to be registered)

Token pairs are identified by an ID (a hash of both token asset IDs)


## 견적 토큰 및 토큰 쌍

전통적인 거래소와 같은 방식으로 기지의 아이디어로 기능
통화 및 견적 통화, UDEX 탈 중앙화 거래소는
다음 원칙에 따라 기본 토큰 및 견적 토큰 :

-거래소 운영자 만 따옴표 토큰을 등록 할 수 있습니다.
-누구나 토큰 쌍을 등록 할 수 있습니다 (단, 견적 토큰은 등록해야 함).

토큰 쌍은 ID (두 토큰 자산 ID의 해시)로 식별됩니다.

# Websocket API

See [WEBSOCKET_API.md](WEBSOCKET_API.md)


# Contribution

Thank you for considering helping the UDEX project !

To make the UDEX project truely revolutionary, we need and accept contributions from anyone and are grateful even for the smallest fixes.

If you want to help UDEX, please fork and setup the development environment of the appropriate repository. In the case you want to submit substantial changes, please get in touch with our development team on #udex channel on [Obyte Discord](https://discord.obyte.org/) to verify those modifications are in line with the general goal of the project and receive early feedback. Otherwise you are welcome to fix, commit and send a pull request for the maintainers to review and merge into the main code base.

Please make sure your contributions adhere to our coding guidelines:

Code must adhere as much as possible to standard conventions (DRY - Separation of concerns - Modular)
Pull requests need to be based and opened against the master branch
Commit messages should properly describe the code modified
Ensure all tests are passing before submitting a pull request
# 기여

UDEX 프로젝트를 도와 주셔서 감사합니다!

UDEX 프로젝트를 진정으로 혁명적으로 만들기 위해, 우리는 누구의 기여도 필요하고 수용하며 아주 작은 수정에도 감사합니다.

UDEX를 돕고 싶다면 적절한 저장소의 개발 환경을 포크하고 설정하십시오. 실질적인 변경 사항을 제출하려면 [Obyte Discord] (https://discord.obyte.org/)의 #udex 채널에서 개발팀에 문의하여 변경 사항이 일반적인 목표에 부합하는지 확인하세요. 프로젝트의 초기 피드백을받습니다. 그렇지 않으면 유지 관리자가 메인 코드베이스를 검토하고 병합 할 수 있도록 수정, 커밋 및 풀 요청을 보낼 수 있습니다.

귀하의 기여가 당사의 코딩 지침을 준수하는지 확인하십시오.

코드는 가능한 한 표준 규칙을 준수해야합니다 (DRY-우려 사항 분리-모듈 식).
풀 리퀘스트는 마스터 브랜치에 기반하고 열어야합니다.
커밋 메시지는 수정 된 코드를 올바르게 설명해야합니다.
pull 요청을 제출하기 전에 모든 테스트가 통과되었는지 확인하십시오.

# Contact

If you have questions, ideas or suggestions, you can reach the development team on Discord in the #udex channel.  [Discord Link](https://discord.obyte.org/)

## Credits

UDEX backend is based on [AMP Exchange](https://github.com/Proofsuite/amp-matching-engine), the most beautiful and easy to use decentralized exchange.

# License

All the code in this repository is licensed under the MIT License, also included in our repository in the LICENSE file.
