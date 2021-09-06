import service from '@/utils/request'

export const getActInfo = (data) => {
    return service({
        url: "/rpc/actInfo",
        method: 'post',
        data: data
    })
}

