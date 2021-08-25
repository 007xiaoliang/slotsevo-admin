export const rpcUser = {
    namespaced: true,
    state: {
        rpcBaseInfo: "",
        rpcThemeInfo: "",
        rpcActivityInfo: ""
    },
    mutations: {
        setRpcBaseInfo(state, rpcBaseInfo) {
            state.rpcBaseInfo = rpcBaseInfo
        },
        setRpcThemeInfo(state, rpcThemeInfo) {
            state.rpcThemeInfo = rpcThemeInfo
        },
        setRpcActivityInfo(state, rpcActivityInfo) {
            state.rpcActivityInfo = rpcActivityInfo
        }
    },
    getters: {
        rpcBaseInfo(state) {
            return state.rpcBaseInfo
        },
        rpcThemeInfo(state) {
            return state.rpcThemeInfo
        },
        rpcActivityInfo(state) {
            return state.rpcActivityInfo
        },

    }
}
