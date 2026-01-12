import {Alias, AliasImage} from "./alias";
import {orderAliases} from "./order_aliases";

const all = [
    new Alias("aws", "AWS"),
    new Alias("gcp", "GCP"),
    new Alias("azure", "Azure"),
    new Alias("digital-ocean", "Digital Ocean"),
    new Alias("vultr", "Vultr"),
];

export function cloudProviders(aliases: Array<string>): Array<Alias> {
    return orderAliases(all, aliases);
}
