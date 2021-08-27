import service from '@/utils/request'

export const getUserInfo = (data) => {
    return service({
        url: "/rpc/userInfo",
        method: 'post',
        data: data
    })
}

