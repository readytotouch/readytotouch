export class AliasImage {
    constructor(
        public readonly alt: string,
        public readonly src: string,
    ) {
    }
}

export class Alias {
    constructor(
        public readonly alias: string,
        public readonly name: string,
        public readonly image: AliasImage = null,
    ) {
    }
}
