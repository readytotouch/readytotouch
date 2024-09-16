import {Alias} from "./alias";
import {createAliasMap} from "./alias_map";

export function orderAliases(all: Array<Alias>, aliases: Array<string>): Array<Alias> {
    if (aliases.length === 0) {
        return [];
    }

    const aliasMap = createAliasMap(aliases, true);

    const result = [];

    for (let alias of all) {
        if (aliasMap.hasOwnProperty(alias.alias)) {
            result.push(alias);
        }
    }

    return result;
}
