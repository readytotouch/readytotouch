import {minWidth} from "./responsive-constants";
import {debounce} from "./debounce";

function mobileFilterToggle($openMobileFilterButton, $body, $filterContainer, $closeMobileFilterButtons) {
    const changeVisible = function (active) {
        if (active) {
            $openMobileFilterButton.classList.remove("is-active");
            $openMobileFilterButton.setAttribute("aria-expanded", "false");
            $openMobileFilterButton.setAttribute("aria-label", "Open the filter");
            $body.classList.remove("disabled-scroll");
            $filterContainer.classList.remove("is-opened");
        } else {
            $openMobileFilterButton.classList.add("is-active");
            $openMobileFilterButton.setAttribute("aria-expanded", "false");
            $openMobileFilterButton.setAttribute("aria-label", "Close the filter");
            $body.classList.add("disabled-scroll");
            $filterContainer.classList.add("is-opened");
        }
    };

    $openMobileFilterButton.addEventListener("click", function () {
        changeVisible($openMobileFilterButton.classList.contains("is-active"));
    });

    for (const $closeMobileFilterButton of $closeMobileFilterButtons) {
        $closeMobileFilterButton.addEventListener("click", function () {
            changeVisible($openMobileFilterButton.classList.contains("is-active"));
        });
    }

    window.addEventListener(
        "resize",
        debounce(function () {
            if ($body.classList.contains("disabled-scroll") && window.innerWidth >= minWidth) {
                changeVisible(true);
            }
        }, 100),
    )
}

export function responsiveFilterWidget() {
    const $body = document.body;
    const $openMobileFilterButton = document.querySelector(".js-mobile-open-filter-container-button");
    const $filterContainer = document.querySelector(".js-filter-container");
    const $closeMobileFilterButtons = document.querySelectorAll(".js-mobile-close-filter-container-button") as any as Array<HTMLElement>;

    if ($openMobileFilterButton !== null) {
        mobileFilterToggle($openMobileFilterButton, $body, $filterContainer, $closeMobileFilterButtons)
    }
}
