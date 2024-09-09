import {TimeCountStats} from "./domain";

export function formatDay(item: TimeCountStats) {
    return new Date(item.time).toLocaleDateString("en-us", {day: "2-digit", month: "long"});
}
