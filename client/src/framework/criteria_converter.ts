import Range from "./range"

export interface CriteriaConverter {
    unmarshal(data: string): any;

    marshal(data: any): string;
}

export class IdentityCriteriaConverter implements CriteriaConverter {
    unmarshal(data) {
        return data;
    }

    marshal(data) {
        return data;
    }
}

export class CheckedCriteriaConverter implements CriteriaConverter {
    unmarshal(data) {
        return data === "1";
    }

    marshal(data) {
        return data === true ? "1" : "";
    }
}

export class RangeCriteriaConverter implements CriteriaConverter {
    private readonly delimiter: string;

    constructor() {
        this.delimiter = "-";
    }

    unmarshal(data): Range {
        const between = data.split(this.delimiter);
        const from = parseInt(between[0], 10);

        if (between.length === 2) {
            const to = parseInt(between[1], 10);

            if (to > 0 && to >= from) {
                return new Range(from, to);
            }

            return null;
        }

        if (from > 0) {
            return new Range(from, 0);
        }

        return new Range(0, 0);
    }

    marshal({from, to}) {
        const data = [from];

        if (to > 0) {
            data.push(to);
        }

        return data.join(this.delimiter);
    }
}

const multiSelectDelimiter = ",";

export class MultiSelectCriteriaConverter implements CriteriaConverter {
    unmarshal(aliases) {
        if (aliases !== "") {
            return aliases.split(multiSelectDelimiter);
        }

        return null;
    }

    marshal(aliases) {
        if (aliases.length === 0) {
            return "";
        }

        return aliases.join(multiSelectDelimiter);
    }
}

// in previous version filtered by "checked" & "unchecked"
const multiCheckboxSuffix = "-unchecked";

// did use in previous, but will use in future
class MultiCheckboxCriteriaConverter implements CriteriaConverter {
    unmarshal(source) {
        if (source !== "") {
            const aliases = source.split(",");

            const result = {};

            for (let i = 0; i < aliases.length; i++) {
                const alias = aliases[i];

                if (alias.endsWith(multiCheckboxSuffix)) {
                    result[alias.substring(0, alias.length - multiCheckboxSuffix.length)] = false;
                } else {
                    result[alias] = true;
                }
            }

            return result;
        }

        return null;
    }

    marshal(aliasCheckMap) {
        const aliases = [];

        for (let alias in aliasCheckMap) {
            if (aliasCheckMap.hasOwnProperty(alias)) {
                const checked = aliasCheckMap[alias];

                if (checked === true) {
                    aliases.push(alias);
                } else if (checked === false) {
                    aliases.push(alias + multiCheckboxSuffix);
                } else {
                    console.error(`wrong state for "${alias}" = "${checked}"`)
                }
            }
        }

        return aliases.join(",");
    }
}

export class AliasCriteriaConverter implements CriteriaConverter {
    unmarshal(source) {
        if (source !== "") {
            const aliases = source.split(",");

            for (let i = 0; i < aliases.length; i++) {
                const alias = aliases[i];

                if (alias.endsWith(multiCheckboxSuffix)) {
                    return alias.substring(0, alias.length - multiCheckboxSuffix.length)
                }

                return alias;
            }
        }

        return null;
    }

    marshal(alias) {
        return alias;
    }
}
