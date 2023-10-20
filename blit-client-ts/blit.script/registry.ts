import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateScript } from "./types/blit/script/tx";
import { MsgRun } from "./types/blit/script/msgrun";
import { MsgUpdateScript } from "./types/blit/script/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/blit.script.MsgCreateScript", MsgCreateScript],
    ["/blit.script.MsgRun", MsgRun],
    ["/blit.script.MsgUpdateScript", MsgUpdateScript],
    
];

export { msgTypes }