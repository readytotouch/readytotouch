import {Alias} from "./alias";
import {orderAliases} from "./order_aliases";

const all = [
    new Alias("2-10", "2-10 employees"),
    new Alias("11-50", "11-50 employees"),
    new Alias("51-200", "51-200 employees"),
    new Alias("201-500", "201-500 employees"),
    new Alias("501-1K", "501-1K employees"),
    new Alias("1K-5K", "1K-5K employees"),
    new Alias("5K-10K", "5K-10K employees"),
    new Alias("10K+", "10K+ employees"),
];

export function linkedinCompanySizes(aliases: Array<string>): Array<Alias> {
    return orderAliases(all, aliases);
}
