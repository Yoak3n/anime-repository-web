import request from "../../utils/request";
import type {RawResponse} from "./type";


enum API {
    API_RAW= '/raw',
}
export const reqGetRaw = () => request.get<RawResponse>(API.API_RAW)