import {Alias} from "./alias";
import {orderAliases} from "./order_aliases";

const all = [
    new Alias("automotive", "Automotive"),
    new Alias("cybersecurity", "CyberSecurity"),
    // new Alias("edtech", "EdTech"),
    // new Alias("e-commerce", "eCommerce"),
    new Alias("healthtech", "HealthTech"),
    // new Alias("medtech", "MedTech"),
    new Alias("fintech", "FinTech"),
    // new Alias("gamedev", "GameDev"),
    // new Alias("iot", "IoT"),
    new Alias("adtech", "AdTech"),
    new Alias("martech", "MarTech"),
    new Alias("devops", "DevOps"),
    // new Alias("cloud-computing", "Cloud Computing"),
    new Alias("big-data", "Big Data"),
    new Alias("social-media", "Social Media"),
    new Alias("entertainment", "Entertainment"),
    new Alias("proptech", "PropTech"),
];

export function industries(aliases: Array<string>): Array<Alias> {
    return orderAliases(all, aliases);
}
