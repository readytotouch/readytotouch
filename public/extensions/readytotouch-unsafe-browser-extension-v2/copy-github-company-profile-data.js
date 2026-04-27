"use strict";

console.log("[RTT] GitHub company profile copy loaded");

// Ctrl+Shift+Y / Ctrl+Shift+Н
document.body.addEventListener("keydown", (event) => {
    if (!event.ctrlKey || !event.shiftKey) return;
    if (event.key !== "Y" && event.key !== "Н") return;

    const login = RTT.parseVanityFromPath(
        window.location.href, "/",
        "Expected URL like https://github.com/company-name/"
    );

    const out = `\t\t\t\tLogin:     "${login}",
\t\t\t\tFollowers: "${getFollowers()}",
\t\t\t\tVerified:  ${document.querySelectorAll('summary[title="Label: Verified"]').length > 0 ? "true" : "false"},
\t\t\t\tDate:      mustDate("${RTT.today()}"),`;

    RTT.copyToClipboard(out);
});

function getFollowers() {
    for (const el of document.querySelectorAll("a.Link--primary")) {
        if (el.href.endsWith("/followers")) {
            return el.querySelector("span")?.textContent.trim() ?? "";
        }
    }
    return "";
}
