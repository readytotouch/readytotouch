import {Alias} from "./alias";
import {orderAliases} from "./order_aliases";

const all = [
    new Alias("cybersecurity", "CyberSecurity"),
    new Alias("embedded", "Embedded"),
    new Alias("edtech", "EdTech"),
    new Alias("ecommerce", "eCommerce"),
    new Alias("healthtech", "HealthTech"),
    new Alias("medtech", "MedTech"),
    new Alias("fintech", "FinTech"),
    new Alias("foodtech", "FoodTech"),
    new Alias("gamedev", "GameDev"),
    new Alias("iot", "IoT"),
    new Alias("ai", "AI/ML/DL"),
    new Alias("blockchain", "Blockchain"),
    new Alias("hrtech", "HRTech"),
    new Alias("adtech", "AdTech"),
    new Alias("proptech", "PropTech"),
    new Alias("insurtech", "InsurTech"),
    new Alias("cryptocurrency", "Cryptocurrency"),
    new Alias("adult", "Adult"),
    new Alias("dating", "Dating"),
    new Alias("gambling", "Gambling"),
    new Alias("betting", "Betting"),
];

export function directions(aliases: Array<string>): Array<Alias> {
    return orderAliases(all, aliases);
}