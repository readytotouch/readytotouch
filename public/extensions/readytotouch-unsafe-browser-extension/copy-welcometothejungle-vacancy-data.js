console.log("Welcome to the Jungle vacancy data copy extension loaded");

// Ctrl + Shift + Y
document.body.addEventListener("keydown", (event) => {
    // Y is for English, Н is for Ukrainian
    if (event.ctrlKey && event.shiftKey && (event.key === "Y" || event.key === "Н")) {

        const title = document.querySelector("h1").innerText.trim()
            .replaceAll(" - ", " – ") // Replace hyphen with dash
            .replace("(m/f/x)", " ").trim()
            .replace("(m/f/d)", " ").trim()
            .replace("Sr.", "Senior")
            .replace("GoLang", "Golang")
            .replace("Goland", "Golang")
            .replace("Back End", "Back-End")
        ;

        const goLinkedInVacancyColumns = `{
						    Title:                "${title}",
						    ShortDescription:     "${document.querySelector("h2").innerText.trim()}",
						    SwitchingOpportunity: "",
						    URL:                  "${window.location.origin + window.location.pathname}",
						    Date:                 mustDate("${date()}"),
						    WithSalary:           ${salary() ? "true" : "false"},
						    Remote:               ${remote() ? "true" : "false"},
						},`

        navigator.clipboard.writeText(goLinkedInVacancyColumns)
            .then(() => console.log("Page info copied to clipboard:", goLinkedInVacancyColumns))
            .catch((err) => console.error("Failed to copy page info:", err));
    }
});

function remote() {
    return false;
}

function salary() {
    return false;
}

function date() {
    return new Intl.DateTimeFormat("en-CA").format(new Date());
}
