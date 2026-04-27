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
    if (!event.ctrlKey || !event.shiftKey) return;

    // Y / Н — full vacancy struct
    if (event.key === "Y" || event.key === "Н") {
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
    const selectors = [
        ".job-details-preferences-and-skills span",
        ".job-details-fit-level-preferences span",
    ];
    for (const sel of selectors) {
        for (const el of document.querySelectorAll(sel)) {
            if (el.textContent.trim().toLowerCase() === "remote") return true;
        }
    }
    return false;
}

function vacancySalary() {
    const currencies = ["$", "€", "£"];
    for (const sel of [".job-details-preferences-and-skills span", ".job-details-fit-level-preferences span"]) {
        for (const el of document.querySelectorAll(sel)) {
            const s = el.textContent.trim().toLowerCase();
            if (currencies.some(c => s.includes(c))) {
                const note = sel.includes("fit-level") ? ` // ${s}` : "";
                return [true, note];
            }
        }
    }
    return [false, ""];
}

function vacancyLocation() {
    const el = document.querySelector(
        ".job-details-jobs-unified-top-card__primary-description-container span.tvm__text"
    );
    return el ? el.textContent.trim() : "";
}

function vacancyDate() {
    for (const el of document.querySelectorAll(
        ".job-details-jobs-unified-top-card__primary-description-container span"
    )) {
        const text = el.textContent.trim();
        if (text) return RTT.relativeDate(text);
    }
    return RTT.today();
}

function normalizeVacancyURL() {
    if (window.location.href.startsWith("https://www.linkedin.com/jobs/search/")) {
        const id = new URLSearchParams(window.location.search).get("currentJobId");
        return `https://www.linkedin.com/jobs/view/${id}/`;
    }
    const raw = window.location.origin + window.location.pathname;
    const prefix = "https://www.linkedin.com/jobs/view/";
    if (raw.startsWith(prefix)) {
        const m = raw.match(/\/jobs\/view\/.*?(\d{7,})/);
        if (m) return prefix + m[1] + "/";
    }
    return raw;
}

// ── Ctrl+Shift+Y/Н — copy company profile data ────────────────────────────

function companyHandler(event) {
    if (!event.ctrlKey || !event.shiftKey) return;
    if (event.key !== "Y" && event.key !== "Н") return;

    const [mainId, affiliatedIds] = linkedinIds();
    const alias = RTT.parseVanityFromPath(
        window.location.href, "/company/",
        "Expected URL like https://www.linkedin.com/company/company-name/"
    );

    const out = `\t\t\t\tID:                ${mainId},
\t\t\t\tIDs:               ${RTT.toGoInts(affiliatedIds)},
\t\t\t\tAlias:             "${alias}",
\t\t\t\tName:              "${document.querySelector("h1")?.innerText.trim() ?? ""}",
\t\t\t\tFollowers:         "${linkedinFollowers()}",
\t\t\t\tEmployees:         "${linkedinEmployees()}",
\t\t\t\tAssociatedMembers: "${linkedinAssociatedMembers()}",
\t\t\t\tVerified:          ${document.querySelectorAll('a[aria-label="Verified"]').length > 0 ? "true" : "false"},
\t\t\t\tDate:              mustDate("${RTT.today()}"),`;

    RTT.copyToClipboard(out);
}

// ── Ctrl+Shift+Y/Н — copy school profile data ─────────────────────────────

function schoolHandler(event) {
    if (!event.ctrlKey || !event.shiftKey) return;
    if (event.key !== "Y" && event.key !== "Н") return;

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

// ── Shared LinkedIn DOM helpers ────────────────────────────────────────────

function linkedinFollowers() {
    for (const el of document.querySelectorAll("div.org-top-card-summary-info-list__info-item")) {
        const t = el.textContent.trim();
        if (/followers?$/.test(t)) return t.replace(/followers?$/, "").trim();
    }
    return "";
}

function linkedinEmployees() {
    for (const el of document.querySelectorAll("a.org-top-card-summary-info-list__info-item")) {
        const t = el.textContent.trim();
        if (/employees?$/.test(t)) return t.replace(/employees?$/, "").trim();
    }
    return "";
}

function linkedinAssociatedMembers() {
    for (const el of document.querySelectorAll("h2")) {
        const t = el.textContent.trim();
        if (/associated members?$/.test(t)) return t.replace(/associated members?$/, "").trim();
    }
    return "";
}

function linkedinIds() {
    for (const code of document.querySelectorAll("code")) {
        try {
            const json = JSON.parse(code.textContent.trim());
            const ids  = json?.data?.data?.organizationDashCompaniesByUniversalName?.["*elements"];
            if (Array.isArray(ids) && ids.length > 0) {
                const mainId = parseInt(ids[0].replace("urn:li:fsd_company:", ""), 10);
                return [mainId, linkedinAffiliatedIds(mainId, json)];
            }
        } catch (_) {}
    }
    return [0, []];
}

function linkedinAffiliatedIds(mainId, json) {
    const result = [];
    if (Array.isArray(json?.included)) {
        for (const el of json.included) {
            const ids = el?.affiliatedOrganizationsByJobs?.["*elements"];
            if (Array.isArray(ids)) {
                for (const id of ids) {
                    const n = parseInt(id.replace("urn:li:fsd_company:", ""), 10);
                    if (n !== mainId) result.push(n);
                }
            }
        }
    }
    if (result.length === 0) return [];
    result.push(mainId);
    result.sort((a, b) => a - b);
    return result;
}
