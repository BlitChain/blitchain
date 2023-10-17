import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRun } from "./types/blit/script/msgrun";
import { MsgUpdateScript } from "./types/blit/script/tx";
import { MsgCreateScript } from "./types/blit/script/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/blit.script.MsgRun", MsgRun],
    ["/blit.script.MsgUpdateScript", MsgUpdateScript],
    ["/blit.script.MsgCreateScript", MsgCreateScript],
    
];

export { msgTypes }