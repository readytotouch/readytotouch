import {parseCurrentProgrammingLanguage} from "./pl";

export function organizersWelcome() {
    const pathname = window.location.pathname;

    return `/${parseCurrentProgrammingLanguage(pathname)}/welcome?redirect=${encodeURIComponent(pathname)}`;
}

export function welcome(companyUrl: string = "") {
    let redirect = window.location.pathname;

    if (companyUrl) {
        redirect = window.location.pathname + "?company-url=" + encodeURIComponent(companyUrl)
    }

    return `/golang/welcome?redirect=${encodeURIComponent(redirect)}`;
}
