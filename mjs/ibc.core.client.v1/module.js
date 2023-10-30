// Generated by Ignite ignite.com/cli
import { Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { Api } from "./rest";
import { IdentifiedClientState as typeIdentifiedClientState } from "./types";
import { ConsensusStateWithHeight as typeConsensusStateWithHeight } from "./types";
import { ClientConsensusStates as typeClientConsensusStates } from "./types";
import { ClientUpdateProposal as typeClientUpdateProposal } from "./types";
import { UpgradeProposal as typeUpgradeProposal } from "./types";
import { Height as typeHeight } from "./types";
import { Params as typeParams } from "./types";
import { GenesisMetadata as typeGenesisMetadata } from "./types";
import { IdentifiedGenesisMetadata as typeIdentifiedGenesisMetadata } from "./types";
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
    return {};
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
            IdentifiedClientState: getStructure(typeIdentifiedClientState.fromPartial({})),
            ConsensusStateWithHeight: getStructure(typeConsensusStateWithHeight.fromPartial({})),
            ClientConsensusStates: getStructure(typeClientConsensusStates.fromPartial({})),
            ClientUpdateProposal: getStructure(typeClientUpdateProposal.fromPartial({})),
            UpgradeProposal: getStructure(typeUpgradeProposal.fromPartial({})),
            Height: getStructure(typeHeight.fromPartial({})),
            Params: getStructure(typeParams.fromPartial({})),
            GenesisMetadata: getStructure(typeGenesisMetadata.fromPartial({})),
            IdentifiedGenesisMetadata: getStructure(typeIdentifiedGenesisMetadata.fromPartial({})),
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
            IbcCoreClientV1: new SDKModule(test)
        },
        registry: msgTypes
    };
};
export default Module;
//# sourceMappingURL=module.js.map