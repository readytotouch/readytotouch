import {htmlToNode} from "./html";

// language=HTML
const example = `<a class="pagination__button pagination__button--prev" href="javascript:void(0);" aria-label="arrow left">
  <img
    class="pagination__icon pagination__icon--prev"
    width="14"
    height="14"
    src="/assets/images/pages/common/left.svg"
    alt="arrow left"
  />
</a>
<a class="pagination__item pagination__item--active" href="javascript:void(0);">1</a>
<a class="pagination__item" href="javascript:void(0);">2</a>
<a class="pagination__item" href="javascript:void(0);">3</a>
<a class="pagination__item" href="javascript:void(0);">4</a>
<a class="pagination__item pagination__item--hidden" href="javascript:void(0);">5</a>
<span class="pagination__item pagination__item--hidden">...</span>
<a class="pagination__item pagination__item--hidden" href="javascript:void(0);">9</a>
<a class="pagination__button pagination__button--prev" href="javascript:void(0);" aria-label="arrow right">
  <img
    class="pagination__icon pagination__icon--prev"
    width="14"
    height="14"
    src="/assets/images/pages/common/right.svg"
    alt="arrow left"
  />
</a>`;
/**
 * The pagination component has a typo in the class name of the next button, but it was left as is because fixing it broke the design.
 **/

export default class Pagination {
    private readonly $element: HTMLElement;
    private readonly setPageHandler: (page: number) => void;

    constructor(setPageHandler: (page: number) => void) {
        this.$element = document.getElementById("js-pagination");
        this.setPageHandler = setPageHandler;
    }

    public render(currentPage: number, totalPages: number, urlByPageBuilder: (page: number) => string): void {
        const $elements = new Array<HTMLElement>();

        const prevPage = currentPage - 1;

        // prev arrow
        {
            const $prevPage = htmlToNode(prev(prevPage, urlByPageBuilder));

            this.addEventListener($prevPage, prevPage);

            $elements.push($prevPage);
        }

        // first 3 …
        {
            const max3 = Math.min(prevPage, 3);

            for (let page = 1; page <= max3; page++) {
                const $page = htmlToNode(`<a class="pagination__item" href="${urlByPageBuilder(page)}">${page}</a>`);

                this.addEventListener($page, page)

                $elements.push($page);
            }

            // …
            if (prevPage > max3 + 1) {
                const page = Math.floor((prevPage + max3) / 2);

                const $page = htmlToNode(`<a class="pagination__item" href="${urlByPageBuilder(page)}">…</a>`);

                this.addEventListener($page, page)

                $elements.push($page);
            }

            // prev
            if (prevPage > max3) {
                const $page = htmlToNode(`<a class="pagination__item" href="${urlByPageBuilder(prevPage)}">${prevPage}</a>`);

                this.addEventListener($page, prevPage)

                $elements.push($page);
            }
        }

        // current
        {
            const $currentPage = htmlToNode(`<a class="pagination__item pagination__item--active" href="${urlByPageBuilder(currentPage)}">${currentPage}</a>`);

            this.addEventListener($currentPage, 0);

            $elements.push($currentPage);
        }

        const nextPage = findNextPage(currentPage, totalPages);

        // end 3 …
        if (nextPage > 0) {
            const min3 = Math.max(nextPage, totalPages - 2);

            // next page
            if (nextPage < min3) {
                const $page = htmlToNode(`<a class="pagination__item" href="${urlByPageBuilder(nextPage)}">${nextPage}</a>`);

                this.addEventListener($page, nextPage)

                $elements.push($page);
            }

            // …
            if (nextPage < min3 - 1) {
                const page = Math.floor((nextPage + min3) / 2);

                const $page = htmlToNode(`<a class="pagination__item" href="${urlByPageBuilder(page)}">…</a>`);

                this.addEventListener($page, page)

                $elements.push($page);
            }

            // end 3
            for (let page = min3; page <= totalPages; page++) {
                const $page = htmlToNode(`<a class="pagination__item" href="${urlByPageBuilder(page)}">${page}</a>`);

                this.addEventListener($page, page)

                $elements.push($page);
            }
        }

        // next arrow
        {
            const $nextPage = htmlToNode(next(nextPage, urlByPageBuilder));

            this.addEventListener($nextPage, nextPage);

            $elements.push($nextPage);
        }

        this.$element.innerHTML = "";
        this.$element.append(...$elements);
    }

    private addEventListener($page: HTMLElement, page: number) {
        const setPageHandler = this.setPageHandler;

        $page.addEventListener("click", function (event) {
            event.preventDefault();

            if (page > 0) {
                setPageHandler(page);
            }
        });
    }
}

function prev(page: number, urlByPageBuilder: (page: number) => string): string {
    const url = page > 0 ? urlByPageBuilder(page) : "javascript:void(0);";

    return `<a
    class="pagination__button pagination__button--prev"
    href="${url}"
    aria-label="arrow left"
>
	<img
        class="pagination__icon pagination__icon--prev"
        width="14"
        height="14"
        src="/assets/images/pages/common/left.svg"
        alt="arrow left"
	/>
</a>`;
}

function next(page: number, urlByPageBuilder: (page: number) => string): string {
    const url = page > 0 ? urlByPageBuilder(page) : "javascript:void(0);";

    return `<a
    class="pagination__button pagination__button--prev"
    href="${url}"
    aria-label="arrow right"
>
    <img
        class="pagination__icon pagination__icon--prev"
        width="14"
        height="14"
        src="/assets/images/pages/common/right.svg"
        alt="arrow left"
    />
</a>`;
}

function findNextPage(currentPage: number, totalPages: number) {
    if (currentPage < totalPages) {
        return currentPage + 1;
    }

    return 0;
}
