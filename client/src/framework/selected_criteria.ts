export function renderSelected(name: string): string {
    return `<li class="filter-used__item">
      ${name}
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
