import service from '@/utils/request'

export const getThemeInfo = (data) => {
    return service({
        url: "/rpc/themeInfo",
        method: 'post',
        data: data
    })
}

export const getThemeWeight = (data) => {
    return service({
        url: "/rpc/themeWeight",
        method: 'post',
        data: data
    })
}
