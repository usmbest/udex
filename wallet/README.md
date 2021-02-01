# UDEX Decentralized Exchange

The UDEX decentralized exchange is a hybrid decentralized exchange that aims at bringing together the ease of use of centralized exchanges along with the security and privacy features of decentralized exchanges. Orders are matched through an off-chain orderbook. After orders are matched and signed, the decentralized exchange operator (matcher) has the sole ability to perform a transaction to the Autonomous Agent. This provides for the best UX as the exchange operator is the only party having to interact directly with the DAG. Exchange users simply sign orders which are broadcasted then to the orderbook. This design enables users to queue and cancel their orders seamlessly.

Several matchers can operate exchanges based on UDEX technology at the same time. They share their orderbooks and exchange all new orders among themselves, thus improving liquidity for all UDEX exchanges. An order can be submitted through any UDEX exchange, however to be matched, both maker and taker orders have to indicate the same matcher. The exchange that was used to submit the order serves as an affliate and can charge a fee from its users.  Anyone can become a matcher or affiliate, or just watch the orders that are being exchanged among the matchers and detect any possible misbehavior by matchers.

# UDEX 탈 중앙화 거래소

UDEX 분산 거래소는 분산 거래소의 보안 및 개인 정보 보호 기능과 함께 중앙 거래소의 사용 편의성을 결합하는 것을 목표로하는 하이브리드 탈 중앙 거래소입니다. 주문은 오프 체인 주문서를 통해 매칭됩니다. 주문이 일치되고 서명 된 후 탈 중앙화 거래소 운영자 (매처)는 Autonomous Agent에 대한 거래를 수행 할 수있는 유일한 능력을 갖습니다. 이는 교환 운영자가 DAG와 직접 상호 작용해야하는 유일한 당사자이기 때문에 최상의 UX를 제공합니다. 교환 사용자는 주문서에 브로드 캐스트 된 주문에 서명하기 만하면됩니다. 이 디자인을 통해 사용자는 주문을 원활하게 대기열에 추가하고 취소 할 수 있습니다.

여러 매 처가 동시에 UDEX 기술을 기반으로 거래소를 운영 할 수 있습니다. 그들은 주문서를 공유하고 모든 새로운 주문을 서로 교환하여 모든 UDEX 거래소의 유동성을 향상시킵니다. 주문은 모든 UDEX 거래소를 통해 제출할 수 있지만 일치하려면 메이커 주문과 테이커 주문 모두 동일한 일치자를 표시해야합니다. 주문을 제출하는 데 사용 된 거래소는 제휴사 역할을하며 사용자에게 수수료를 부과 할 수 있습니다. 누구든지 매처 또는 제휴사가 될 수 있으며, 매처간에 교환되는 주문을보고 매처의 가능한 오작동을 감지 할 수 있습니다.

## Install
```
git clone https://github.com/citypayorg/udex/wallet
cd udex-wallet
npm install
```
Copy `.env.testnet` file to `.env` if you are working on testnet.

## Run
```
node run.js
```
Run this from `screen` or `tmux` session and detach after starting, or start it in the background while redirecting all output:
`screen` 또는`tmux` 세션에서 실행하고 시작 후 분리하거나 모든 출력을 리디렉션하는 동안 백그라운드에서 시작합니다.
```
node run.js 1>log 2>errlog &
```

