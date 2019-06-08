// Copyright 2019 The Energi Core Authors
// This file is part of Energi Core.
//
// Energi Core is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Energi Core is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Energi Core. If not, see <http://www.gnu.org/licenses/>.

// Energi Governance system is the fundamental part of Energi Core.

// NOTE: It's not allowed to change the compiler due to byte-to-byte
//       match requirement.
pragma solidity 0.5.9;
//pragma experimental SMTChecker;

import { GlobalConstants } from "./constants.sol";
import { IGovernedContract, GovernedContract } from "./GovernedContract.sol";
import { IBlockReward } from "./IBlockReward.sol";

/**
 * Genesis hardcoded version of BackboneReward
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract BackboneRewardV1 is
    GlobalConstants,
    GovernedContract,
    IBlockReward
{
    constructor(address _proxy) public GovernedContract(_proxy) {}
    function migrate(IGovernedContract) external requireProxy {}
    function destroy(IGovernedContract) external requireProxy {}
    function () external payable {}

    function reward(uint _amount) external payable {
    }

    function getReward(uint _blockNumber)
        external view
        returns(uint amount)
    {
        if (_blockNumber > 0) {
            amount = REWARD_BACKBONE_V1;
        }
    }
}

