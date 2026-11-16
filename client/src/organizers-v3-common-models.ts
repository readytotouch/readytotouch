export type CompanyType = string;

export class Industry {
    constructor(
        public readonly alias: string,
        public readonly name: string,
    ) {}
}

export interface CompanyHidden {
    id: number;
    industries: Industry[],
    hidden: boolean,
}

export class CloudProvider {
    constructor(
        public readonly alias: string,
        public readonly name: string,
    ) {}
}

export class Country {
    constructor(
        public readonly alias: string,
        public readonly name: string,
    ) {}
}

export class LocationCountryResponse {
    constructor(
        public readonly code: string,
    ) {
    }
}

export class LocationResponse {
    constructor(
        public readonly raw: string,
        public readonly country: LocationCountryResponse,
    ) {
    }
}

export function companyEmptyIndustries(company: CompanyHidden): boolean {
    return company.industries === null || company.industries.length === 0;
}

export function companyFirstIndustry(company: CompanyHidden): Industry | null {
    if (companyEmptyIndustries(company)) {
        return null;
    }

    return company.industries[0];
}
