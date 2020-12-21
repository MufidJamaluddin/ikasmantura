export function ToFormData(data: Object = {}): FormData {
    let formData = new FormData()
    Object.keys(data).map(key => {
        formData.append(key, data[key])
    })
    return formData
}
