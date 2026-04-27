"use strict";

console.log("[RTT] Indeed vacancy copy loaded");

// Ctrl+Shift+Y / Ctrl+Shift+Н
document.body.addEventListener("keydown", (event) => {
    if (!event.ctrlKey || !event.shiftKey) return;
    if (event.key !== "Y" && event.key !== "Н") return;

    const titleRaw = document.querySelector("h2")?.innerText.trim() ?? "";
    const title = RTT.normalizeTitle(titleRaw);

    const out = `{
\t\t\t\t\t\t    Title:                "${title}",
\t\t\t\t\t\t    ShortDescription:     "",
\t\t\t\t\t\t    SwitchingOpportunity: "",
\t\t\t\t\t\t    Location:             "",
\t\t\t\t\t\t    URL:                  "${getURL()}",
\t\t\t\t\t\t    Date:                 mustDate("${RTT.today()}"),
\t\t\t\t\t\t    WithSalary:           ${hasSalary() ? "true" : "false"},
\t\t\t\t\t\t    Remote:               ${isRemote() ? "true" : "false"},
\t\t\t\t\t\t},`;

    RTT.copyToClipboard(out);
});

function getURL() {
    const params = new URLSearchParams(window.location.search);
    const jk = params.get("jk");
    if (jk) {
        return `${window.location.origin + window.location.pathname}?jk=${jk}`;
    }

    return "";
}

function remote() {
    const $elements = document.querySelectorAll('[class^="js-match-insights-provider-"]');

    return Array
        .from($elements)
        .some($element => $element.textContent.trim().toLowerCase() === "remote");
}

function hasSalary() {
    const $elements = document.querySelectorAll('[class^="js-match-insights-provider-"]');

    return Array
        .from($elements)
        .some($element => $element.textContent.trim().toLowerCase().includes("$"));
}
