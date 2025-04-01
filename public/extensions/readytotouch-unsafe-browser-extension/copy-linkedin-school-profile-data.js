console.log("LinkedIn company profile data copy extension loaded");

// Ctrl + Shift + Y
document.body.addEventListener("keydown", (event) => {
    // Y is for English, Н is for Ukrainian
    if (event.ctrlKey && event.shiftKey && (event.key === "Y" || event.key === "Н")) {
        const [mainId, affiliatedIds] = ids();

        const goLinkedInProfileColumns = `				ID:                ${mainId},
				IDs:               ${toGoInts(affiliatedIds)},
				Alias:             "${parseVanityName(window.location.href)}",
				Name:              "${document.querySelector("h1").innerText.trim()}",
				Followers:         "${followers()}",
				Employees:         "${employees()}",
				AssociatedMembers: "${associatedMembers()}",
				Verified:          ${document.querySelectorAll('a[aria-label="Verified"]').length > 0 ? "true" : "false"},`

        navigator.clipboard.writeText(goLinkedInProfileColumns)
            .then(() => console.log("Page info copied to clipboard:", goLinkedInProfileColumns))
            .catch((err) => console.error("Failed to copy page info:", err));
    }
});

function parseVanityName(url) {
    // Your errors = your pain
    const error = "Expected URL like https://www.linkedin.com/school/company-name/";

    let parsedUrl = null;

    try {
        parsedUrl = new URL(url);
    } catch (e) {
        alert(error);

        return "";
    }

    const prefix = "/school/";

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

function followers() {
    const elements = document.querySelectorAll("div.org-top-card-summary-info-list__info-item");

    for (const element of elements) {
        const text = element.textContent.trim();

        if (text.endsWith("followers")) {
            return text.replace("followers", "").trim();
        }

        if (text.endsWith("follower")) {
            return text.replace("follower", "").trim();
        }
    }

    return "";
}

function employees() {
    const elements = document.querySelectorAll("a.org-top-card-summary-info-list__info-item");

    for (const element of elements) {
        const text = element.textContent.trim();

        if (text.endsWith("employees")) {
            return text.replace("employees", "").trim();
        }

        if (text.endsWith("employee")) {
            return text.replace("employee", "").trim();
        }
    }

    return "";
}

function associatedMembers() {
    const elements = document.querySelectorAll("h2");

    for (const element of elements) {
        const text = element.textContent.trim();

        if (text.endsWith("associated members")) {
            return text.replace("associated members", "").trim();
        }

        if (text.endsWith("associated member")) {
            return text.replace("associated member", "").trim();
        }
    }

    return "";
}

function ids() {
    const $codes = document.querySelectorAll("code");

    for (const $code of $codes) {
        const text = $code.textContent.trim();

        try {
            const json = JSON.parse(text);
            const ids = json?.data?.data?.organizationDashCompaniesByUniversalName?.["*elements"];

            if (Array.isArray(ids) && ids.length > 0) {
                const id = parseInt(ids[0].replace("urn:li:fsd_company:", ""), 10);

                return [id, affiliatedOrganizationsIds(id, json)];
            }
        } catch (error) {

        }
    }

    return [0, []];
}

function affiliatedOrganizationsIds(mainId, json) {
    const result = [];

    const elements = json?.included;

    if (Array.isArray(elements)) {
        for (const element of elements) {
            const ids = element?.affiliatedOrganizationsByJobs?.["*elements"];

            if (Array.isArray(ids)) {
                for (const id of ids) {
                    const affiliatedId = parseInt(id.replace("urn:li:fsd_company:", ""), 10);
                    if (affiliatedId !== mainId) {
                        result.push(affiliatedId);
                    }
                }
            }
        }
    }

    if (result.length === 0) {
        return [];
    }

    result.push(mainId);

    result.sort((a, b) => a - b)

    return result;
}

function toGoInts(ids) {
    if (ids.length === 0) {
        return "nil";
    }

    return `[]int{${ids.join(", ")}}`;
}
