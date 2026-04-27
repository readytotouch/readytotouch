"use strict";

console.log("[RTT] Indeed company profile copy loaded");

// Ctrl+Shift+Y / Ctrl+Shift+Н
document.body.addEventListener("keydown", (event) => {
    if (!event.ctrlKey || !event.shiftKey) return;
    if (event.key !== "Y" && event.key !== "Н") return;

    const alias = RTT.parseVanityFromPath(
        window.location.href, "/cmp/",
        "Expected URL like https://www.indeed.com/cmp/Google"
    );

    RTT.copyToClipboard(`\t\t\t\tAlias: "${alias}",`);
});
