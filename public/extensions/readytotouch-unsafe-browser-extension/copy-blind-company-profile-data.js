console.log("Blind company profile data copy extension loaded");

// Ctrl + Shift + Y
document.body.addEventListener("keydown", (event) => {
    // Y is for English, Н is for Ukrainian
    if (event.ctrlKey && event.shiftKey && (event.key === "Y" || event.key === "Н")) {
        const {ratingValue, ratingCount} = getEmployerRating();
        const {employees, salary} = getCompanyInfo();

        const goBlindProfileColumns = `				Alias:       "${parseVanityName(window.location.href)}",
				Employees:   "${employees}",
				Salary:      "${salary}",
				Reviews:     "${ratingCount}",
				ReviewsRate: "${ratingValue}",`

        navigator.clipboard.writeText(goBlindProfileColumns)
            .then(() => console.log("Page info copied to clipboard:", goBlindProfileColumns))
            .catch((err) => console.error("Failed to copy page info:", err));
    }
});

function parseVanityName(url) {
    // Your errors = your pain
    const error = "Expected URL like https://www.teamblind.com/company/company-name/";

    let parsedUrl = null;

    try {
        parsedUrl = new URL(url);
    } catch (e) {
        alert(error);

        return "";
    }

    const prefix = "/company/";

    if (parsedUrl.pathname.indexOf(prefix) === -1) {
        alert(error);

        return "";
    }

    const end = parsedUrl.pathname.indexOf("/", prefix.length);
    if (end === -1) {
        return parsedUrl.pathname.substring(prefix.length);
    }

    return parsedUrl.pathname.substring(prefix.length, end);
}

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
    const $overview = document.querySelector("h1");

    if ($overview === null) {
        return {
            employees: "",
            salary: "",
        };
    }

    let employees = "";
    let salary = "";

    const $children = $overview.nextElementSibling.querySelectorAll("div.text-sm");
    for (const $child of $children) {
        if ($child.textContent.trim() === "Size") {
            employees = $child.nextElementSibling.textContent.trim()
                .replace("employees", "")
                .trim();
        }
        if ($child.textContent.trim() === "Salary") {
            salary = $child.nextElementSibling.textContent.trim();

            if (salary === "-") {
                salary = "";
            }
        }
    }

    return {employees, salary};
}
