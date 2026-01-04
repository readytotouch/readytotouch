import {Alias, AliasImage} from "./alias";
import {orderAliases} from "./order_aliases";

const all = [
    new Alias("3.0", "3.0 stars and above"),
    new Alias("3.5", "3.5 stars and above"),
    new Alias("4.0", "4.0 stars and above"),
    new Alias("4.5", "4.5 stars and above"),
];

export function ratingRatesAliases(aliases: Array<string>): Array<Alias> {
    return orderAliases(all, aliases);
}
