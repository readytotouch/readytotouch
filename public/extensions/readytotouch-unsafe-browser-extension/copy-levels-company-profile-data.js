console.log("Levels.fyi company profile data copy extension loaded");

// Ctrl + Shift + Y
document.body.addEventListener("keydown", (event) => {
    // Y is for English, Н is for Ukrainian
    if (event.ctrlKey && event.shiftKey && (event.key === "Y" || event.key === "Н")) {
        const {employees} = getCompanyInfo();

        const goBlindProfileColumns = `				Alias:     "${parseVanityName(window.location.href)}",
				Employees: "${employees}",`

        navigator.clipboard.writeText(goBlindProfileColumns)
            .then(() => console.log("Page info copied to clipboard:", goBlindProfileColumns))
            .catch((err) => console.error("Failed to copy page info:", err));
    }
});

function parseVanityName(url) {
    // Your errors = your pain
    const error = "Expected URL like https://www.levels.fyi/companies/company-name/";

    let parsedUrl = null;

    try {
        parsedUrl = new URL(url);
    } catch (e) {
        alert(error);

        return "";
    }

    const prefix = "/companies/";

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

function getCompanyInfo() {
    let employees = "";

    const $elements = document.querySelectorAll("h6");
    for (const $element of $elements) {

        if ($element.nextElementSibling !== null && $element.nextElementSibling.textContent.trim() === "# of Employees") {
            employees = $element.textContent.trim();
        }
    }

    return {employees};
}
