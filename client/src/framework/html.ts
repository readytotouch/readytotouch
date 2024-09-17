// https://github.com/vuejs/vue/blob/dev/src/compiler/parser/entity-decoder.js

const decoder = document.createElement("div");

// xss prevent
export function htmlDecode(html: string): string {
    decoder.innerHTML = html;
    return decoder.textContent;
}

export function s(html: string): string {
    return htmlDecode(html);
}

export function htmlToNode(html: string): HTMLElement {
    decoder.innerHTML = html.trim();
    return decoder.firstChild as HTMLElement;
}
