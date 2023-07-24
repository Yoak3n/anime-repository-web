export interface SettingResponse {
    code: number;
    message: string;
    data: Setting;
}

export interface Setting {
    key:string
    use_proxy:boolean
    proxy_url:string
    raw_path:string
    tv_path:string
    movie_path:string
    delay:number
}