console.log("LinkedIn vacancy data copy extension loaded");

// Ctrl + Shift + Y
document.body.addEventListener("keydown", (event) => {
    // Y is for English, Н is for Ukrainian
    if (event.ctrlKey && event.shiftKey && (event.key === "Y" || event.key === "Н")) {

        const title = document.querySelector("h1").innerText.trim()
            .replace(" - ", " – ") // Replace hyphen with dash
            .replace("GoLang", "Golang") // Replace GoLang with Golang
            .replace("Goland", "Golang") // Replace Goland with Golang
        ;
        const goLinkedInVacancyColumns = `{
						    Title:                "${title}",
						    ShortDescription:     "",
						    SwitchingOpportunity: "",
						    URL:                  "${window.location.origin + window.location.pathname}",
						    Date:                 mustDate("${date()}"),
						    WithSalary:           false,
						    Remote:               ${remote() ? "true" : "false"},
						},`

        navigator.clipboard.writeText(goLinkedInVacancyColumns)
            .then(() => console.log("Page info copied to clipboard:", goLinkedInVacancyColumns))
            .catch((err) => console.error("Failed to copy page info:", err));
    }
});

function remote() {
    const $elements = document.querySelectorAll(".job-details-preferences-and-skills span");
    for (const $element of $elements) {
        if ($element.textContent.trim().toLowerCase() === "remote") {
            return true;
        }
    }

    return false;
}

function date() {
    const $elements = document.querySelectorAll('.job-details-jobs-unified-top-card__primary-description-container span');

    for (const $element of $elements) {
        const text = $element.textContent.trim().toLowerCase();

        if (text.includes("hour ago") || text.includes("hours ago")) {
            return new Intl.DateTimeFormat("en-CA").format(new Date());
        }
    }

    return new Intl.DateTimeFormat("en-CA").format(new Date());
}
