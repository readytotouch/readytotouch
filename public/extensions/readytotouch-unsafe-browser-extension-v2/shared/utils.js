"use strict";

// RTT namespace — loaded first in every content_scripts entry that needs it.
// All scripts access shared logic via window.RTT to avoid redeclaring globals.
window.RTT = window.RTT || {};

// ── Title normalization ────────────────────────────────────────────────────

RTT.normalizeTitle = function (raw) {
    return raw
        .replaceAll(" - ", " – ")
        .replaceAll(" — ", " – ")
        .replace(/\(m\/f\/[xd]\)/gi, "")
        .replace(/\(d\/f\/m\)/gi, "")
        .replace(/\(f\/m\/d\)/gi, "")
        .replace(/\(m\/w\/d\)/gi, "")
        .replace(/\(all genders\)/gi, "")
        .replace("Sr.", "Senior")
        .replace(/go\s*\/\s*golang/gi, "Golang")
        .replace(/golang\s*\/\s*go/gi, "Golang")
        .replace(/goland/gi, "Golang")
        .replace("Back End", "Back-End")
        .replace(/\s{2,}/g, " ")
        .trim();
};

// ── Dates ──────────────────────────────────────────────────────────────────

RTT.today = function () {
    return new Intl.DateTimeFormat("en-CA").format(new Date());
};

/**
 * Converts LinkedIn relative strings ("3 days ago") to an ISO date.
 * Falls back to today when the string isn't recognized.
 */
RTT.relativeDate = function (text) {
    const t = text.toLowerCase().trim();
    const now = new Date();

    const rules = [
        [/^(\d+) hours? ago$/,  (n) => now.setHours(now.getHours() - n)],
        [/^(\d+) days? ago$/,   (n) => now.setDate(now.getDate() - n)],
        [/^(\d+) weeks? ago$/,  (n) => now.setDate(now.getDate() - n * 7)],
        [/^(\d+) months? ago$/, (n) => now.setMonth(now.getMonth() - n)],
        [/^(\d+) years? ago$/,  (n) => now.setFullYear(now.getFullYear() - n)],
    ];

    for (const [re, apply] of rules) {
        const m = t.match(re);
        if (m) { apply(parseInt(m[1], 10)); break; }
    }

    return new Intl.DateTimeFormat("en-CA").format(now);
};

// ── Cloud providers ────────────────────────────────────────────────────────

RTT.cloudProviders = function (text) {
    const providers = [
        {pattern: /\b(aws|amazon web services)\b/i, value: "domain.AWS"},
        {pattern: /\b(gcp|google cloud)\b/i, value: "domain.GCP"},
        {pattern: /\b(azure|microsoft cloud)\b/i, value: "domain.Azure"}
    ];

    return providers
        .filter(p => p.pattern.test(text))
        .map(p => p.value);
};

// ── Clipboard + toast ──────────────────────────────────────────────────────

RTT.copyToClipboard = function (text) {
    return navigator.clipboard.writeText(text)
        .then(() => {
            console.log("[RTT] copied:", text);
            RTT.showToast("✓ Copied");
        })
        .catch((err) => console.error("[RTT] clipboard error:", err));
};

RTT.showToast = function (msg) {
    const el = document.createElement("div");
    el.textContent = msg || "✓ Copied";
    Object.assign(el.style, {
        position:     "fixed",
        bottom:       "20px",
        right:        "20px",
        background:   "#222",
        color:        "#fff",
        padding:      "8px 16px",
        borderRadius: "8px",
        fontSize:     "14px",
        zIndex:       "999999",
        opacity:      "1",
        transition:   "opacity 0.5s",
        fontFamily:   "sans-serif",
        boxShadow:    "0 4px 12px rgba(0,0,0,0.3)",
        pointerEvents:"none",
    });
    document.body.appendChild(el);
    setTimeout(() => {
        el.style.opacity = "0";
        setTimeout(() => el.remove(), 500);
    }, 2000);
};

// ── URL helpers ────────────────────────────────────────────────────────────

/**
 * Extracts the slug between `prefix` and the next `/`.
 * Example: parseVanityFromPath(url, "/company/", "Expected …") → "stripe"
 */
RTT.parseVanityFromPath = function (url, prefix, errorMsg) {
    let parsedUrl;
    try {
        parsedUrl = new URL(url);
    } catch (_) {
        alert(errorMsg);
        return "";
    }
    if (!parsedUrl.pathname.includes(prefix)) {
        alert(errorMsg);
        return "";
    }
    const start = parsedUrl.pathname.indexOf(prefix) + prefix.length;
    const end   = parsedUrl.pathname.indexOf("/", start);
    return end === -1
        ? parsedUrl.pathname.substring(start)
        : parsedUrl.pathname.substring(start, end);
};

// ── Go codegen helpers ─────────────────────────────────────────────────────

RTT.toGoInts = function (ids) {
    if (ids.length === 0) return "nil";
    return `[]int{${ids.join(", ")}}`;
};
