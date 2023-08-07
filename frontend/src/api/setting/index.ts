import request from "../../utils/request";
import { Setting, SettingResponse } from "./type";


enum API {
    API_CONFIG= '/config',
}

export const reqGetSetting = () => request.get<SettingResponse>(API.API_CONFIG)
export const  reqPostSetting = (data: Setting) => request.post<SettingResponse>(API.API_CONFIG, data)