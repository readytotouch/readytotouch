export function firstQuerySelector($element: HTMLElement, selector: string): HTMLElement {
    return $element.querySelector(selector) as HTMLElement;
}
