export const defaultDateFormatter = new Intl.DateTimeFormat("en-CA", {
    day: "2-digit",
    month: "2-digit",
    year: "numeric",
});

export const prettyDateFormatter = new Intl.DateTimeFormat("en-GB", {
    day: "2-digit",
    month: "short",
    year: "numeric",
});

export function formatDiffDate(date: Date): string {
    const today = new Date();
    today.setUTCHours(0, 0, 0, 0);

    if (date >= today) {
        return "Today";
    }

    const yesterday = new Date(today);
    yesterday.setUTCDate(today.getUTCDate() - 1);

    if (date >= yesterday) {
        return "Yesterday";
    }

    const diffMs = today.getTime() - date.getTime();
    const days = Math.floor(diffMs / (24 * 60 * 60 * 1000));

    if (days === 1) {
        return "1 day ago";
    }

    return `${days} days ago`;
}
