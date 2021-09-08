import service from '@/utils/request'

export const getActInfo = (data) => {
    return service({
        url: "/rpc/actInfo",
        method: 'post',
        data: data
    })
}
export const getActWeight = (data) => {
    return service({
        url: "/rpc/actWeight",
        method: 'post',
        data: data
    })
}

