export interface RawResponse {
    code: number;
    message: string;
    data: RawItem[];
}
export interface RawNameResponse {
    code: number;
    message: string;
    data: string;
}
export interface RawItem {
    full_path: string;
    name: string;
    path: string;
}
