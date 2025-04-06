console.log("Indeed vacancy data copy extension loaded");

// Ctrl + Shift + Y
document.body.addEventListener("keydown", (event) => {
    // Y is for English, Н is for Ukrainian
    if (event.ctrlKey && event.shiftKey && (event.key === "Y" || event.key === "Н")) {

        const title = document.querySelector("h2").innerText.trim()
            .replaceAll(" - ", " – ") // Replace hyphen with dash
            .replace("(m/f/x)", " ").trim()
            .replace("(m/f/d)", " ").trim()
            .replace("(d/f/m)", " ").trim()
            .replace("Sr.", "Senior")
            .replace("GoLang", "Golang")
            .replace("Goland", "Golang")
            .replace("Back End", "Back-End")
        ;

        const goLinkedInVacancyColumns = `{
						    Title:                "${title}",
						    ShortDescription:     "",
						    SwitchingOpportunity: "",
						    URL:                  "${getURL()}",
						    Date:                 mustDate("${date()}"),
						    WithSalary:           ${salary() ? "true" : "false"},
						    Remote:               ${remote() ? "true" : "false"},
						},`

        navigator.clipboard.writeText(goLinkedInVacancyColumns)
            .then(() => console.log("Page info copied to clipboard:", goLinkedInVacancyColumns))
            .catch((err) => console.error("Failed to copy page info:", err));
    }
});

function getURL() {
    const params = new URLSearchParams(window.location.search);
    const jk = params.get("jk");
    if (jk) {
        return `${window.location.origin + window.location.pathname}?jk=${jk}`;
    }

    return "";
}

function remote() {
    const $elements = document.querySelectorAll('[class^="js-match-insights-provider-"]');

    return Array.from($elements).some($element => $element.textContent.trim().toLowerCase() === "remote");
}

function salary() {
    const $elements = document.querySelectorAll('[class^="js-match-insights-provider-"]');

    return Array.from($elements).some($element => $element.textContent.trim().toLowerCase().includes("$"));
}

function date() {
    return new Intl.DateTimeFormat("en-CA").format(new Date());
}
