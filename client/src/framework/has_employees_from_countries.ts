import {Alias, AliasImage} from "./alias";
import {orderAliases} from "./order_aliases";

const all = [
    new Alias("ukraine", "Has Ukrainian employees", new AliasImage("Flag of Ukraine with waves", "/assets/images/pages/online-new/ua_flag_with_waves.svg")),
    new Alias("czechia", "Has Czechs employees", new AliasImage("Flag of Czechia", "/assets/images/pages/online-new/cz_flag.svg")),
];

export function hasEmployeesFromCountries(aliases: Array<string>): Array<Alias> {
    return orderAliases(all, aliases);
}
