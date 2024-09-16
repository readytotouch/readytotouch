import {CriteriaConverter} from "./criteria_converter"

export default class UrlStateContainer {
    private page: number;
    private criteriaMap: { [s: string]: any; };
    private criteria: string;
    private readonly criteriaNameConverterMap: { [s: string]: CriteriaConverter; };

    constructor(criteriaNames: Array<string>, criteriaNameConverterMap: { [s: string]: CriteriaConverter; }) {
        assertConverterMap(criteriaNames, criteriaNameConverterMap);

        const query = new URLSearchParams(window.location.search.substring(1));

        this.page = parsePage(query.get("page"));

        this.criteriaNameConverterMap = criteriaNameConverterMap;

        this.buildCriteriaMap(parseCriteriaMap(query, criteriaNames, criteriaNameConverterMap));
    }

    public getCriteria(criteriaName: string, defaultValue: any): any {
        if (this.criteriaMap.hasOwnProperty(criteriaName)) {
            return this.criteriaMap[criteriaName];
        }

        return defaultValue;
    }

    public setStringCriteria(key: string, value: string) {
        if (value !== "") {
            this.criteriaMap[key] = value;
        } else {
            delete this.criteriaMap[key];
        }
    }

    public setAnyCriteria(key: string, value: any) {
        this.criteriaMap[key] = value;
    }

    public setArrayCriteria(key: string, value: Array<string>) {
        if (value.length > 0) {
            this.criteriaMap[key] = value;
        } else {
            delete this.criteriaMap[key];
        }
    }

    public setBoolCriteria(key: string, value: boolean) {
        if (value === true) {
            this.criteriaMap[key] = value;
        } else {
            delete this.criteriaMap[key];
        }
    }

    public removeAlias(key: string, alias: string) {
        const current = this.getCriteria(key, []);

        const next = current.filter(value => value !== alias);

        this.setArrayCriteria(key, next);
    }

    public remove(key: string) {
        delete this.criteriaMap[key];
    }

    public keep(...keys: string[]) {
        for (let key in this.criteriaMap) {
            if (keys.indexOf(key) === -1) {
                delete this.criteriaMap[key];
            }
        }
    }

    public getPage(): number {
        return this.page;
    }

    public setPage(page: number) {
        this.page = page;
    }

    public storeCurrentState() {
        this.buildCriteriaMap(this.criteriaMap);

        this.update();
    }

    public createUrlByPageBuilder(): (page: number) => string {
        const criteria = this.criteria;

        return function (page: number): string {
            return window.location.pathname + getQueryParams(criteria, page);
        }
    }

    private update() {
        window.history.pushState(
            null,
            "",
            window.location.pathname + getQueryParams(this.criteria, this.page)
        );
    }

    private buildCriteriaMap(criteriaMap: {}) {
        const criteria = [];

        for (let criteriaName in criteriaMap) {
            if (criteriaMap.hasOwnProperty(criteriaName)) {
                const converter = this.criteriaNameConverterMap[criteriaName];

                criteria.push(
                    criteriaName + "=" + encodeURIComponent(converter.marshal(criteriaMap[criteriaName]))
                );
            }
        }

        this.criteriaMap = criteriaMap;
        this.criteria = criteria.join("&");
    }
}


function parseCriteriaMap(query, criteriaNames: Array<string>, criteriaNameConverterMap: { [s: string]: CriteriaConverter }): { [s: string]: any } {
    const criteriaMap = {};

    for (let i = 0; i < criteriaNames.length; i++) {
        const criteriaName = criteriaNames[i];

        if (query.has(criteriaName)) {
            const converter = criteriaNameConverterMap[criteriaName];

            const source = decodeURIComponent(query.get(criteriaName)).trim();

            if (source !== "") {
                const criteria = converter.unmarshal(source);

                if (criteria !== null) {
                    criteriaMap[criteriaName] = criteria;
                }
            }
        }
    }

    return criteriaMap;
}

function parsePage(s: string): number {
    if (s === "") {
        return 1;
    }

    const result = parseInt(s, 10);

    // if NaN | 0 | 1 -> false
    if (result > 1) {
        return result;
    }

    return 1;
}

function getQueryParams(criteria: string, page: number): string {
    if (criteria === "" && page === 1) {
        return "";
    }

    const result: Array<string> = [];

    if (criteria !== "") {
        result.push(criteria);
    }

    if (page > 1) {
        result.push(`page=${page}`);
    }

    return "?" + result.join("&");
}

function assertConverterMap(criteriaNames: Array<string>, criteriaNameConverterMap: { [s: string]: any; }) {
    for (let i = 0; i < criteriaNames.length; i++) {
        const criteriaName = criteriaNames[i];

        if (criteriaNameConverterMap.hasOwnProperty(criteriaName)) {
            continue;
        }

        throw new Error(`missing converter for "${criteriaName}"`);
    }
}
