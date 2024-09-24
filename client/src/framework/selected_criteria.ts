import {AliasImage} from "./alias";

function renderImage(image: AliasImage) {
    if (image === null) {
        return "";
    }

    return `<img class="filter-used__content-image" 
        alt="${image.alt}" 
        width="20" 
        height="20" 
        src="${image.src}"
/>`;
}

export function renderSelected(name: string, image: AliasImage = null): string {
    return `<li class="filter-used__item">
      ${name}
      ${renderImage(image)}
      <button class="filter-used__link" type="button">
        <img
          class="filter-used__cross-icon"
          alt="cross icon"
          width="10"
          height="10"
          title="info"
          src="/assets/images/pages/common/cross.svg"
        />
      </button>
    </li>`;
}
