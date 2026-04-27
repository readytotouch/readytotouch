"use strict";

console.log("[RTT] Levels.fyi company profile copy loaded");

// Ctrl+Shift+Y / Ctrl+Shift+Н
document.body.addEventListener("keydown", (event) => {
    if (!event.ctrlKey || !event.shiftKey) return;
    if (event.key !== "Y" && event.key !== "Н") return;

    const alias     = RTT.parseVanityFromPath(
        window.location.href, "/companies/",
        "Expected URL like https://www.levels.fyi/companies/company-name/"
    );
    const employees = getEmployees();

    const out = `\t\t\t\tAlias:     "${alias}",
\t\t\t\tEmployees: "${employees}",`;

    RTT.copyToClipboard(out);
});

function getEmployees() {
    for (const el of document.querySelectorAll("h6")) {
        if (el.nextElementSibling?.textContent.trim() === "# of Employees") {
            return el.textContent.trim();
        }
    }
    return "";
}
