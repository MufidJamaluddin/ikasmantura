import moment from "moment";

export function dateParser(date){
    return moment(date).format("YYYY-MM-DDTHH:mm:ssZ")
}