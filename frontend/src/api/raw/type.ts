export interface RawResponse {
    code: number;
    message: string;
    data: RawItem[];
}

export interface RawItem {
    full_path: string;
    name: string;
    path: string;
}
