/**
 * stripTags
 *
 * @parm mixed output
 * @param input
 */

export function strip_tags(input) {
    if (input) {
        let tags = /(<([^>]+)>)/ig;
        if (!Array.isArray(input)) {
            input = input.replace(tags,'')
        }
        else {
            let i = input.length;
            while(i--) {
                input[i] = input[i].replace(tags,'')
            }
        }
        return input;
    }
    return false;
}
