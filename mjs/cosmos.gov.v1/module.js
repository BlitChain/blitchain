// Generated by Ignite ignite.com/cli
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { Api } from "./rest";
import { MsgVote } from "./types/cosmos/gov/v1/tx";
import { MsgSubmitProposal } from "./types/cosmos/gov/v1/tx";
import { MsgDeposit } from "./types/cosmos/gov/v1/tx";
import { MsgUpdateParams } from "./types/cosmos/gov/v1/tx";
import { MsgVoteWeighted } from "./types/cosmos/gov/v1/tx";
import { WeightedVoteOption as typeWeightedVoteOption } from "./types";
import { Deposit as typeDeposit } from "./types";
import { Proposal as typeProposal } from "./types";
import { TallyResult as typeTallyResult } from "./types";
import { Vote as typeVote } from "./types";
import { DepositParams as typeDepositParams } from "./types";
import { VotingParams as typeVotingParams } from "./types";
import { TallyParams as typeTallyParams } from "./types";
import { Params as typeParams } from "./types";
export { MsgVote, MsgSubmitProposal, MsgDeposit, MsgUpdateParams, MsgVoteWeighted };
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
        async sendMsgVote({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgVote: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry, prefix });
                let msg = this.msgVote({ value: MsgVote.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgVote: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendMsgSubmitProposal({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgSubmitProposal: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry, prefix });
                let msg = this.msgSubmitProposal({ value: MsgSubmitProposal.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgSubmitProposal: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendMsgDeposit({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgDeposit: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry, prefix });
                let msg = this.msgDeposit({ value: MsgDeposit.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgDeposit: Could not broadcast Tx: ' + e.message);
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
        async sendMsgVoteWeighted({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgVoteWeighted: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry, prefix });
                let msg = this.msgVoteWeighted({ value: MsgVoteWeighted.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgVoteWeighted: Could not broadcast Tx: ' + e.message);
            }
        },
        msgVote({ value }) {
            try {
                return { typeUrl: "/cosmos.gov.v1.MsgVote", value: MsgVote.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgVote: Could not create message: ' + e.message);
            }
        },
        msgSubmitProposal({ value }) {
            try {
                return { typeUrl: "/cosmos.gov.v1.MsgSubmitProposal", value: MsgSubmitProposal.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgSubmitProposal: Could not create message: ' + e.message);
            }
        },
        msgDeposit({ value }) {
            try {
                return { typeUrl: "/cosmos.gov.v1.MsgDeposit", value: MsgDeposit.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgDeposit: Could not create message: ' + e.message);
            }
        },
        msgUpdateParams({ value }) {
            try {
                return { typeUrl: "/cosmos.gov.v1.MsgUpdateParams", value: MsgUpdateParams.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgUpdateParams: Could not create message: ' + e.message);
            }
        },
        msgVoteWeighted({ value }) {
            try {
                return { typeUrl: "/cosmos.gov.v1.MsgVoteWeighted", value: MsgVoteWeighted.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgVoteWeighted: Could not create message: ' + e.message);
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
            WeightedVoteOption: getStructure(typeWeightedVoteOption.fromPartial({})),
            Deposit: getStructure(typeDeposit.fromPartial({})),
            Proposal: getStructure(typeProposal.fromPartial({})),
            TallyResult: getStructure(typeTallyResult.fromPartial({})),
            Vote: getStructure(typeVote.fromPartial({})),
            DepositParams: getStructure(typeDepositParams.fromPartial({})),
            VotingParams: getStructure(typeVotingParams.fromPartial({})),
            TallyParams: getStructure(typeTallyParams.fromPartial({})),
            Params: getStructure(typeParams.fromPartial({})),
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
            CosmosGovV1: new SDKModule(test)
        },
        registry: msgTypes
    };
};
export default Module;
//# sourceMappingURL=module.js.map