# UDEX decentralized exchange front end

The UDEX decentralized exchange is a hybrid decentralized exchange that aims at bringing together the ease of use of centralized exchanges along with the security and privacy features of decentralized exchanges. Orders are matched through an off-chain orderbook. After orders are matched and signed, the decentralized exchange operator (matcher) has the sole ability to perform a transaction to the Autonomous Agent. This provides for the best UX as the exchange operator is the only party having to interact directly with the DAG. Exchange users simply sign orders which are broadcasted then to the orderbook. This design enables users to queue and cancel their orders seamlessly.

Several matchers can operate exchanges based on UDEX technology at the same time. They share their orderbooks and exchange all new orders among themselves, thus improving liquidity for all UDEX exchanges. An order can be submitted through any UDEX exchange, however to be matched, both maker and taker orders have to indicate the same matcher. The exchange that was used to submit the order serves as an affliate and can charge a fee from its users.  Anyone can become a matcher or affiliate, or just watch the orders that are being exchanged among the matchers and detect any possible misbehavior by matchers.

# UDEX 탈 중앙화 거래소 프런트 엔드

UDEX 분산 거래소는 분산 거래소의 보안 및 개인 정보 보호 기능과 함께 중앙 거래소의 사용 편의성을 결합하는 것을 목표로하는 하이브리드 탈 중앙 거래소입니다. 주문은 오프 체인 주문서를 통해 매칭됩니다. 주문이 일치되고 서명 된 후 탈 중앙화 거래소 운영자 (매처)는 Autonomous Agent에 대한 거래를 수행 할 수있는 유일한 능력을 갖습니다. 이는 교환 운영자가 DAG와 직접 상호 작용해야하는 유일한 당사자이기 때문에 최상의 UX를 제공합니다. 교환 사용자는 주문서에 브로드 캐스트 된 주문에 서명하기 만하면됩니다. 이 디자인을 통해 사용자는 주문을 원활하게 대기열에 추가하고 취소 할 수 있습니다.

여러 명의 매 처가 동시에 UDEX 기술을 기반으로 교환을 운영 할 수 있습니다. 그들은 주문서를 공유하고 모든 새로운 주문을 서로 교환하여 모든 UDEX 거래소의 유동성을 향상시킵니다. 주문은 UDEX 거래소를 통해 제출할 수 있지만 일치하려면 메이커 주문과 테이커 주문 모두 동일한 일치자를 표시해야합니다. 주문을 제출하는 데 사용 된 거래소는 제휴사 역할을하며 사용자에게 수수료를 부과 할 수 있습니다. 누구든지 매처 또는 제휴사가 될 수 있으며, 매처간에 교환되는 주문을보고 매처의 가능한 오작동을 감지 할 수 있습니다.

# Installation
```
git clone https://github.com/citypayorg/udex/frontend
cd udex-frontend
yarn
```
Copy the appropriate `.env.XXXX` file to `.env`. Edit if necessary to reflect your domain name.

To start the frontend in development environment:
```
yarn start
```
To deploy a production (livenet or testnet) frontend:
```
yarn sass
yarn build
```
This builds all static files and puts them in `build` folder. Set this folder as web root in your nginx. See [an example of nginx vertual host config](blob/develop/nginx.conf).
이렇게하면 모든 정적 파일이 빌드되고`build` 폴더에 저장됩니다. 이 폴더를 nginx의 웹 루트로 설정하십시오. [nginx 수직 호스트 구성의 예] (blob / develop / nginx.conf)를 참조하십시오.

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

UDEX를 돕고 싶다면 적절한 저장소의 개발 환경을 포크하고 설정하십시오. 실질적인 변경 사항을 제출하려면 [Obyte discord] (https://discord.obyte.org)의 #udex 채널에서 개발팀에 연락하여 변경 사항이 일반적인 목표에 부합하는지 확인하시기 바랍니다. 프로젝트를 시작하고 초기 피드백을받습니다. 그렇지 않으면 유지 관리자가 메인 코드베이스를 검토하고 병합 할 수 있도록 수정, 커밋 및 풀 요청을 보낼 수 있습니다.

귀하의 기여가 당사의 코딩 지침을 준수하는지 확인하십시오.

코드는 가능한 한 표준 규칙을 준수해야합니다 (DRY-우려 사항 분리-모듈 식).
풀 리퀘스트는 마스터 브랜치에 기반하고 열어야합니다.
커밋 메시지는 수정 된 코드를 올바르게 설명해야합니다.
pull 요청을 제출하기 전에 모든 테스트가 통과되었는지 확인하십시오.

# Contact

If you have questions, ideas or suggestions, you can reach our development team on #udex channel in [Obyte discord](https://discord.obyte.org)

## Credits

UDEX frontend is based on [AMP Exchange](https://github.com/Proofsuite/amp-client), the most beautiful and easy to use decentralized exchange.

# License

All the code in this repository is licensed under the MIT License, also included here in the LICENSE file.


# 연락처

질문, 아이디어 또는 제안이 있으면 [Obyte discord] (https://discord.obyte.org)의 #udex 채널에서 개발팀에 문의 할 수 있습니다.

## 크레딧

UDEX 프런트 엔드는 가장 아름답고 사용하기 쉬운 분산 형 거래소 인 [AMP Exchange] (https://github.com/Proofsuite/amp-client)를 기반으로합니다.

# 라이선스

이 저장소의 모든 코드는 여기 LICENSE 파일에도 포함 된 MIT 라이선스에 따라 라이선스가 부여됩니다.