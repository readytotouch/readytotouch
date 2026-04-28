"use strict";

console.log("[RTT] LinkedIn coordinator loaded");

// ── SPA routing ────────────────────────────────────────────────────────────
//
// LinkedIn is a SPA: navigation happens via history.pushState without a full
// page reload. Browser content scripts are injected once and stay alive, so
// the old keydown handler from /jobs/view/ would remain active on /company/.
//
// Fix: intercept pushState/popstate, swap the handler on every URL change.

const ROUTES = [
    { pattern: /\/jobs\/(view|search)\//, handler: vacancyHandler },
    { pattern: /\/company\//,             handler: companyHandler },
    { pattern: /\/school\//,              handler: schoolHandler  },
];

let _activeHandler = null;

(function patchHistory() {
    const _push = history.pushState.bind(history);
    history.pushState = function (...args) {
        _push(...args);
        activate();
    };
    window.addEventListener("popstate", activate);
})();

function activate() {
    const url     = window.location.href;
    const matched = ROUTES.find(r => r.pattern.test(url));
    const next    = matched ? matched.handler : null;

    if (_activeHandler === next) return; // same route, nothing to swap

    if (_activeHandler) {
        document.body.removeEventListener("keydown", _activeHandler);
        console.log("[RTT] removed handler for previous route");
    }

    _activeHandler = next;

    if (_activeHandler) {
        document.body.addEventListener("keydown", _activeHandler);
        console.log("[RTT] activated handler for", url);
    }
}

activate(); // run once on initial page load

// ── Ctrl+Shift+Y/Н — copy vacancy data ────────────────────────────────────

function vacancyHandler(event) {
    if (event.ctrlKey && event.shiftKey && (event.key === "Y" || event.key === "Н")) {
        const titleRaw   = document.querySelector("h1")?.innerText.trim() ?? "";
        const title      = RTT.normalizeTitle(titleRaw);
        const descText   = (document.querySelector("h2.text-heading-large + div.mt4")?.innerText ?? "");
        const providers  = RTT.cloudProviders(title + " " + descText);
        const [hasSalary, salaryNote] = vacancySalary();

        const out = `{
\t\t\t\t\t\t    Title:                "${title}",
\t\t\t\t\t\t    ShortDescription:     "",
\t\t\t\t\t\t    SwitchingOpportunity: "",
\t\t\t\t\t\t    URL:                  "${normalizeVacancyURL()}",
\t\t\t\t\t\t    Location:             "${vacancyLocation()}",
\t\t\t\t\t\t    CloudProviders:       []domain.CloudProvider{${providers.join(", ")}},
\t\t\t\t\t\t    Date:                 mustDate("${vacancyDate()}"),
\t\t\t\t\t\t    WithSalary:           ${hasSalary ? "true" : "false"},${salaryNote}
\t\t\t\t\t\t    Remote:               ${vacancyRemote() ? "true" : "false"},
\t\t\t\t\t\t},`;

        RTT.copyToClipboard(out);
    }

    // U / Г — just the date (quick-update shortcut)
    if (event.key === "U" || event.key === "Г") {
        RTT.copyToClipboard(`mustDate("${vacancyDate()}"), // `);
    }
}

function vacancyRemote() {
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

function vacancySalary() {
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

function vacancyDate() {
    const $elements = document.querySelectorAll(".job-details-jobs-unified-top-card__primary-description-container span");

    for (const $element of $elements) {
        const publishedAt = $element.textContent.trim().toLowerCase();
        if (publishedAt) {
            return RTT.relativeDate(publishedAt);
        }
    }

    return RTT.today();
}

function normalizeVacancyURL() {
    if (window.location.href.startsWith("https://www.linkedin.com/jobs/search/")) {
        const params = new URLSearchParams(window.location.search);

        const id = params.get("currentJobId");

        return normalizeURL(`https://www.linkedin.com/jobs/view/${id}/`);
    }

    return normalizeURL(window.location.origin + window.location.pathname);
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

// ── Ctrl+Shift+Y/Н — copy company profile data ────────────────────────────

function companyHandler(event) {
    if (event.ctrlKey && event.shiftKey && (event.key === "Y" || event.key === "Н")) {
        const [mainId, affiliatedIds] = ids();
        const alias = RTT.parseVanityFromPath(
            window.location.href, "/company/",
            "Expected URL like https://www.linkedin.com/company/company-name/"
        );

        const out = `\t\t\t\tID:                ${mainId},
\t\t\t\tIDs:               ${RTT.toGoInts(affiliatedIds)},
\t\t\t\tAlias:             "${alias}",
\t\t\t\tName:              "${document.querySelector("h1").innerText.trim()}",
\t\t\t\tFollowers:         "${followers()}",
\t\t\t\tEmployees:         "${employees()}",
\t\t\t\tAssociatedMembers: "${associatedMembers()}",
\t\t\t\tVerified:          ${document.querySelectorAll('a[aria-label="Verified"]').length > 0 ? "true" : "false"},
\t\t\t\tDate:              mustDate("${RTT.today()}"),`;

        RTT.copyToClipboard(out);
    }
}

// ── Ctrl+Shift+Y/Н — copy school profile data ─────────────────────────────

function schoolHandler(event) {
    if (event.ctrlKey && event.shiftKey && (event.key === "Y" || event.key === "Н")) {
        const [mainId, affiliatedIds] = linkedinIds();
        const alias = RTT.parseVanityFromPath(
            window.location.href, "/school/",
            "Expected URL like https://www.linkedin.com/school/school-name/"
        );

        const out = `\t\t\t\tID:                ${mainId},
\t\t\t\tIDs:               ${RTT.toGoInts(affiliatedIds)},
\t\t\t\tAlias:             "${alias}",
\t\t\t\tName:              "${document.querySelector("h1")?.innerText.trim() ?? ""}",
\t\t\t\tFollowers:         "${linkedinFollowers()}",
\t\t\t\tEmployees:         "${linkedinEmployees()}",
\t\t\t\tAssociatedMembers: "${linkedinAssociatedMembers()}",
\t\t\t\tVerified:          ${document.querySelectorAll('a[aria-label="Verified"]').length > 0 ? "true" : "false"},`;

        RTT.copyToClipboard(out);
    }
}

// ── Shared LinkedIn DOM helpers ────────────────────────────────────────────

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
    const $elements = document.querySelectorAll("h2");

    for (const $element of $elements) {
        const $text = $element.textContent.trim();

        if ($text.endsWith("associated members")) {
            return $text.replace("associated members", "").trim();
        }

        if ($text.endsWith("associated member")) {
            return $text.replace("associated member", "").trim();
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
