import request from "../../utils/request";
import type {RawResponse,RawNameResponse} from "./type";


enum API {
    API_RAW= '/raw',
    API_HOST = '/host'
}
export const reqGetTVName = (id:string)=> request.get<RawNameResponse>(API.API_RAW+"/tv/"+id)
export const reqGetMovieName = (id:string)=> request.get<RawNameResponse>(API.API_RAW+"/movie/"+id)

export const reqGetRaw = () => request.get<RawResponse>(API.API_RAW)
export const reqGetHost = () => request.get<RawNameResponse>(API.API_HOST)