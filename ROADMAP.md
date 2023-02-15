# ROADMAP

This document contains the roadmap for the Juno project. It is a living document and will be updated as the project progresses. For the most update to date information, please follow the Notion and tracking issue links in each section.

---

## V14 - Q3 or Q4 2023 / TBD
<!-- - [Medium Blog](https://medium.com/@reecepbcups/juno-v12-update-4bab64640a62) -->

- [Notion Plan Page](https://fluffy-conifer-309.notion.site/123261ebfe2040d9ac559f7e7d3c5cd2?v=6d59a04765f543738676f8db21ae8525)

- [V13 Tracking Issue](https://github.com/EZStaking/baobab/issues/475)

This update will focus more on upgrading the base layer of the Juno stack, bringing new features and pushing us to the latest versions of the software.

### Features Planned

- SDK v0.47, Tendermint 0.37
- IBC v5/6, ICA v3 (optional)
- Using [Skip's MEV Tendermint fork](https://github.com/skip-mev/mev-tendermint) by default
- [Native Liquid Staking](https://github.com/iqlusioninc/liquidity-staking-module)
- IBCTest (optional)
- Faster Block times (with respect to oracle)
- [Improving the Nakamoto Coefficient](https://github.com/EZStaking/baobab/issues/474)

---

## V13 - Q1 2023

Links:

- [Medium Blog](https://medium.com/@reecepbcups/juno-v12-update-4bab64640a62)

- [Notion Plan Page](https://fluffy-conifer-309.notion.site/123261ebfe2040d9ac559f7e7d3c5cd2?v=6d59a04765f543738676f8db21ae8525)

- [3 Tracking Issue](https://github.com/EZStaking/baobab/issues/268)

The V13 update is Juno's largest update, bringing many new features for developers, users, and relayers.

### V13 PRs

- [x/FeeShare (CosmWasm)](https://github.com/EZStaking/baobab/pull/385)
- [x/TokenFactory](https://github.com/EZStaking/baobab/pull/368)
- [x/Oracle](https://github.com/EZStaking/baobab/pull/329)
- [x/GlobalFee](https://github.com/EZStaking/baobab/pull/411)
- [x/inter-tx](https://github.com/EZStaking/baobab/pull/215)
- [More ICA Messages](https://github.com/EZStaking/baobab/pull/436/files)
- [Governance Spam Prevention](https://github.com/EZStaking/baobab/pull/394)
- [x/wasmd 30](https://github.com/EZStaking/baobab/pull/387)
- [x/ibc V4](https://github.com/EZStaking/baobab/pull/387)
- [x/ibc-fees](https://github.com/EZStaking/baobab/pull/432)

V13 is targeted at developers with relayer and user experience improvements as well.

**FeeShare** will allow contract developers to receive 50% of gas fees executed on their contract. Providing an alternative income source for new business use cases. This also enhances current business models to support developers & grow the ecosystem further.

The **TokenFactory** will make developers' lives easier, and also make querying users' [DAO](https://daodao.zone/) tokens via MintScan and Keplr possible. By default, CosmWasm smart contracts accept native tokens. However, the only initial native tokens are the staking demons for most chains. This gives the ability for a user to create their token, and manage the tokenomics behind it. Then accept it just as they would any other denomination via the standard [x/bank](https://github.com/cosmos/cosmos-sdk/tree/main/x/bank) module.

The **oracle** brings the ability to query external data sources, such as the price of JUNO or ATOM. This is a powerful tool for De-Fi applications on Juno like trading games, perpetual swaps, and more.

Governance can now also deny and allow IBC-based denominations (tokens) via **GlobalFee**, bringing massive User Experience improvements with it. First, all nodes will be required to accept the minimum fee. Given this, a user can query the required fee for gas, improving front-end UX. Second, it reduces operational costs for Validators and Relayers by whitelisting select IBC and ORACLE transactions. Now, these transactions will be free and take the strain off all parties.

Speaking of relayers, **IBCFees** now helps to fund those who relayer your packets! In the above paragraph, we mention how IBC transfers are feeless for relayers. Fees can still be sent with these packets and bring some income for relayers, thus maintaining public goods infrastructure. The relayers still have to pay the fee on the other chain's token, but this is a positive step in the right direction for variables we can control.

Juno is now a controller chain via **inter-tx**, which will allow it to control accounts on other chains. [HERE](https://github.com/EZStaking/baobab/pull/436/files).
