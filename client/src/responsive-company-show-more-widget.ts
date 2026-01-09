function addEvent($button) {
    $button.addEventListener("click", function (event) {
        const $text = event.target.querySelector(".js-toggle-text");

        if (event.target.nextElementSibling.classList.contains("is-visible")) {
            event.target.nextElementSibling.classList.remove("is-visible");
            $text.innerText = "Show more info";
        } else {
            event.target.nextElementSibling.classList.add("is-visible");
            $text.innerText = "Hide info";
        }
    });
}

export function responsiveCompanyShowMoreWidget() {
    const $buttons = document.querySelectorAll(".js-mobile-company-show-more-info-button");

    $buttons.forEach(addEvent);
}

export function addCompanyShowMoreEvent($company: HTMLElement) {
    const $button = $company.querySelector(".js-mobile-company-show-more-info-button");

    addEvent($button);
}
