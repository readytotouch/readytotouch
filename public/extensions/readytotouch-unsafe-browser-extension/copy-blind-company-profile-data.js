"use strict";

console.log("[RTT] Blind company profile copy loaded");

// Ctrl+Shift+Y / Ctrl+Shift+Н
document.body.addEventListener("keydown", (event) => {
    if (event.ctrlKey && event.shiftKey && (event.key === "Y" || event.key === "Н")) {

        const alias = RTT.parseVanityFromPath(
            window.location.href, "/company/",
            "Expected URL like https://www.teamblind.com/company/company-name/"
        );
        const {ratingValue, ratingCount} = getEmployerRating();
        const {employees, salary} = getCompanyInfo();

        const out = `\t\t\t\tAlias:       "${alias}",
\t\t\t\tEmployees:   "${employees}",
\t\t\t\tSalary:      "${salary}",
\t\t\t\tReviews:     "${ratingCount}",
\t\t\t\tReviewsRate: "${ratingValue}",
\t\t\t\tDate:        mustDate("${RTT.today()}"),`;

        RTT.copyToClipboard(out);
    }
});

function getEmployerRating() {
    const $scripts = document.querySelectorAll('script[type="application/ld+json"]');

    for (const $script of $scripts) {
        try {
            const json = JSON.parse($script.textContent.trim());
            if (json["@type"] === "EmployerAggregateRating") {
                return {
                    ratingValue: json.ratingValue,
                    ratingCount: json.ratingCount.toLocaleString("en-US"),
                };
            }
        } catch (error) {
            console.error("Error parsing JSON-LD:", error);
        }
    }

    return {
        ratingValue: "",
        ratingCount: ""
    };
}

function getCompanyInfo() {
    let employees = "",
        salary = "";

    const $rows = document.querySelectorAll("div.text-sm");
    for (const $row of $rows) {
        const $first = $row.firstChild;
        if ($first.textContent.trim() === "Size") {
            employees = $first.nextElementSibling.textContent.trim()
                .replace("employees", "")
                .trim();
        } else if ($first.textContent.trim() === "Salary") {
            salary = $first.nextElementSibling.textContent.trim();

            if (salary === "-") {
                salary = "";
            }
        }
    }

    return {employees, salary};
}
