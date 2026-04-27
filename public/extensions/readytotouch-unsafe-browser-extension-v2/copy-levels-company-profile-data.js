"use strict";

console.log("[RTT] Levels.fyi company profile copy loaded");

// Ctrl+Shift+Y / Ctrl+Shift+Н
document.body.addEventListener("keydown", (event) => {
    if (event.ctrlKey && event.shiftKey && (event.key === "Y" || event.key === "Н")) {
        const alias     = RTT.parseVanityFromPath(
            window.location.href, "/companies/",
            "Expected URL like https://www.levels.fyi/companies/company-name/"
        );
        const employees = getEmployees();

        const out = `\t\t\t\tAlias:     "${alias}",
\t\t\t\tEmployees: "${employees}",
\t\t\t\tDate:      mustDate("${RTT.today()}"),`;

        RTT.copyToClipboard(out);
    }
});

function getEmployees() {
    let employees = "";

    const $elements = document.querySelectorAll("h6");
    for (const $element of $elements) {

        if ($element.nextElementSibling !== null && $element.nextElementSibling.textContent.trim() === "# of Employees") {
            employees = $element.textContent.trim();
        }
    }

    return employees;
}
