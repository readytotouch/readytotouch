export default class Range {
    constructor(
        public readonly from: number,
        public readonly to: number,
    ) {
    }

    equals(range: Range): boolean {
        return this.from === range.from && this.to === range.to;
    }

    great(range: Range): boolean {
        return (range.to > 0 && this.to >= range.to) || (this.from > 0 && this.from >= range.from);
    }
}

export function ParseStringRange(fromString: string, toString: string): Range {
    return new Range(
        parse(fromString),
        parse(toString),
    );
}

function parse(value: string): number {
    if (value === "") {
        return 0;
    }

    const result = parseInt(value);

    if (result > 0) {
        return result;
    }

    return 0;
}
