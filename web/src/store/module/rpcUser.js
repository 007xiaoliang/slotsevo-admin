export const rpcUser = {
    namespaced: true,
    state: {
        rpcUserID: "",
        rpcBaseInfo: "",
        rpcThemeInfo: "",
        rpcActivityInfo: ""
    },
    mutations: {
        setRpcUserID(state, rpcUserID) {
            state.rpcUserID = rpcUserID
        },
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
        rpcUserID(state) {
            return state.rpcUserID
        },
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
