type action = 'GET' | 'POST' | 'PUT' | 'DELETE'

interface DownloadAction {
    formId?: string
    formName?: string
    action: action
    path: string
}

export function download(downloadAction: DownloadAction) {
    const my_form: any = document.createElement('FORM')
    my_form.name = downloadAction.formName ?? 'myDownloadForm'
    my_form.method = downloadAction.action
    my_form.action = process.env.PUBLIC_URL + downloadAction.path

    if(downloadAction.formId) {
        let nodes = document.getElementById(downloadAction.formId).children

        for (let i = 0; i < nodes.length; i++) {
            my_form.appendChild(nodes.item(i).cloneNode(true));
        }
    }

    document.body.appendChild(my_form);
    my_form.submit();
}