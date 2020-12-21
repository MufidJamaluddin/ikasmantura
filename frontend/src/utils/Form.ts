export function ToFormData(data: Object = {}): FormData {
    let formData = new FormData()
    Object.keys(data).forEach(key => {
        formData.append(key, data[key])
    })
    return formData
}
