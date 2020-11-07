# Flares

**Flares** is a blockchain application built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

- Flares is designed as a payment network.  
- It supports multiple assets and multiple payment scenarios, such as games, dapps, DEX, DeFi, etc.  
- It is intended to discover the value of all crypto assets.  
- It's an IBC network.  
- It's a no-tokens-binding blockchain.  
- It's a value consensus.  
- It's a multi to multi contract or module.  

## Introduction

The goal of my project is to build a payment network that supports multiple assets and decentralized services.  

The first reason is:  
I like games, and games will definitely have a great development in the blockchain. However, if the game depends on a certain kind of encryption asset, it will inevitably increase the operation complexity of game developers and game players. I hope that the payment network can provide a more direct, simpler and more efficient way.  

Secondly:  
although encryption assets are very rich, their functions are relatively single. The non-unique binding payment network provides more abundant usage scenarios for each kind of encrypted assets through the open decentralized evaluation mechanism.  

Thirdly:  
all decentralized assets should have the property of payment and liquidity. Only when they have the actual liquidity and ability to pay and have fair value can they really be called assets.  

### Concept

- Contract

Contract is the core of the whole flare module. As an intermediate module for payment, secure and realize the value consensus of multiple crypto assets through services such as DEC and DeFi. It is responsible for identifying and managing the transactions of users and evaluating the value of trading assets. The contract will determine whether the transaction has been completed and whether an exchange is required and the final payment is made. There are three modes of contract.  

The first kind is simple payment.  
People transfer money to the contract, and then the contract judges whether the contract conditions are met, that is, to complete the transaction, and finally transfer the money to the payee, that is, the contract creator.  

The second one is auction.  
The contract ends at the specified block height or duration height. Then it will judge whether the transaction is successful or not and determine the maximum amount. (to query the evaluated value of each transaction means voucher, but do not generate voucher.) If the transaction meets the contract conditions, the successful transaction amount (token) will be transferred to the contract creator, and the failed transaction will be returned.  

Thirdly, game.  
The contract execution block period is determined by the clearing component defined by the service. Since the game players will need to further operate their own vouchers, when the contract receives the transfer transaction from the player, the voucher of corresponding value will be automatically cast and sent to the player. The player also holds the corresponding certificate in the game. Players may have different numbers of credentials in the same game. But the unit value of the voucher is the same. When the clearing component call is completed and successful, the certificate of the contract casting (Mint) will be immediately recovered and destroyed (burn).  

At the same time, the execution of the contract is two-tier. The flares module has a common standard contract. Some basic parameters can be configured in the flare module. For example, contract mode, effective period, etc.  

When facing different services, some customized or extended functions may be needed. Therefore, services (such as Nameservice and sealed monsters) also need to define their own clearing methods. During contract execution, the standard contract will call the service registered service custom clearing component according to the configuration.  

- Voucher

Digital credentials for encrypted assets after value consensus.  

Voucher is issued temporarily in order to ensure that the payment network does not depend on a specific currency (token), and to avoid the value fluctuation of assets in long-term transactions (such as games). The voucher exists only during the payment process. Different assets will generate different amount of vouchers according to the value evaluation mechanism. When the transaction is completed, the voucher is destroyed immediately. The unified name denom is used in the project. The algorithm needs to be further improved before testing the network line. And add ID or hash to ensure that the voucher generated by different transactions will not produce confusion. 

- Board

Provides queries and exchange interfaces for encrypted assets.  

Board is a component of the contract used to evaluate the value of assets. At present, only the ratio of base token assets and accept token assets is used to calculate the valuation value of assets. In order to ensure the accuracy and credibility of asset evaluation, it is necessary to connect the real exchangeable service before starting the test-net. For example, decentralized exchange, DEX clearing.

- Clearing

This component realizes the functions of managing and serencing encrypted assets and reallocating them according to the contract application.  

The clearing component is the second layer of the contract. It can be called service customization contract also. The liquidation component is responsible for the final liquidation of the contract completion phase. For example, in the game sealed monsters, the standard contract will send the forged voucher to the sender after receiving the send transaction. The player ,I mean transferor, holds the voucher in the game. When the game judges the end of the game, it triggers the end of the contract. The contract will call the service customized clearing component to complete the final payment. In sealed monsters, the 100th block height after the creation of monster determines the end of the game, and the liquidation method will determine who wins the game. It will return all the assets paid by the winner first. Then, all winners will distribute the bonus, the bonus means the monster reward and the assets paid by the loser, in proportion to the voucher paid by the previous winners.  

## Demos

### Sealed Monsters - flares with games

Refer to a demo project [scavenge](https://tutorials.cosmos.network/scavenge/tutorial/01-background.html) and make some changes.

![Sealed Monsters - flares with games](https://raw.githubusercontent.com/wangfeiping/flares/hackatom2020/docs/flares-game-sealed-monsters.png)

### Name Service - flares with payment or auctions

Refer to a demo project [nameservice](https://tutorials.cosmos.network/nameservice/tutorial/00-intro.html) and make some changes.

![Name Service - flares with payment or auctions](https://raw.githubusercontent.com/wangfeiping/flares/hackatom2020/docs/flares-nameservice-payment-or-auctions.png)

### Future view

![Future view](https://raw.githubusercontent.com/wangfeiping/flares/hackatom2020/docs/flares-future-view.png)

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos Tutorials](https://tutorials.cosmos.network)
- [Channel on Discord](https://discord.gg/W8trcGV)
- [BDD Test](https://github.com/onsi/ginkgo)
