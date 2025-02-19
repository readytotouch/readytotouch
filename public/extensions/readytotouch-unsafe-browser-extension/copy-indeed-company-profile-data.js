console.log("Blind company profile data copy extension loaded");

// Ctrl + Shift + Y
document.body.addEventListener("keydown", (event) => {
    // Y is for English, Н is for Ukrainian
    if (event.ctrlKey && event.shiftKey && (event.key === "Y" || event.key === "Н")) {
        const goIndeedProfileColumns = `				Alias: "${parseVanityName(window.location.href)}",`

        navigator.clipboard.writeText(goIndeedProfileColumns)
            .then(() => console.log("Page info copied to clipboard:", goIndeedProfileColumns))
            .catch((err) => console.error("Failed to copy page info:", err));
    }
});

function parseVanityName(url) {
    // Your errors = your pain
    const error = "Expected URL like https://www.indeed.com/cmp/Google";

    let parsedUrl = null;

    try {
        parsedUrl = new URL(url);
    } catch (e) {
        alert(error);

        return "";
    }

    const prefix = "/cmp/";

    if (parsedUrl.pathname.indexOf(prefix) === -1) {
        alert(error);

        return "";
    }

    const end = parsedUrl.pathname.indexOf("/", prefix.length);
    if (end === -1) {
        return parsedUrl.pathname.substring(prefix.length);
    }

    return parsedUrl.pathname.substring(prefix.length, end);
}
