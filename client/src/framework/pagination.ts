import {htmlToNode} from "./html";

// language=HTML
const example = `<a class="pagination__button pagination__button--prev" href="javascript:void(0);" aria-label="arrow left">
    <img class="pagination__icon pagination__icon--prev"
         width="14"
         height="14"
         src="/assets/images/pages/vacancies2/arrow-left.svg"
         alt="arrow left"
    >
</a>

<a class="pagination__item pagination__item--active" href="javascript:void(0);">1</a>
<a class="pagination__item" href="javascript:void(0);">2</a>
<a class="pagination__item" href="javascript:void(0);">3</a>
<span class="pagination__item">…</span>
<a class="pagination__item" href="javascript:void(0);">55</a>
<a class="pagination__item" href="javascript:void(0);">56</a>
<a class="pagination__item" href="javascript:void(0);">57</a>
<span class="pagination__item">…</span>
<a class="pagination__item" href="javascript:void(0);">98</a>
<a class="pagination__item" href="javascript:void(0);">99</a>
<a class="pagination__item" href="javascript:void(0);">100</a>

<a class="pagination__button pagination__button--next" href="javascript:void(0);" aria-label="arrow right"
>
    <img class="pagination__icon pagination__icon--next"
         width="14"
         height="14"
         src="/assets/images/pages/vacancies2/arrow-left.svg"
         alt="arrow right"
    >
</a>`;

export default class Pagination {
    private readonly $element: HTMLElement;
    private readonly setPageHandler: (page: number) => void;
    private readonly appVersion: number;

    constructor(setPageHandler: (page: number) => void, appVersion: number) {
        this.$element = document.getElementById("js-pagination");
        this.setPageHandler = setPageHandler;
        this.appVersion = appVersion;
    }

    public render(currentPage: number, totalPages: number, urlByPageBuilder: (page: number) => string): void {
        const $elements = new Array<HTMLElement>();

        const prevPage = currentPage - 1;

        // prev arrow
        {
            const $prevPage = htmlToNode(prev(prevPage, urlByPageBuilder, this.appVersion));

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
            const $nextPage = htmlToNode(next(nextPage, urlByPageBuilder, this.appVersion));

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

function prev(page: number, urlByPageBuilder: (page: number) => string, appVersion: number): string {
    const url = page > 0 ? urlByPageBuilder(page) : "javascript:void(0);";

    return `<a class="pagination__button pagination__button--prev"
   href="${url}"
   aria-label="arrow left"
>
	<img class="pagination__icon pagination__icon--prev"
		 width="14"
		 height="14"
		 src="${prevArrowSrc(appVersion)}"
		 alt="arrow left"
	>
</a>`;
}

function prevArrowSrc(appVersion: number): string {
    switch (appVersion) {
        case 1:
            return "/assets/images/pages/vacancies/arrow-left.svg";
        case 2:
            return "/assets/images/pages/vacancies2/arrow-left.svg";
    }

    return "";
}

function next(page: number, urlByPageBuilder: (page: number) => string, appVersion: number): string {
    const url = page > 0 ? urlByPageBuilder(page) : "javascript:void(0);";

    return `<a class="pagination__button pagination__button--next"
   href="${url}"
   aria-label="arrow right"
>
	<img class="pagination__icon pagination__icon--next"
		 width="14"
		 height="14"
		 src="${nextArrowSrc(appVersion)}"
		 alt="arrow right"
	>
</a>`;
}

function nextArrowSrc(appVersion: number): string {
    switch (appVersion) {
        case 1:
            return "/assets/images/pages/vacancies/arrow-right.svg";
        case 2:
            return "/assets/images/pages/vacancies2/arrow-left.svg";
    }

    return "";
}

function findNextPage(currentPage: number, totalPages: number) {
    if (currentPage < totalPages) {
        return currentPage + 1;
    }

    return 0;
}
