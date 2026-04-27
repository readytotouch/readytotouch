"use strict";

console.log("[RTT] Glassdoor company profile copy loaded");

// Ctrl+Shift+Y / Ctrl+Shift+Н
document.body.addEventListener("keydown", (event) => {
    if (!event.ctrlKey || !event.shiftKey) return;
    if (event.key !== "Y" && event.key !== "Н") return;

    const out = `\t\t\t\tOverviewURL: "${qs('div[class^="EmployerInfo_employerInfoWrapper"] a[href^="/Overview/"]')?.href ?? ""}",
\t\t\t\tReviewsURL:  "${qs('div[class^="EmployerInfo_employerInfoWrapper"] a[href^="/Reviews/"]')?.href ?? ""}",
\t\t\t\tJobsURL:     "${qs('div[class^="EmployerInfo_employerInfoWrapper"] a[href^="/Jobs/"]')?.href ?? ""}",
\t\t\t\tJobs:        "", // Becomes empty in the new design version
\t\t\t\tReviews:     "", // Becomes empty in the new design version
\t\t\t\tSalaries:    "", // Becomes empty in the new design version
\t\t\t\tReviewsRate: "${getEmployerRating()}",
\t\t\t\tVerified:    ${isEngagedEmployer() ? "true" : "false"},
\t\t\t\tDate:        mustDate("${RTT.today()}"),`;

    RTT.copyToClipboard(out);
});

function qs(selector) {
    return document.querySelector(selector);
}

function isEngagedEmployer() {
    return Array.from(document.querySelectorAll('div[class^="EmployerInfo_employerInfoEngagementStatus"]'))
        .some(el => el.textContent.trim() === "Engaged Employer");
}

function getEmployerRating() {
    const el = document.querySelector('span[class^="employer-overview_employerOverviewRating"]');
    return el ? el.childNodes[0].textContent.trim() : "";
}
