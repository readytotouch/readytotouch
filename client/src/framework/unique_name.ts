export default class UniqueName {
    private readonly list: Array<string>;
    private readonly map: { [s: string]: boolean; };

    constructor() {
        this.list = [];
        this.map = {};
    }

    public unique(key: string): string {
        if (this.map.hasOwnProperty(key)) {
            throw new Error(`already exists: ${key}`);
        }

        this.map[key] = true;
        this.list.push(key);

        return key;
    }

    public names(): Array<string> {
        return this.list;
    }
}
