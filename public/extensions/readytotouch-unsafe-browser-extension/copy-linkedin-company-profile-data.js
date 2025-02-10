console.log("LinkedIn company profile data copy extension loaded");

// Ctrl + Shift + Y
document.body.addEventListener("keydown", (event) => {
    if (event.ctrlKey && event.shiftKey && event.key === "Y") {
        const goLinkedInProfileColumns = `
				ID:                0,
				Alias:             "${parseVanityName(window.location.href)}",
				Name:              "${document.querySelector("h1").innerText.trim()}",
				Followers:         "${followers()}",
				Employees:         "${employees()}",
				AssociatedMembers: "${associatedMembers()}",
				Verified:          ${document.querySelectorAll('a[aria-label="Verified"]').length > 0 ? "true" : "false"},
        `

        navigator.clipboard.writeText(goLinkedInProfileColumns)
            .then(() => console.log("Page info copied to clipboard:", goLinkedInProfileColumns))
            .catch((err) => console.error("Failed to copy page info:", err));
    }
});

function parseVanityName(url) {
    // Your errors = your pain
    const error = "Expected URL like https://www.linkedin.com/company/company-name/";

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

function followers() {
    const elements = document.querySelectorAll("div.org-top-card-summary-info-list__info-item");

    for (const element of elements) {
        const text = element.textContent.trim();

        if (text.endsWith("followers")) {
            return text.replace("followers", "").trim();
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
    }

    return "";
}
