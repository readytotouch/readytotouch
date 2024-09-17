import {Alias} from "./alias";
import {orderAliases} from "./order_aliases";

const all = [
    new Alias("product", "Product"),
    new Alias("startup", "Startup"),
    /*
        new Alias("outsource", "Outsource"),
        new Alias("outstaff", "Outstaff"),
     */
];

export function companyTypes(aliases: Array<string>): Array<Alias> {
    return orderAliases(all, aliases);
}
