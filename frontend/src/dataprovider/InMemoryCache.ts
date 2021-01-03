const InMemoryCache = () => {

    let __cache = {}

    const getCache = (key) => {
        return __cache[key] ?? null
    }

    const setCache = (key, value) => {
        __cache[key] = value
    }

    const resetCache = () => {
        __cache = {}
    }

    return {
        getCache, setCache, resetCache
    }
}

export default InMemoryCache()
