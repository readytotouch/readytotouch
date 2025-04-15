console.log("GitHub company profile data copy extension loaded");

// Ctrl + Shift + Y
document.body.addEventListener("keydown", (event) => {
    // Y is for English, Н is for Ukrainian
    if (event.ctrlKey && event.shiftKey && (event.key === "Y" || event.key === "Н")) {
        const goLinkedInProfileColumns = `				Login:     "${parseVanityName(window.location.href)}",
				Followers: "${followers()}",
				Verified:  ${document.querySelectorAll('summary[title="Label: Verified"]').length > 0 ? "true" : "false"},`

        navigator.clipboard.writeText(goLinkedInProfileColumns)
            .then(() => console.log("Page info copied to clipboard:", goLinkedInProfileColumns))
            .catch((err) => console.error("Failed to copy page info:", err));
    }
});

function parseVanityName(url) {
    // Your errors = your pain
    const error = "Expected URL like https://github.com/company-name/";

    let parsedUrl = null;

    try {
        parsedUrl = new URL(url);
    } catch (e) {
        alert(error);

        return "";
    }

    const prefix = "/";

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

function followers() {
    const $elements = document.querySelectorAll("a.Link--primary");

    for (const $element of $elements) {
        if ($element.href.endsWith("/followers")) {
            return $element.querySelector("span").textContent.trim();
        }
    }

    return "";
}