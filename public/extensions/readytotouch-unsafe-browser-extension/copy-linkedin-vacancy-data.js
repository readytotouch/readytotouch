console.log("LinkedIn vacancy data copy extension loaded");

// Ctrl + Shift + Y
// Ctrl + Shift + U
document.body.addEventListener("keydown", (event) => {
    // Y is for English, Н is for Ukrainian
    if (event.ctrlKey && event.shiftKey && (event.key === "Y" || event.key === "Н")) {

        const title = document.querySelector("h1").innerText.trim()
            .replaceAll(" - ", " – ") // Replace hyphen with dash
            .replaceAll(" — ", " – ") // Replace hyphen with dash
            .replace("(m/f/x)", " ").trim()
            .replace("(m/f/d)", " ").trim()
            .replace("(d/f/m)", " ").trim()
            .replace("(f/m/d)", " ").trim()
            .replace("(m/w/d)", " ").trim()
            .replace("(all genders)", " ").trim()
            .replace("Sr.", "Senior")
            .replace("GoLang", "Golang")
            .replace("Goland", "Golang")
            .replace("Go / Golang", "Golang")
            .replace("Go/Golang", "Golang")
            .replace("Back End", "Back-End")
        ;

        const url = normalizeURL(window.location.origin + window.location.pathname);

        const [withSalary, salaryComment] = salary();

        const goLinkedInVacancyColumns = `{
						    Title:                "${title}",
						    ShortDescription:     "",
						    SwitchingOpportunity: "",
						    URL:                  "${url}",
						    Location:             "${vacancyLocation()}",
						    Date:                 mustDate("${date()}"),
						    WithSalary:           ${withSalary ? "true" : "false"},${salaryComment}
						    Remote:               ${remote() ? "true" : "false"},
						},`

        navigator.clipboard.writeText(goLinkedInVacancyColumns)
            .then(() => console.log("Page info copied to clipboard:", goLinkedInVacancyColumns))
            .catch((err) => console.error("Failed to copy page info:", err));
    }

    // U is for English, Г is for Ukrainian
    if (event.ctrlKey && event.shiftKey && (event.key === "U" || event.key === "Г")) {
        const goLinkedInVacancyColumns = `mustDate("${date()}"), // `;

        navigator.clipboard.writeText(goLinkedInVacancyColumns)
            .then(() => console.log("Page info copied to clipboard:", goLinkedInVacancyColumns))
            .catch((err) => console.error("Failed to copy page info:", err));
    }
});

function remote() {
    {
        const $elements = document.querySelectorAll(".job-details-preferences-and-skills span");
        for (const $element of $elements) {
            if ($element.textContent.trim().toLowerCase() === "remote") {
                return true;
            }
        }
    }

    {
        const $elements = document.querySelectorAll(".job-details-fit-level-preferences span");
        for (const $element of $elements) {
            if ($element.textContent.trim().toLowerCase() === "remote") {
                return true;
            }
        }
    }

    return false;
}

function salary() {
    {
        const $elements = document.querySelectorAll(".job-details-preferences-and-skills span");
        for (const $element of $elements) {
            const s = $element.textContent.trim().toLowerCase();

            if (s.includes("$") || s.includes("€") || s.includes("£")) {
                return [true, ""];
            }
        }
    }

    {
        const $elements = document.querySelectorAll(".job-details-fit-level-preferences span");
        for (const $element of $elements) {
            const s = $element.textContent.trim().toLowerCase();

            if (s.includes("$") || s.includes("€") || s.includes("£")) {
                return [true, " // " + s];
            }
        }
    }

    return [false, ""];
}

function vacancyLocation() {
    const $elements = document.querySelectorAll(".job-details-jobs-unified-top-card__primary-description-container span.tvm__text");

    for (const $element of $elements) {
        return $element.textContent.trim();
    }

    return "";
}

function date() {
    const $elements = document.querySelectorAll(".job-details-jobs-unified-top-card__primary-description-container span");

    for (const $element of $elements) {
        const publishedAt = $element.textContent.trim().toLowerCase();

        if (publishedAt.includes("hour ago")) {
            const past = new Date();
            past.setHours(past.getHours() - 1);

            return new Intl.DateTimeFormat("en-CA").format(past);
        }

        const matchHours = publishedAt.match(/^(\d+) hours ago$/);
        if (matchHours) {
            const past = new Date();
            past.setHours(past.getHours() - parseInt(matchHours[1], 10));
            return new Intl.DateTimeFormat("en-CA").format(past);
        }

        if (publishedAt.includes("1 day ago")) {
            const yesterday = new Date();
            yesterday.setDate(yesterday.getDate() - 1);

            return new Intl.DateTimeFormat("en-CA").format(yesterday);
        }

        const matchDays = publishedAt.match(/^(\d+) days ago$/);
        if (matchDays) {
            const past = new Date();
            past.setDate(past.getDate() - parseInt(matchDays[1], 10));
            return new Intl.DateTimeFormat("en-CA").format(past);
        }

        if (publishedAt.includes("1 week ago")) {
            const past = new Date();
            past.setDate(past.getDate() - 7);

            return new Intl.DateTimeFormat("en-CA").format(past);
        }

        const matchWeeks = publishedAt.match(/^(\d+) weeks ago$/);
        if (matchWeeks) {
            const past = new Date();
            past.setDate(past.getDate() - 7 * parseInt(matchWeeks[1], 10));
            return new Intl.DateTimeFormat("en-CA").format(past);
        }

        if (publishedAt.includes("1 month ago")) {
            const past = new Date();
            past.setMonth(past.getMonth() - 1);

            return new Intl.DateTimeFormat("en-CA").format(past);
        }

        const matchMonths = publishedAt.match(/^(\d+) months ago$/);
        if (matchMonths) {
            const past = new Date();
            past.setMonth(past.getMonth() - parseInt(matchMonths[1], 10));

            return new Intl.DateTimeFormat("en-CA").format(past);
        }

        if (publishedAt.includes("1 year ago")) {
            const past = new Date();
            past.setMonth(past.getMonth() - 12);

            return new Intl.DateTimeFormat("en-CA").format(past);
        }

        const matchYears = publishedAt.match(/^(\d+) years ago$/);
        if (matchYears) {
            const past = new Date();
            past.setMonth(past.getMonth() - 12 * parseInt(matchYears[1], 10));

            return new Intl.DateTimeFormat("en-CA").format(past);
        }
    }

    return new Intl.DateTimeFormat("en-CA").format(new Date());
}

function normalizeURL(url) {
    const prefix = "https://www.linkedin.com/jobs/view/";

    if (url.startsWith(prefix)) {
        const match = url.match(/\/jobs\/view\/.*?(\d{7,})/);
        if (match && match[1]) {
            return prefix + match[1] + "/";
        }
    }

    return url;
}