export type CompanyType = string;

export class Industry {
    constructor(
        public readonly alias: string,
        public readonly name: string,
    ) {}
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
