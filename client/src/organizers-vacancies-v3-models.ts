import {VacancySource} from "./vacancy-source";
import {
    CompanyType,
    Industry,
    CloudProvider,
    Country,
    LocationResponse,
} from "./organizers-v3-common-models";

export class VacancyResponse {
    constructor(
        public readonly id: number,
        public readonly title: string,
        public readonly short_description: string,
        public readonly location: LocationResponse,
        public readonly source: VacancySource,
        public readonly cloud_providers: CloudProvider[],
        public readonly remote: boolean,
        public readonly date: string,
        public readonly pinned_until: string | null,
        public readonly monthly_views: number,
        public company: VacancyCompanyResponse, // to be filled later
        public favorite: boolean,
    ) {
    }
}

export class VacancyCompanyResponse {
    constructor(
        public readonly id: number,
        public readonly type: CompanyType,
        /**
         * key: size (e.g. "72x72", "144x144")
         * value: image URL
         */
        public readonly logo: Record<string, string>,
        public readonly name: string,
        public readonly linkedin_profile: VacancyCompanyLinkedInProfileResponse,
        public readonly glassdoor_profile: VacancyCompanyGlassdoorProfileResponse,
        public readonly industries: Industry[],
        public readonly has_employees_from_countries: Country[],
        public readonly rust_foundation_member: boolean,
    ) {
    }
}

export class VacancyCompanyLinkedInProfileResponse {
    constructor(
        public readonly alias: string, // vanity name
        public readonly employees: string,
    ) {
    }
}

export class VacancyCompanyGlassdoorProfileResponse {
    constructor(
        public readonly reviews_rate: string,
    ) {
    }
}