## Architecture
The DEX consists of 4 layers:
* An Autonomous Agent, written in [Oscript](https://developer.obyte.org/autonomous-agents), that tracks user balances and executes trades.
* DEX wallet, written in nodejs (this repo). It includes an Obyte node and is reponsible for:
	* sending matched trades to the AA foe execution;
	* watching the DAG for deposits/withdrawals and sending events to the backend (which in turn forwards them to clients);
	* receiving orders and cancels from users through a chatbot;
	* logging in users through a chatbot;
	* exchanging new orders, cancels, and trades with other matchers (USMDEXes) through websocket connections.
* [DEX backend](https://github.com/citypayorg/udex/backend), written in go. It is responsible for:
	* matching orders;
	* serving REST and websocket endponts for frontend and bot clients;
	* forwarding events from DEX wallet to clients (browsers and bots).
* [DEX frontend](https://github.com/citypayorg/udex/frontend), written in react. It is the UI users use to interact with the exchange.

Orders, trades, and other information is stored in a mongodb database. It is most actively used by UDEX backend but UDEX wallet also has access to it.

Backend and wallet interact through JSON-RPC and websocket connections.
## 건축물
DEX는 4 개의 레이어로 구성됩니다.
* 사용자 잔액을 추적하고 거래를 실행하는 [Oscript] (https://developer.obyte.org/autonomous-agents)로 작성된 Autonomous Agent.
* nodejs (이 저장소)로 작성된 DEX 지갑. 여기에는 Obyte 노드가 포함되어 있으며 다음을 담당합니다.
* 일치하는 거래를 AA 적 처형에 보내기;
* 입금 / 출금을 위해 DAG를 감시하고 이벤트를 백엔드로 전송 (이를 통해 클라이언트에게 전달)
* 챗봇을 통해 사용자의 주문 및 취소를 수신합니다.
* 챗봇을 통해 사용자 로그인
* 웹 소켓 연결을 통해 다른 매처 (UDEX)와 새로운 주문, 취소 및 거래를 교환합니다.
* [DEX 백엔드] (https://github.com/citypayorg/udex/backend), go로 작성되었습니다. 다음을 담당합니다.
* 일치하는 주문;
* 프론트 엔드 및 봇 클라이언트를위한 REST 및 웹 소켓 엔드 폰트 제공
* DEX 지갑에서 클라이언트 (브라우저 및 봇)로 이벤트 전달.
* [DEX frontend] (https://github.com/citypayorg/udex/frontend), react로 작성되었습니다. 사용자가 거래소와 상호 작용하는 데 사용하는 UI입니다.

주문, 거래 및 기타 정보는 mongodb 데이터베이스에 저장됩니다. UDEX 백엔드에서 가장 많이 사용되지만 UDEX 지갑도 액세스 할 수 있습니다.

백엔드와 지갑은 JSON-RPC 및 웹 소켓 연결을 통해 상호 작용합니다.

## How to run your own UDEX node

* Install all 3 repos: wallet, backend, and frontend (the AA is a public service and is shared among all UDEX instances in order to share liquidity).
* Configure your node:
	* decide which network you run, livenet or testnet, and copy the corresponding .env.XXX file to .env in wallet and frontend;
	* find all mentions of usm.best domain name in the frontend and replace them with your domain name;
	* set up nginx according to example config file in the frontend repo (make sure you edit the domain name and document root path);
	* edit the chatbot name (`deviceName`) in conf.js or conf.json of your wallet. Make sure it reflects your donmain name so that users don't confuse it with chatbots of other USMDEXes.
	* add `admin_email` and `from_email` in conf.js or conf.json of your wallet. This is the address where you'll receive notifications in case of issues with your wallet, such as insufficient balance to send a trade to the AA. Set up `sendmail` or another way to deliver email (such as user/pass for an email account).
	* decide whether you run full or light node. Light node is much faster to start and takes much less disk space but it is slower to see the finality of transactions and you might run into issues as your exchange node grows. Migration from light to full is easy should you need it. Set the corresponding `bLight` option in conf.js or conf.json of your wallet.
	* when you first start the wallet, note the pairing code that it prints, copy it to `CHATBOT_TESTNET_URL` or `CHATBOT_LIVENET_URL` (depending on which network you run) in `src/config/urls.js` of the frontend.
	* when you first start the wallet, note the address that it prints. Send some Bytes (0.1 GBYTE is recommended) to this address for your node to be able to pay for fees when it sends trades for execution to the AA.
	* set the `myUrl` and `port` in conf.js or conf.json of your wallet. Other UDEX nodes will connect to your `myUrl` to exchange information about new orders, cancels, and trades. `port` is for your nginx to proxy websocket connections to your node. Set up nginx to accept websocket connections on `myUrl` and proxy them to localhost:port.
	* if you run a full node, it will automatically discover other UDEX nodes and connect to them. If you run a light node, you might want to help it discover other ODExes by specifying the URLs of their Obyte nodes in `light_peers` of the wallet (see an example in conf.js).
	* optionally set up TOR for security (attackers won't know your IP) and privacy. Specify `socksHost` and `socksPort` in conf.js or conf.json of your wallet. Note that by running a web server you are exposing your IP, use a reverse proxy such as cloudflare to keep it private.
	* review other options you might want to edit in wallet's conf.js. In particular, check out `matcher_fee`, `affiliate_fee`, `MIN_BALANCE_FOR_REFILL`, `MIN_BALANCE_FOR_NOTIFICATION`.
* follow the instructions in frontend and wallet repos to start them.

## 자신의 UDEX 노드를 실행하는 방법

* 지갑, 백엔드, 프런트 엔드의 3 가지 저장소를 모두 설치합니다 (AA는 공공 서비스이며 유동성을 공유하기 위해 모든 UDEX 인스턴스간에 공유됩니다).
* 노드 구성 :
* 실행중인 네트워크, 라이브 넷 또는 테스트 넷을 결정하고 해당 .env.XXX 파일을 지갑 및 프런트 엔드의 .env에 복사합니다.
* 프론트 엔드에서 usm.best 도메인 이름에 대한 모든 언급을 찾아 귀하의 도메인 이름으로 대체합니다.
* 프런트 엔드 저장소의 예제 구성 파일에 따라 nginx를 설정합니다 (도메인 이름과 문서 루트 경로를 편집해야합니다).
* 지갑의 conf.js 또는 conf.json에서 챗봇 이름 (`deviceName`)을 편집하십시오. 사용자가 다른 USMDEX의 챗봇과 혼동하지 않도록 donmain 이름을 반영해야합니다.
* 지갑의 conf.js 또는 conf.json에`admin_email` 및`from_email`을 추가하십시오. 이것은 AA로 거래를 보내기에 부족한 잔액과 같은 지갑에 문제가있는 경우 알림을받을 주소입니다. 'sendmail'또는 이메일을 전달하는 다른 방법 (예 : 이메일 계정의 사용자 / 패스)을 설정합니다.
* 전체 또는 가벼운 노드 실행 여부를 결정하십시오. 라이트 노드는 시작하는 데 훨씬 빠르며 디스크 공간을 훨씬 적게 차지하지만 트랜잭션의 최종성을 확인하는 속도가 느리고 교환 노드가 커짐에 따라 문제가 발생할 수 있습니다. 필요한 경우 라이트에서 전체로 쉽게 마이그레이션 할 수 있습니다. 지갑의 conf.js 또는 conf.json에서 해당`bLight` 옵션을 설정하십시오.
* 지갑을 처음 시작할 때 인쇄되는 페어링 코드를 확인하고 프런트 엔드의`src / config / urls.js`에있는`CHATBOT_TESTNET_URL` 또는`CHATBOT_LIVENET_URL` (실행하는 네트워크에 따라 다름)에 복사합니다.
* 지갑을 처음 시작할 때 인쇄 된 주소를 기록해 두십시오. 노드가 거래를 AA로 보낼 때 수수료를 지불 할 수 있도록이 주소로 약간의 바이트 (0.1 GBYTE 권장)를 보냅니다.

* 지갑의 conf.js 또는 conf.json에서`myUrl`과`port`를 설정하십시오. 다른 UDEX 노드는 'myUrl'에 연결하여 새로운 주문, 취소 및 거래에 대한 정보를 교환합니다. `port`는 nginx에서 노드에 대한 websocket 연결을 프록시하는 것입니다. 'myUrl'에서 웹 소켓 연결을 허용하도록 nginx를 설정하고이를 localhost : port에 프록시합니다.
* 전체 노드를 실행하면 다른 UDEX 노드를 자동으로 검색하여 연결합니다. 라이트 노드를 실행하는 경우 지갑의 'light_peers'에 Obyte 노드의 URL을 지정하여 다른 ODEx를 발견하는 데 도움이 될 수 있습니다 (conf.js의 예 참조).
* 선택적으로 보안 및 개인 정보 보호를 위해 TOR를 설정합니다. 지갑의 conf.js 또는 conf.json에`socksHost` 및`socksPort`를 지정하십시오. 웹 서버를 실행하여 IP를 노출하고 Cloudflare와 같은 역방향 프록시를 사용하여 비공개로 유지합니다.
* 지갑의 conf.js에서 편집 할 수있는 다른 옵션을 검토하십시오. 특히`matcher_fee`,`affiliate_fee`,`MIN_BALANCE_FOR_REFILL`,`MIN_BALANCE_FOR_NOTIFICATION`을 확인하세요.
* 프런트 엔드 및 지갑 리포지토리의 지침에 따라 시작하십시오.


# Contributions

Thank you for considering helping the UDEX project! We accept contributions from anyone and are grateful even for the smallest fixes.

If you want to help UDEX, please fork and setup the development environment of the appropriate repository. In the case you want to submit substantial changes, please get in touch with our development team on #udex channel in [Obyte discord](https://discord.obyte.org) to verify those modifications are in line with the general goal of the project and receive early feedback. Otherwise you are welcome to fix, commit and send a pull request for the maintainers to review and merge into the main code base.

Please make sure your contributions adhere to our coding guidelines:

Code must adhere as much as possible to standard conventions (DRY - Separation of concerns - Modular)
Pull requests need to be based and opened against the master branch
Commit messages should properly describe the code modified
Ensure all tests are passing before submitting a pull request

# 기여

UDEX 프로젝트를 도와 주셔서 감사합니다! 우리는 누구의 기여도 받아들이고 아주 작은 수정에도 감사합니다.

USMDEX를 돕고 싶다면 적절한 저장소의 개발 환경을 포크하고 설정하십시오. 실질적인 변경 사항을 제출하려면 [Obyte discord] (https://discord.obyte.org)의 #udex 채널 개발팀에 연락하여 변경 사항이 일반적인 목표에 부합하는지 확인하시기 바랍니다. 프로젝트를 시작하고 초기 피드백을받습니다. 그렇지 않으면 유지 관리자가 메인 코드베이스를 검토하고 병합 할 수 있도록 수정, 커밋 및 풀 요청을 보낼 수 있습니다.

귀하의 기여가 당사의 코딩 지침을 준수하는지 확인하십시오.

코드는 가능한 한 표준 규칙을 준수해야합니다 (DRY-우려 사항 분리-모듈 식).
풀 리퀘스트는 마스터 브랜치에 기반하고 열어야합니다.
커밋 메시지는 수정 된 코드를 올바르게 설명해야합니다.
pull 요청을 제출하기 전에 모든 테스트가 통과되었는지 확인하십시오.

# Contact

If you have questions, ideas or suggestions, you can reach our development team on #udex channel in [Obyte discord](https://discord.obyte.org)

# License

All the code in this repository is licensed under the MIT License, also included here in the LICENSE file.


# 연락처

질문, 아이디어 또는 제안이 있으면 [Obyte discord] (https://discord.obyte.org)의 #udex 채널에서 개발팀에 문의 할 수 있습니다.

# 라이선스

이 저장소의 모든 코드는 여기 LICENSE 파일에도 포함 된 MIT 라이선스에 따라 라이선스가 부여됩니다.
