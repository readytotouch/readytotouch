import {
    CompanyType,
    CloudProvider,
    Country,
    Industry,
} from "./organizers-v3-common-models";

export class CompanyResponse {
    constructor(
        public readonly id: number,
        public readonly type: CompanyType,
        /**
         * key: size (e.g. "72x72", "144x144")
         * value: image URL
         */
        public readonly logo: Record<string, string>,
        public readonly name: string,
        public readonly base_url: string,
        public readonly careers_url: string,
        public readonly about_url: string,
        public readonly blog_url: string,
        public readonly linkedin_profile: LinkedInProfileResponse,
        public readonly github_profile: GitHubProfileResponse,
        public readonly glassdoor_profile: GlassdoorProfileResponse,
        public readonly short_description: string,
        public readonly industries: Industry[],
        public readonly cloud_providers: CloudProvider[],
        public readonly has_employees_from_countries: Country[],
        public readonly rust_foundation_member: boolean,
        public readonly pinned_until: string | null,
        public readonly remote: boolean,
        public readonly latest_vacancy_date: string | null,
        public readonly github_repository_count: number,
        public favorite: boolean,
        public hidden: boolean,
    ) {
    }

    public emptyIndustries(): boolean {
        return this.industries === null || this.industries.length === 0;
    }

    public firstIndustry(): Industry | null {
        if (this.emptyIndustries()) {
            return null;
        }

        return this.industries[0];
    }
}

export class LinkedInProfileResponse {
    constructor(
        public readonly id: number,
        public readonly ids: number[],
        public readonly alias: string,
        public readonly name: string,
        public readonly employees: string,
        public readonly verified: boolean,
    ) {
    }
}

export class GitHubProfileResponse {
    constructor(
        public readonly login: string,
        public readonly followers: string,
        public readonly verified: boolean,
    ) {
    }
}

export class GlassdoorProfileResponse {
    constructor(
        public readonly overview_url: string,
        public readonly reviews_url: string,
        public readonly reviews_rate: string,
        public readonly verified: boolean,
    ) {
    }
}
