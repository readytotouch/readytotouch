export function toEnter(handler: () => void) {
    return function (event) {
        if (event.code === "Enter") {
            handler();
        }
    }
}
