import {createAliasMap} from "./alias_map";

export interface Checkboxes {
    setState(aliases: Array<string>)
    onChange(handler: (state: Array<string>) => void)
}

export class InputCheckboxes implements Checkboxes{
    constructor(private readonly $elements: Array<HTMLInputElement>) {
    }

    setState(aliases: Array<string>) {
        const aliasMap = createAliasMap(aliases, true);

        this.$elements.forEach(function ($element) {
            $element.checked = aliasMap.hasOwnProperty($element.getAttribute("data-alias"));
        })
    }

    onChange(handler: (state: Array<string>) => void) {
        const getState = this.getState.bind(this);

        this.$elements.forEach(function ($element) {
            $element.addEventListener("change", function () {
                handler(getState());
            });
        });
    }

    private getState(): Array<string> {
        const result = new Array<string>();

        this.$elements.forEach(function ($element) {
            if ($element.checked) {
                result.push($element.getAttribute("data-alias"));
            }
        })

        return result;
    }
}

export class ButtonCheckboxes implements Checkboxes{
    private readonly buttons: Array<ButtonCheckbox>;

    constructor($elements: Array<HTMLElement>) {
        this.buttons = Array.from($elements).map(function ($element) {
            return new ButtonCheckbox($element);
        })
    }

    setState(aliases: Array<string>) {
        const aliasMap = createAliasMap(aliases, true);

        this.buttons.forEach(function (button) {
            const checked = aliasMap.hasOwnProperty(button.$element.getAttribute("data-alias"));

            button.$input.checked = checked;
            button.$element.classList.toggle("is-checked", checked);
        })
    }

    onChange(handler: (state: Array<string>) => void) {
        const getState = this.getState.bind(this);

        this.buttons.forEach(function (button) {
            button.$element.addEventListener("click", (event) => {
                const checked = !button.$input.checked;
                button.$input.checked = checked;

                button.$element.classList.toggle("is-checked", checked);

                handler(getState());
            });
        });
    }

    private getState(): Array<string> {
        const result = new Array<string>();

        this.buttons.forEach(function ($button) {
            if ($button.$input.checked) {
                result.push($button.$element.getAttribute("data-alias"));
            }
        })

        return result;
    }
}

class ButtonCheckbox {
    public readonly $element: HTMLElement;
    public readonly $input: HTMLInputElement;

    constructor($element: HTMLElement) {
        this.$element = $element;
        this.$input = $element.querySelector("input") as HTMLInputElement;
    }
}
