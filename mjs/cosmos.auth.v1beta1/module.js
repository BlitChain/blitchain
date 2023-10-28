// Generated by Ignite ignite.com/cli
import { Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { Api } from "./rest";
import { BaseAccount as typeBaseAccount } from "./types";
import { ModuleAccount as typeModuleAccount } from "./types";
import { ModuleCredential as typeModuleCredential } from "./types";
import { Params as typeParams } from "./types";
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
            BaseAccount: getStructure(typeBaseAccount.fromPartial({})),
            ModuleAccount: getStructure(typeModuleAccount.fromPartial({})),
            ModuleCredential: getStructure(typeModuleCredential.fromPartial({})),
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
            CosmosAuthV1Beta1: new SDKModule(test)
        },
        registry: msgTypes
    };
};
export default Module;
//# sourceMappingURL=module.js.map