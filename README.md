# Kloud-数字资产交易平台

![](2022-06-13-20-33-20.png)

##  可行性分析报告 

### 引言

#### 编写目的

可行性分析报告是为“Kloud”开发的可能性、可行性、必要性提供论据，为开发人员进行系统总体规划设计及具体实施开发工程提供必要的参考资料，在系统开发完成后期为系统的测试、验收提供帮助。其编写过程由张云鹏完成。预期读者是从事Kloud开发的相关人员。

#### 项目背景

本项目名称为“Kloud数字资产交易平台”。系统功能主要包括：能够存储用户的数字资产, 提供主流数字货币的便捷交易, 支持对主流交易对的限价单, 市价单委托、用户还可通过API接口获取实时行情数据信息授权访问, 以及Okx和Binance的交易大数据, 本项目的任务提出者为北京交通大学软件实践工程课程，开发者为计算机学院张云鹏。

#### 参考资料
1. 张海藩.《软件工程导论》.人民邮电出版社.2006年1月
2. jammy928. Bizzan. https://github.com/bruce2233/CoinExchange_CryptoExchange_Java

### 项目概述

#### 要求

系统能够存储用户的数字资产, 提供主流数字货币的便捷交易, 支持对主流交易对的限价单, 市价单委托、用户还可通过API接口获取实时行情数据信息授权访问, 以及Okx和Binance的交易大数据

#### 功能

系统主要功能是数字资产存储, 限价委托, 市价委托交易和行情数据实时查询

#### 性能

交易系统对安全性要求极高,在不开放API的情况下, 交易全部通过UI接口, 要求较低. 但性能需求会随用户人数增加, 所以必须具备良好的可拓展性. 并发数达到10000笔/秒

#### 系统输出

- 行情数据
- 用户资产
- 账单记录

#### 系统输入

- 订单委托
- 用户资产
- 用户充值/提现

#### 处理流程和数据流程

![](docs/%E5%A4%84%E7%90%86%E6%B5%81%E7%A8%8B%E5%92%8C%E6%95%B0%E6%8D%AE%E6%B5%81%E7%A8%8B.svg)

#### 可靠性和安全性需求

由于交易平台的订单数量和频次较高, 同时涉及财产, 所以数据的必须依靠一致性较强的原子操作. 对于整个系统, 需要完整的权限控制. 尤其是充值/提现和订单委托的权限管理. 同时redis数据必须即使持久化, 便于意外恢复. 

#### 项目完成期限

本项目的完成期限为2022年6月24日. 具体进度见软件项目计划.

### 项目基本目标
1. 用户能够通过交易平台与其他用户交易数字资产.
2. 系统能存储数字资产到交易平台.
3. 系统需要较好的安全性和容灾恢复机制.

### 