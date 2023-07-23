import axios from 'axios'

let request  = axios.create({
    baseURL: '/v1',
    timeout: 10000,

})

export default request