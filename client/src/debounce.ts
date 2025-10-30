export function debounce(handler: () => void, wait: number) {
    let timeout

    const later = function () {
        timeout = null;

        handler();
    };

    return function () {
        clearTimeout(timeout)

        timeout = setTimeout(later, wait)
    }
}
