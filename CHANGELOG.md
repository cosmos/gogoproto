# Changelog

## [Unreleased]

## [v1.4.3](https://github.com/cosmos/gogoproto/releases/tag/v1.4.2) - 2022-10-14

### Bug Fixes

- [#24](https://github.com/cosmos/gogoproto/pull/24) Fix `CompactTextString` panics with nested Anys and private fields.
- [#14](https://github.com/cosmos/gogoproto/pull/14) Fix `make regenerate`.

## [v1.4.2](https://github.com/cosmos/gogoproto/releases/tag/v1.4.2) - 2022-09-14

### Features

- [#13](https://github.com/cosmos/gogoproto/pull/13) Add `AllFileDescriptors` function.

### Improvements

- [#8](https://github.com/cosmos/gogoproto/pull/8) Fix typo in `doc.go`.
- [#8](https://github.com/cosmos/gogoproto/pull/8) Support for merging messages implementing Merger which are embedded by value.
- [#8](https://github.com/cosmos/gogoproto/pull/8) Use reflect.Value.String() for String kinds in proto equal.

## [v1.4.1](https://github.com/cosmos/gogoproto/releases/tag/v1.4.1) - 2022-08-30

### Improvements

- [#6](https://github.com/cosmos/gogoproto/pull/6) Add buf.yaml for cosmos/gogo-proto module.

### Bug Fixes

- [226206f](https://github.com/cosmos/gogoproto/commit/226206f39bd7276e88ec684ea0028c18ec2c91ae) Fixed order of imports, make stable generation result.

## [v1.4.0](https://github.com/cosmos/gogoproto/releases/tag/v1.4.0) - 2022-03-18

- Migration from [regen-network/protobuf](https://github.com/regen-network/protobuf), a fork of [gogo/protobuf](https://github.com/gogo/protobuf) used by the Cosmos SDK to [cosmos/gogoproto](https://github.com/cosmos/gogoproto) (this repository).
