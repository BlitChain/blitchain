import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateStorage } from "./types/blit/storage/tx";
import { MsgCreateStorage } from "./types/blit/storage/tx";
import { MsgDeleteStorage } from "./types/blit/storage/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/blit.storage.MsgUpdateStorage", MsgUpdateStorage],
    ["/blit.storage.MsgCreateStorage", MsgCreateStorage],
    ["/blit.storage.MsgDeleteStorage", MsgDeleteStorage],
    
];

export { msgTypes }