// Generated by Ignite ignite.com/cli
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { Api } from "./rest";
import { MsgFundCommunityPool } from "./types/cosmos/distribution/v1beta1/tx";
import { MsgWithdrawDelegatorReward } from "./types/cosmos/distribution/v1beta1/tx";
import { MsgSetWithdrawAddress } from "./types/cosmos/distribution/v1beta1/tx";
import { MsgUpdateParams } from "./types/cosmos/distribution/v1beta1/tx";
import { MsgCommunityPoolSpend } from "./types/cosmos/distribution/v1beta1/tx";
import { MsgWithdrawValidatorCommission } from "./types/cosmos/distribution/v1beta1/tx";
import { Params as typeParams } from "./types";
import { ValidatorHistoricalRewards as typeValidatorHistoricalRewards } from "./types";
import { ValidatorCurrentRewards as typeValidatorCurrentRewards } from "./types";
import { ValidatorAccumulatedCommission as typeValidatorAccumulatedCommission } from "./types";
import { ValidatorOutstandingRewards as typeValidatorOutstandingRewards } from "./types";
import { ValidatorSlashEvent as typeValidatorSlashEvent } from "./types";
import { ValidatorSlashEvents as typeValidatorSlashEvents } from "./types";
import { FeePool as typeFeePool } from "./types";
import { CommunityPoolSpendProposal as typeCommunityPoolSpendProposal } from "./types";
import { DelegatorStartingInfo as typeDelegatorStartingInfo } from "./types";
import { DelegationDelegatorReward as typeDelegationDelegatorReward } from "./types";
import { CommunityPoolSpendProposalWithDeposit as typeCommunityPoolSpendProposalWithDeposit } from "./types";
import { DelegatorWithdrawInfo as typeDelegatorWithdrawInfo } from "./types";
import { ValidatorOutstandingRewardsRecord as typeValidatorOutstandingRewardsRecord } from "./types";
import { ValidatorAccumulatedCommissionRecord as typeValidatorAccumulatedCommissionRecord } from "./types";
import { ValidatorHistoricalRewardsRecord as typeValidatorHistoricalRewardsRecord } from "./types";
import { ValidatorCurrentRewardsRecord as typeValidatorCurrentRewardsRecord } from "./types";
import { DelegatorStartingInfoRecord as typeDelegatorStartingInfoRecord } from "./types";
import { ValidatorSlashEventRecord as typeValidatorSlashEventRecord } from "./types";
export { MsgFundCommunityPool, MsgWithdrawDelegatorReward, MsgSetWithdrawAddress, MsgUpdateParams, MsgCommunityPoolSpend, MsgWithdrawValidatorCommission };
export const registry = new Registry(msgTypes);
function getStructure(template) {
    const structure = { fields: [] };
    for (let [key, value] of Object.entries(template)) {
        let field = { name: key, type: typeof value };
        structure.fields.push(field);
    }
    return structure;
}
const defaultFee = {
    amount: [],
    gas: "200000",
};
export const txClient = ({ signer, prefix, addr } = { addr: "http://localhost:26657", prefix: "cosmos" }) => {
    return {
        async sendMsgFundCommunityPool({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgFundCommunityPool: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry, prefix });
                let msg = this.msgFundCommunityPool({ value: MsgFundCommunityPool.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgFundCommunityPool: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendMsgWithdrawDelegatorReward({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgWithdrawDelegatorReward: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry, prefix });
                let msg = this.msgWithdrawDelegatorReward({ value: MsgWithdrawDelegatorReward.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgWithdrawDelegatorReward: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendMsgSetWithdrawAddress({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgSetWithdrawAddress: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry, prefix });
                let msg = this.msgSetWithdrawAddress({ value: MsgSetWithdrawAddress.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgSetWithdrawAddress: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendMsgUpdateParams({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgUpdateParams: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry, prefix });
                let msg = this.msgUpdateParams({ value: MsgUpdateParams.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgUpdateParams: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendMsgCommunityPoolSpend({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgCommunityPoolSpend: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry, prefix });
                let msg = this.msgCommunityPoolSpend({ value: MsgCommunityPoolSpend.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgCommunityPoolSpend: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendMsgWithdrawValidatorCommission({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgWithdrawValidatorCommission: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry, prefix });
                let msg = this.msgWithdrawValidatorCommission({ value: MsgWithdrawValidatorCommission.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgWithdrawValidatorCommission: Could not broadcast Tx: ' + e.message);
            }
        },
        msgFundCommunityPool({ value }) {
            try {
                return { typeUrl: "/cosmos.distribution.v1beta1.MsgFundCommunityPool", value: MsgFundCommunityPool.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgFundCommunityPool: Could not create message: ' + e.message);
            }
        },
        msgWithdrawDelegatorReward({ value }) {
            try {
                return { typeUrl: "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward", value: MsgWithdrawDelegatorReward.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgWithdrawDelegatorReward: Could not create message: ' + e.message);
            }
        },
        msgSetWithdrawAddress({ value }) {
            try {
                return { typeUrl: "/cosmos.distribution.v1beta1.MsgSetWithdrawAddress", value: MsgSetWithdrawAddress.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgSetWithdrawAddress: Could not create message: ' + e.message);
            }
        },
        msgUpdateParams({ value }) {
            try {
                return { typeUrl: "/cosmos.distribution.v1beta1.MsgUpdateParams", value: MsgUpdateParams.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgUpdateParams: Could not create message: ' + e.message);
            }
        },
        msgCommunityPoolSpend({ value }) {
            try {
                return { typeUrl: "/cosmos.distribution.v1beta1.MsgCommunityPoolSpend", value: MsgCommunityPoolSpend.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgCommunityPoolSpend: Could not create message: ' + e.message);
            }
        },
        msgWithdrawValidatorCommission({ value }) {
            try {
                return { typeUrl: "/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission", value: MsgWithdrawValidatorCommission.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgWithdrawValidatorCommission: Could not create message: ' + e.message);
            }
        },
    };
};
export const queryClient = ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseURL: addr });
};
class SDKModule {
    query;
    tx;
    structure;
    registry = [];
    constructor(client) {
        this.query = queryClient({ addr: client.env.apiURL });
        this.updateTX(client);
        this.structure = {
            Params: getStructure(typeParams.fromPartial({})),
            ValidatorHistoricalRewards: getStructure(typeValidatorHistoricalRewards.fromPartial({})),
            ValidatorCurrentRewards: getStructure(typeValidatorCurrentRewards.fromPartial({})),
            ValidatorAccumulatedCommission: getStructure(typeValidatorAccumulatedCommission.fromPartial({})),
            ValidatorOutstandingRewards: getStructure(typeValidatorOutstandingRewards.fromPartial({})),
            ValidatorSlashEvent: getStructure(typeValidatorSlashEvent.fromPartial({})),
            ValidatorSlashEvents: getStructure(typeValidatorSlashEvents.fromPartial({})),
            FeePool: getStructure(typeFeePool.fromPartial({})),
            CommunityPoolSpendProposal: getStructure(typeCommunityPoolSpendProposal.fromPartial({})),
            DelegatorStartingInfo: getStructure(typeDelegatorStartingInfo.fromPartial({})),
            DelegationDelegatorReward: getStructure(typeDelegationDelegatorReward.fromPartial({})),
            CommunityPoolSpendProposalWithDeposit: getStructure(typeCommunityPoolSpendProposalWithDeposit.fromPartial({})),
            DelegatorWithdrawInfo: getStructure(typeDelegatorWithdrawInfo.fromPartial({})),
            ValidatorOutstandingRewardsRecord: getStructure(typeValidatorOutstandingRewardsRecord.fromPartial({})),
            ValidatorAccumulatedCommissionRecord: getStructure(typeValidatorAccumulatedCommissionRecord.fromPartial({})),
            ValidatorHistoricalRewardsRecord: getStructure(typeValidatorHistoricalRewardsRecord.fromPartial({})),
            ValidatorCurrentRewardsRecord: getStructure(typeValidatorCurrentRewardsRecord.fromPartial({})),
            DelegatorStartingInfoRecord: getStructure(typeDelegatorStartingInfoRecord.fromPartial({})),
            ValidatorSlashEventRecord: getStructure(typeValidatorSlashEventRecord.fromPartial({})),
        };
        client.on('signer-changed', (signer) => {
            this.updateTX(client);
        });
    }
    updateTX(client) {
        const methods = txClient({
            signer: client.signer,
            addr: client.env.rpcURL,
            prefix: client.env.prefix ?? "cosmos",
        });
        this.tx = methods;
        for (let m in methods) {
            this.tx[m] = methods[m].bind(this.tx);
        }
    }
}
;
const Module = (test) => {
    return {
        module: {
            CosmosDistributionV1Beta1: new SDKModule(test)
        },
        registry: msgTypes
    };
};
export default Module;
//# sourceMappingURL=module.js.map