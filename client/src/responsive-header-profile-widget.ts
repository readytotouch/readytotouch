import {maxWidth, minWidth} from "./responsive-constants";
import {debounce} from "./debounce";

function hamburgerToggle($hamburger, $body, $mainHeaderNav) {
    const changeVisible = function (active) {
        if (active) {
            $hamburger.classList.remove("is-active");
            $hamburger.setAttribute("aria-expanded", "false");
            $hamburger.setAttribute("aria-label", "Open the menu");
            $body.classList.remove("disabled-scroll");
            $mainHeaderNav.classList.remove("is-opened");
        } else {
            $hamburger.classList.add("is-active");
            $hamburger.setAttribute("aria-expanded", "true");
            $hamburger.setAttribute("aria-label", "Close the menu");
            $body.classList.add("disabled-scroll");
            $mainHeaderNav.classList.add("is-opened");
        }
    }

    $hamburger.addEventListener("click", function () {
        changeVisible($hamburger.classList.contains("is-active"));
    });

    window.addEventListener(
        "resize",
        debounce(function () {
            if ($body.classList.contains("disabled-scroll") && window.innerWidth >= minWidth) {
                changeVisible(true);
            }
        }, 100),
    )
}

function profileToggle($headerProfileButton, $modalProfile, $headerProfileAvatar) {
    $headerProfileButton.addEventListener("click", function (event) {
        if (window.innerWidth >= minWidth) {
            event.stopPropagation();
            $modalProfile.classList.toggle("active");
        }
    })

    $headerProfileAvatar.addEventListener("click", function () {
        if (window.innerWidth <= maxWidth) {
            $headerProfileAvatar.classList.toggle("is-mobile-profile-opened");
        }
    })

    document.addEventListener("click", function (event) {
        if (window.innerWidth >= minWidth && !$modalProfile.contains(event.target) && event.target !== $headerProfileButton) {
            $modalProfile.classList.remove("active");
        }
    })
}

export function responsiveHeaderProfileWidget() {
    const $body = document.body;
    const $hamburger = document.querySelector(".js-hamburger");
    const $mainHeaderNav = document.querySelector(".js-header-nav");
    const $headerProfileButton = document.querySelector(".js-header-profile-button");
    const $modalProfile = document.querySelector(".js-modal-profile");
    const $headerProfileAvatar = document.querySelector(".js-header-profile-avatar");

    if ($hamburger !== null) {
        hamburgerToggle($hamburger, $body, $mainHeaderNav)
    }

    if ($modalProfile !== null || $headerProfileAvatar !== null) {
        profileToggle($headerProfileButton, $modalProfile, $headerProfileAvatar)
    }
}
