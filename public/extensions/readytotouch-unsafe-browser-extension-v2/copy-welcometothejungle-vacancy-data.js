"use strict";

console.log("[RTT] Welcome to the Jungle vacancy copy loaded");

// Ctrl+Shift+Y / Ctrl+Shift+Н
document.body.addEventListener("keydown", (event) => {
    if (!event.ctrlKey || !event.shiftKey) return;
    if (event.key !== "Y" && event.key !== "Н") return;

    const titleRaw = document.querySelector("h1")?.innerText.trim() ?? "";
    const title    = RTT.normalizeTitle(titleRaw);

    const out = `{
\t\t\t\t\t\t    Title:                "${title}",
\t\t\t\t\t\t    ShortDescription:     "",
\t\t\t\t\t\t    SwitchingOpportunity: "",
\t\t\t\t\t\t    URL:                  "${window.location.origin + window.location.pathname}",
\t\t\t\t\t\t    Date:                 mustDate("${RTT.today()}"),
\t\t\t\t\t\t    WithSalary:           false,
\t\t\t\t\t\t    Remote:               false,
\t\t\t\t\t\t},`;

    RTT.copyToClipboard(out);
});
