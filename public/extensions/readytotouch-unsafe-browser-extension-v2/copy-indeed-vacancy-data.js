"use strict";

console.log("[RTT] Indeed vacancy copy loaded");

// Ctrl+Shift+Y / Ctrl+Shift+Н
document.body.addEventListener("keydown", (event) => {
    if (!event.ctrlKey || !event.shiftKey) return;
    if (event.key !== "Y" && event.key !== "Н") return;

    const titleRaw = document.querySelector("h2")?.innerText.trim() ?? "";
    const title    = RTT.normalizeTitle(titleRaw);

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
    const jk = new URLSearchParams(window.location.search).get("jk");
    return jk ? `${window.location.origin + window.location.pathname}?jk=${jk}` : "";
}

function isRemote() {
    return Array.from(document.querySelectorAll('[class^="js-match-insights-provider-"]'))
        .some(el => el.textContent.trim().toLowerCase() === "remote");
}

function hasSalary() {
    return Array.from(document.querySelectorAll('[class^="js-match-insights-provider-"]'))
        .some(el => el.textContent.trim().toLowerCase().includes("$"));
}
