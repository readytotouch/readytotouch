export function organizersWelcome() {
    const pathname = window.location.pathname;

    const index = pathname.indexOf("/", 2);

    return `${pathname.substring(0, index)}/welcome?redirect=${encodeURIComponent(pathname)}`;
}

export function welcome(companyUrl: string = '') {
    let redirect = window.location.pathname;

    if (companyUrl) {
        redirect = window.location.pathname + "?company-url=" + encodeURIComponent(companyUrl)
    }

    return `/golang/welcome?redirect=${encodeURIComponent(redirect)}`;
}
