export interface rule {
    id: number;
    vid:string;
    provider:string;
    file_extract_reg:string
    season:number
    language:string
    episode_extract_reg:string
    episode_position:number
    episode_offset:number
    type:string,
    name:string
}

export interface ruleRequestData {
    vid:string;
    provider:string;
    file_extract_reg:string
    season?:number
    language?:string
    episode_extract_reg:string
    episode_position?:number
    episode_offset?:number
    type:string,
    name?:string
}

export interface RuleResponse {
    code: number;
    message: string;
    data: rule[];
}