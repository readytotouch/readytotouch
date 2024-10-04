export function organizersWelcome() {
    const pathname = window.location.pathname;

    const index = pathname.lastIndexOf('/');

    return `${pathname.substring(0, index)}/welcome?redirect=${encodeURIComponent(pathname)}`;
}
