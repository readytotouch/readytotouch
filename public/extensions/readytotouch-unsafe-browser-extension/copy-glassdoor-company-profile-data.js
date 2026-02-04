console.log("Glassdoor company profile data copy extension loaded");

// Ctrl + Shift + Y
document.body.addEventListener("keydown", (event) => {
    // Y is for English, Н is for Ukrainian
    if (event.ctrlKey && event.shiftKey && (event.key === "Y" || event.key === "Н")) {
        const goGlassdoorProfileColumns = `				OverviewURL: "${document.querySelector('a[href^="/Overview/"]').href}",
				ReviewsURL:  "${document.querySelector('a[href^="/Reviews/"]').href}",
				JobsURL:     "${document.querySelector('a[href^="/Jobs/"]').href}",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "${getEmployerRating()}",
				Verified:    ${isEngagedEmployer() ? "true" : "false"},
				Date:        mustDate("${date()}"),`

        navigator.clipboard.writeText(goGlassdoorProfileColumns)
            .then(() => console.log("Page info copied to clipboard:", goGlassdoorProfileColumns))
            .catch((err) => console.error("Failed to copy page info:", err));
    }
});

function trim($element) {
    const value = $element.textContent.trim();
    if (value === "--") {
        return "";
    }

    return value;
}

function isEngagedEmployer() {
    const elements = document.querySelectorAll(`p[class^="employer-engagement-status_engagement"]`);

    return Array.from(elements).some($element => $element.textContent.trim() === "Engaged Employer");
}

function getEmployerRating() {
    const $rating = document.querySelector(`span[class^="employer-overview_employerOverviewRating"]`);
    if ($rating) {
        return $rating.childNodes[0].textContent.trim();
    }

    return "";
}

function date() {
    return new Intl.DateTimeFormat("en-CA").format(new Date());
}
