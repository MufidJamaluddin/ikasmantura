import {match} from "react-router";

export function getSubroute(basepath: match<any>, subpath:string): string
{
    let path:string

    if(subpath === null || subpath === '') path = basepath.url
    else if (basepath.url === '/') path = `/${subpath}`
    else path = `${basepath.url}/${subpath}`

    return path
}