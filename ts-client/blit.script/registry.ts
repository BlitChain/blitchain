import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRun } from "./types/blit/script/msgrun";
import { MsgCreateScript } from "./types/blit/script/tx";
import { MsgUpdateScript } from "./types/blit/script/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/blit.script.MsgRun", MsgRun],
    ["/blit.script.MsgCreateScript", MsgCreateScript],
    ["/blit.script.MsgUpdateScript", MsgUpdateScript],
    
];

export { msgTypes }