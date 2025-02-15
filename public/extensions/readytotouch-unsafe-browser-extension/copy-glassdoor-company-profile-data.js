console.log("Glassdoor company profile data copy extension loaded");

// Ctrl + Shift + Y
document.body.addEventListener("keydown", (event) => {
    // Y is for English, Н is for Ukrainian
    if (event.ctrlKey && event.shiftKey && (event.key === "Y" || event.key === "Н")) {
        const goGlassdoorProfileColumns = `				OverviewURL: "${document.querySelector("#overview > a").href}",
				ReviewsURL:  "${document.querySelector("#reviews > a").href}",
				JobsURL:     "${document.querySelector("#jobs > a").href}",
				Jobs:        "${document.querySelector("#jobs div:first-of-type span:first-of-type").textContent.trim()}",
				Reviews:     "${document.querySelector("#reviews div:first-of-type span:first-of-type").textContent.trim()}",
				Salaries:    "${document.querySelector("#salaries div:first-of-type span:first-of-type").textContent.trim()}",
				ReviewsRate: "${getEmployerRating()}",
				Verified:    ${isEngagedEmployer() ? "true" : "false"},`

        navigator.clipboard.writeText(goGlassdoorProfileColumns)
            .then(() => console.log("Page info copied to clipboard:", goGlassdoorProfileColumns))
            .catch((err) => console.error("Failed to copy page info:", err));
    }
});


function isEngagedEmployer() {
    const elements = document.querySelectorAll(`p[class^="employer-engagement-status_engementTrigger"]`);

    return Array.from(elements).some($element => $element.textContent.trim() === "Engaged Employer");
}

function getEmployerRating() {
    const $rating = document.querySelector(`span[class^="employer-overview_employerOverviewRating"]`);
    if ($rating) {
        return $rating.childNodes[0].textContent.trim();
    }

    return "";
}
