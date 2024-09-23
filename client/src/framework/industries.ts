import {Alias} from "./alias";
import {orderAliases} from "./order_aliases";

const all = [
    new Alias("cybersecurity", "CyberSecurity"),
    new Alias("edtech", "EdTech"),
    new Alias("ecommerce", "eCommerce"),
    new Alias("healthtech", "HealthTech"),
    new Alias("medtech", "MedTech"),
    new Alias("fintech", "FinTech"),
    new Alias("gamedev", "GameDev"),
    new Alias("iot", "IoT"),
    new Alias("adtech", "AdTech"),
];

export function industries(aliases: Array<string>): Array<Alias> {
    return orderAliases(all, aliases);
}
