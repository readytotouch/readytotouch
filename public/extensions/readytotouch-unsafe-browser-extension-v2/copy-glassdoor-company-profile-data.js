"use strict";

console.log("[RTT] Glassdoor company profile copy loaded");

// Ctrl+Shift+Y / Ctrl+Shift+Н
document.body.addEventListener("keydown", (event) => {
    if (event.ctrlKey && event.shiftKey && (event.key === "Y" || event.key === "Н")) {

        const out = `\t\t\t\tOverviewURL: "${document.querySelector('div[class^="EmployerInfo_employerInfoWrapper"] a[href^="/Overview/"]')?.href ?? ""}",
\t\t\t\tReviewsURL:  "${document.querySelector('div[class^="EmployerInfo_employerInfoWrapper"] a[href^="/Reviews/"]')?.href ?? ""}",
\t\t\t\tJobsURL:     "${document.querySelector('div[class^="EmployerInfo_employerInfoWrapper"] a[href^="/Jobs/"]')?.href ?? ""}",
\t\t\t\tJobs:        "", // Becomes empty in the new design version
\t\t\t\tReviews:     "", // Becomes empty in the new design version
\t\t\t\tSalaries:    "", // Becomes empty in the new design version
\t\t\t\tReviewsRate: "${getEmployerRating()}",
\t\t\t\tVerified:    ${isEngagedEmployer() ? "true" : "false"},
\t\t\t\tDate:        mustDate("${RTT.today()}"),`;

        RTT.copyToClipboard(out);
    }
});

function isEngagedEmployer() {
    const $elements = document.querySelectorAll('div[class^="EmployerInfo_employerInfoEngagementStatus"]');

    return Array
        .from($elements)
        .some($element => $element.textContent.trim() === "Engaged Employer");
}

function getEmployerRating() {
    const $rating = document.querySelector(`div[class^="CompanyOverview_overallRating"]`);
    if ($rating) {
        return $rating.childNodes[0].textContent.trim();
    }

    return "";
}
