"use strict";

console.log("[RTT] Blind company profile copy loaded");

// Ctrl+Shift+Y / Ctrl+Shift+Н
document.body.addEventListener("keydown", (event) => {
    if (!event.ctrlKey || !event.shiftKey) return;
    if (event.key !== "Y" && event.key !== "Н") return;

    const alias               = RTT.parseVanityFromPath(
        window.location.href, "/company/",
        "Expected URL like https://www.teamblind.com/company/company-name/"
    );
    const { ratingValue, ratingCount } = getEmployerRating();
    const { employees, salary }        = getCompanyInfo();

    const out = `\t\t\t\tAlias:       "${alias}",
\t\t\t\tEmployees:   "${employees}",
\t\t\t\tSalary:      "${salary}",
\t\t\t\tReviews:     "${ratingCount}",
\t\t\t\tReviewsRate: "${ratingValue}",`;

    RTT.copyToClipboard(out);
});

function getEmployerRating() {
    for (const el of document.querySelectorAll('script[type="application/ld+json"]')) {
        try {
            const json = JSON.parse(el.textContent.trim());
            if (json["@type"] === "EmployerAggregateRating") {
                return {
                    ratingValue: json.ratingValue,
                    ratingCount: json.ratingCount.toLocaleString("en-US"),
                };
            }
        } catch (_) {}
    }
    return { ratingValue: "", ratingCount: "" };
}

function getCompanyInfo() {
    const h1 = document.querySelector("h1");
    if (!h1) return { employees: "", salary: "" };

    let employees = "", salary = "";

    for (const el of h1.nextElementSibling?.querySelectorAll("div.text-sm") ?? []) {
        if (el.textContent.trim() === "Size") {
            employees = el.nextElementSibling?.textContent.trim().replace("employees", "").trim() ?? "";
        }
        if (el.textContent.trim() === "Salary") {
            salary = el.nextElementSibling?.textContent.trim() ?? "";
            if (salary === "-") salary = "";
        }
    }

    return { employees, salary };
}
