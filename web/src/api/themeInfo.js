import service from '@/utils/request'

export const getThemeInfo = (data) => {
    return service({
        url: "/rpc/themeInfo",
        method: 'post',
        data: data
    })
}

