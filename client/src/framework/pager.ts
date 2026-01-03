export class Pager {
    private limit: number;
    private offset: number;
    private page: number;

    constructor(limit: number) {
        this.limit = limit;
        this.offset = 0;
        this.page = 1;
    }

    public setPage(page: number) {
        this.page = page;
        this.offset = this.limit * (page - 1);
    }

    public getPage(): number {
        return this.page;
    }

    public getOffset(): number {
        return this.offset;
    }

    public increment() {
        this.offset += this.limit;
        this.page += 1;
    }

    public incrementOffsetOnly() {
        this.offset += this.limit;
    }

    public reset() {
        this.offset = 0;
        this.page = 1;
    }
}

export function TotalPages(total: number, limit: number): number {
    if (total < limit) {
        return 1;
    }

    return Math.ceil(total / limit);
}