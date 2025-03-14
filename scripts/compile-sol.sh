#! /bin/bash

solc --base-path . --include-path .. ../contracts/impl/GToken.sol --bin --abi -o ../abi/AccountDID --overwrite
