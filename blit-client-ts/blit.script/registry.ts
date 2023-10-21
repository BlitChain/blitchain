import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateScript } from "./types/blit/script/tx";
import { MsgRun } from "./types/blit/script/msgrun";
import { MsgCreateScript } from "./types/blit/script/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/blit.script.MsgUpdateScript", MsgUpdateScript],
    ["/blit.script.MsgRun", MsgRun],
    ["/blit.script.MsgCreateScript", MsgCreateScript],
    
];

export { msgTypes }