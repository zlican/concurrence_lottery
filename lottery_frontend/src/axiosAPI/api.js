import request from "./config";

export const api = {
    getHome(){
        return request({
            url: '/home',
            method: 'GET'
        })
	},
    getAllPrize(){
        return request({
            url: '/prize/all',
            method: 'GET'
        })
    },
    lottery(){
        return request({
            url: '/prize/lottery',
            method: 'GET'
        })
    }
}