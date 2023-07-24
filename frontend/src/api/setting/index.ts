import request from "../../utils/request";
import { Setting, SettingResponse } from "./type";


enum API {
    API_RAW= '/config',
}

export const reqGetSetting = () => request.get<SettingResponse>(API.API_RAW)
export const  reqPostSetting = (data: Setting) => request.post<SettingResponse>(API.API_RAW, data)