import request from "../../utils/request";
import type {RuleResponse,ruleRequestData} from "./type";


enum API {
   API_RULE = "/rule"
}


export const reqGetRule = () => request.get<RuleResponse>(API.API_RULE)
export const reqAddRule = (data:ruleRequestData) =>request.post<RuleResponse>(API.API_RULE,data)
export const reqDeleteRule = (id:number)=>request.delete<RuleResponse>(API.API_RULE+`/${id}`)
export const reqModifyRule = (id:number,data:ruleRequestData)=>request.patch<RuleResponse>(API.API_RULE+`/${id}`,data)