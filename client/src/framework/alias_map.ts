export function createAliasMap(aliases: Array<string>, state: boolean): { [s: string]: boolean; } {
    const map = {};

    for (let i = 0; i < aliases.length; i++) {
        map[aliases[i]] = state;
    }

    return map;
}

export function filterAliasMap(aliasMap: { [s: string]: boolean; }): Array<string> {
    const result = [];

    for (let alias in aliasMap) {
        if (aliasMap.hasOwnProperty(alias) && aliasMap[alias] === true) {
            result.push(alias);
        }
    }

    return result;
}
