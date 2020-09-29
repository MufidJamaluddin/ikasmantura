export function toSentenceCase(str?: string) {
    if(str === null || str === undefined) return ''

    return str.replace(/(?:^\w|[A-Z]|\b\w)/g, function(word, index) {
        return word.toUpperCase();
    }).replace(/\s+/g, '');
}